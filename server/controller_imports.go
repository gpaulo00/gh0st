package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gpaulo00/gh0st/models"
	"github.com/gpaulo00/gh0st/plugins/custom"
	"github.com/gpaulo00/gh0st/plugins/nmap"
)

var errNoWorkspace = errors.New("no specified workspace")

// ImportController is a HTTP controller to manage imports
type ImportController struct{}

// Import inserts some data into database
func (ctl *ImportController) Import(c *gin.Context) {
	// binding
	var form custom.ImportSource
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, Error(err))
		return
	}

	// run in transaction
	db := models.DB()
	err := db.RunInTransaction(form.Import)

	// handle error
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}

	// result
	c.JSON(http.StatusOK, Done)
}

// Nmap parses a nmap xml scan and insert its data into database
func (ctl *ImportController) Nmap(c *gin.Context) {
	// get workspace
	raw, ok := c.GetQuery("workspace")
	if !ok {
		c.JSON(http.StatusBadRequest, Error(errNoWorkspace))
		return
	}
	ws, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error(err))
		return
	}

	// get file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, Error(err))
		return
	}

	// open file
	fp, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}
	defer fp.Close()

	// run in transaction
	db := models.DB()
	if err := nmap.New(db, ws).Parse(fp); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}

	// result
	c.JSON(http.StatusOK, Done)
}

// Route configures gin to route the controller
func (ctl *ImportController) Route(r gin.IRouter) {
	r.POST("/import", ctl.Import)
	r.POST("/import/nmap", ctl.Nmap)
}
