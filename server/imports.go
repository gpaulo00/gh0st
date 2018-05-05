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
	var infos []interface{}

	// transaction
	db := models.DB()
	err := db.RunInTransaction(func(tx *pg.Tx) error {
		// add source
		if err := tx.Insert(&form.Source); err != nil {
			return err
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

		// parse services & infos
		for i := range form.Hosts {
			hostID := hosts[i].(*models.Host).ID
			srv := form.Hosts[i].Services
			inf := form.Hosts[i].Infos

			// TODO: add concurrency
			// parse services
			for j := range srv {
				srv[j].HostID = hostID
				services = append(services, &srv[j])
			}

			// parse infos
			for j := range inf {
				inf[j].HostID = hostID
				infos = append(infos, &inf[j])
			}
		}

		// add services & infos
		if _, err := tx.Model(services...).Insert(); err != nil {
			return err
		}
		_, err := tx.Model(infos...).Insert()

		return err
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
		Infos:    len(infos),
	}
	log.Debug(result.String())
	c.JSON(http.StatusOK, result)
}

// Route configures gin to route the controller
func (ctl *ImportController) Route(r gin.IRouter) {
	r.POST(models.ImportPath, ctl.Import)
}
