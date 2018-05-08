package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gpaulo00/gh0st/models"
)

// ServiceController is a HTTP controller to manage Services
type ServiceController struct{}

// List returns a list of all Services
func (ctl *ServiceController) List(c *gin.Context) {
	w := []models.Service{}
	if err := models.DB().Model(&w).Select(); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Get return a single Service
func (ctl *ServiceController) Get(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// find Service
	w := models.Service{ID: id}
	if err := models.DB().Select(&w); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Create adds a new Service
func (ctl *ServiceController) Create(c *gin.Context) {
	// binding
	var form models.Service
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

// Delete removes a Service
func (ctl *ServiceController) Delete(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// delete
	w := models.Service{ID: id}
	if err := models.DB().Delete(&w); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}
	c.JSON(http.StatusOK, Done)
}

// Route configures gin to route the controller
func (ctl *ServiceController) Route(r gin.IRouter) {
	const path = "/services"
	r.GET(path, ctl.List)
	r.POST(path, ctl.Create)
	r.GET(path+"/:id", ctl.Get)
	r.DELETE(path+"/:id", ctl.Delete)
}
