package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/orm"
	"github.com/gpaulo00/gh0st/models"
)

// HostController is a HTTP controller to manage Hosts
type HostController struct{}

// List returns a list of all Hosts
func (ctl *HostController) List(c *gin.Context) {
	w := []models.Host{}
	err := models.DB().Model(&w).
		Apply(orm.Pagination(c.Request.URL.Query())).
		Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Get return a single Host
func (ctl *HostController) Get(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// find Host
	w := models.Host{ID: id}
	if err := models.DB().Select(&w); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Create adds a new Host
func (ctl *HostController) Create(c *gin.Context) {
	// binding
	var form models.Host
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, Error(err))
		return
	}

	// insert
	if _, err := models.DB().Model(&form).Returning("*").Insert(); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}
	c.JSON(http.StatusOK, form)
}

// Delete removes a Host
func (ctl *HostController) Delete(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// delete
	w := models.Host{ID: id}
	if err := models.DB().Delete(&w); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}
	c.JSON(http.StatusOK, Done)
}

// Route configures gin to route the controller
func (ctl *HostController) Route(r gin.IRouter) {
	const path = "/hosts"
	r.GET(path, ctl.List)
	r.POST(path, ctl.Create)
	r.GET(path+"/:id", ctl.Get)
	r.DELETE(path+"/:id", ctl.Delete)
}
