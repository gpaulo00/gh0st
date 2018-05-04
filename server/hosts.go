package server

import (
	"net/http"

	"github.com/gpaulo00/gh0st/models"
	"github.com/gin-gonic/gin"
)

// HostController is a HTTP controller to manage Hosts
type HostController struct{}

// List returns a list of all Hosts
func (ctl *HostController) List(c *gin.Context) {
	var w []models.Host
	if err := models.DB().Model(&w).Select(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": w})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": w})
}

// Create adds a new Host
func (ctl *HostController) Create(c *gin.Context) {
	// binding
	var form models.Host
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": true})
}

// Route configures gin to route the controller
func (ctl *HostController) Route(r gin.IRouter) {
	r.GET("/hosts", ctl.List)
	r.GET("/hosts/:id", ctl.Get)
	r.POST("/hosts", ctl.Create)
	r.DELETE("/hosts/:id", ctl.Delete)
}
