package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/gpaulo00/gh0st/models"
	log "github.com/sirupsen/logrus"
)

// ImportController is a HTTP controller to manage imports
type ImportController struct{}

func (ctl *ImportController) insertMany(tx *pg.Tx, input []interface{}) error {
	if len(input) > 0 {
		if _, err := tx.Model(input...).Insert(); err != nil {
			return err
		}
	}
	return nil
}

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
	var (
		services []interface{}
		notes    []interface{}
		issues   []interface{}
	)

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
			var wg sync.WaitGroup

			// parse services
			go func(i []models.Service) {
				wg.Add(1)
				defer wg.Done()
				for _, srv := range i {
					srv.HostID = hostID
					services = append(services, &srv)
				}
			}(form.Hosts[i].Services)

			// parse notes
			go func(i []models.Note) {
				wg.Add(1)
				defer wg.Done()
				for _, note := range i {
					note.HostID = hostID
					notes = append(notes, &note)
				}
			}(form.Hosts[i].Notes)

			// parse issues
			go func(i []models.Issue) {
				wg.Add(1)
				defer wg.Done()
				for _, issue := range i {
					issue.HostID = hostID
					issues = append(issues, &issue)
					fmt.Println(issues)
				}
			}(form.Hosts[i].Issues)

			wg.Wait()
		}

		// add services, notes & issues
		if err := ctl.insertMany(tx, services); err != nil {
			return err
		}
		if err := ctl.insertMany(tx, notes); err != nil {
			return err
		}
		return ctl.insertMany(tx, issues)
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
