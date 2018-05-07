package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/gpaulo00/gh0st/models"
	log "github.com/sirupsen/logrus"
)

// ImportController is a HTTP controller to manage imports
type ImportController struct{}

// Import inserts some data into the database
func (ctl *ImportController) Import(c *gin.Context) {
	// binding
	var form models.ImportForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, models.Error(err))
		return
	}

	// variables
	hosts := make([]interface{}, len(form.Hosts))
	var services []interface{}
	var notes []interface{}
	var issues []interface{}

	// transaction
	db := models.DB()
	err := db.RunInTransaction(func(tx *pg.Tx) error {
		// add source
		if err := tx.Insert(&form.Source); err != nil {
			return err
		}

		// without hosts?
		if len(form.Hosts) <= 0 {
			return nil
		}

		// add hosts
		for i := range form.Hosts {
			host := &form.Hosts[i].Host
			host.SourceID = form.Source.ID
			hosts[i] = host
		}
		if _, err := tx.Model(hosts...).Insert(); err != nil {
			return err
		}

		// parse services, notes & issues
		for i := range form.Hosts {
			hostID := hosts[i].(*models.Host).ID

			// TODO: add concurrency
			// parse services
			for _, srv := range form.Hosts[i].Services {
				srv.HostID = hostID
				services = append(services, &srv)
			}

			// parse notes
			for _, note := range form.Hosts[i].Notes {
				note.HostID = hostID
				notes = append(notes, &note)
			}

			// parse issues
			for _, issue := range form.Hosts[i].Issues {
				issue.HostID = hostID
				issues = append(issues, &issue)
			}
		}

		// add services, notes & issues
		if len(services) > 0 {
			if _, err := tx.Model(services...).Insert(); err != nil {
				return err
			}
		}

		if len(notes) > 0 {
			if _, err := tx.Model(notes...).Insert(); err != nil {
				return err
			}
		}

		if len(issues) > 0 {
			if _, err := tx.Model(issues...).Insert(); err != nil {
				return err
			}
		}

		return nil
	})

	// handle error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	// result
	result := models.ImportResult{
		Hosts:    len(hosts),
		Services: len(services),
		Notes:    len(notes),
		Issues:   len(issues),
	}
	log.Debug(result.String())
	c.JSON(http.StatusOK, result)
}

// Route configures gin to route the controller
func (ctl *ImportController) Route(r gin.IRouter) {
	r.POST(models.ImportPath, ctl.Import)
}
