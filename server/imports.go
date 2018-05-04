package server

import (
	"net/http"

	"github.com/gpaulo00/gh0st/models"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
)

// ImportController is a HTTP controller to manage imports
type ImportController struct{}

type importForm struct {
	Source models.Source `json:"source" binding:"required"`
	Hosts  []struct {
		Host     models.Host      `json:"host" binding:"required"`
		Infos    []models.Info    `json:"infos"`
		Services []models.Service `json:"services"`
	} `json:"hosts" binding:"required"`
}

// Import inserts some data into the database
func (ctl *ImportController) Import(c *gin.Context) {
	// binding
	var form importForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// result
	log.Debugf(
		"imported data: hosts = %d, services = %d, infos = %d",
		len(hosts), len(services), len(infos),
	)
	c.JSON(http.StatusOK, gin.H{
		"result": gin.H{
			"hosts":    len(hosts),
			"services": len(services),
			"infos":    len(infos),
		},
	})
}

// Route configures gin to route the controller
func (ctl *ImportController) Route(r gin.IRouter) {
	r.POST("/import", ctl.Import)
}
