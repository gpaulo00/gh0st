package server

import (
	"net/http"

	"github.com/gpaulo00/gh0st/models"
	"github.com/gin-gonic/gin"
)

// ServiceController is a HTTP controller to manage Services
type ServiceController struct{}

// List returns a list of all Services
func (ctl *ServiceController) List(c *gin.Context) {
	var w []models.Service
	if err := models.DB().Model(&w).Select(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": w})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": w})
}

// Create adds a new Service
func (ctl *ServiceController) Create(c *gin.Context) {
	// binding
	var form models.Service
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insert
	if _, err := models.DB().Model(&form).Returning("*").Insert(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": form})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": true})
}

// Route configures gin to route the controller
func (ctl *ServiceController) Route(r gin.IRouter) {
	r.GET("/services", ctl.List)
	r.GET("/services/:id", ctl.Get)
	r.POST("/services", ctl.Create)
	r.DELETE("/services/:id", ctl.Delete)
}
