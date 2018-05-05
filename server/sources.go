package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gpaulo00/gh0st/models"
)

// SourceController is a HTTP controller to manage Sources
type SourceController struct{}

// List returns a list of all Sources
func (ctl *SourceController) List(c *gin.Context) {
	w := []models.Source{}
	if err := models.DB().Model(&w).Select(); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Get return a single Source
func (ctl *SourceController) Get(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// find Source
	w := models.Source{ID: id}
	if err := models.DB().Select(&w); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Create adds a new Source
func (ctl *SourceController) Create(c *gin.Context) {
	// binding
	var form models.Source
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, models.Error(err))
		return
	}

	// insert
	if _, err := models.DB().Model(&form).Returning("*").Insert(); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}
	c.JSON(http.StatusOK, form)
}

// Delete removes a Source
func (ctl *SourceController) Delete(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// delete
	w := models.Source{ID: id}
	if err := models.DB().Delete(&w); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}
	c.JSON(http.StatusOK, models.Done)
}

// Route configures gin to route the controller
func (ctl *SourceController) Route(r gin.IRouter) {
	r.GET(models.SourcePath.String(), ctl.List)
	r.POST(models.SourcePath.String(), ctl.Create)
	r.GET(models.SourcePath.ID(), ctl.Get)
	r.DELETE(models.SourcePath.ID(), ctl.Delete)
}
