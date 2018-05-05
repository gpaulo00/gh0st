package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gpaulo00/gh0st/models"
)

// InfoController is a HTTP controller to manage Infos
type InfoController struct{}

// List returns a list of all Infos
func (ctl *InfoController) List(c *gin.Context) {
	w := []models.Info{}
	if err := models.DB().Model(&w).Select(); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Get return a single Info
func (ctl *InfoController) Get(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// find Info
	w := models.Info{ID: id}
	if err := models.DB().Select(&w); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Create adds a new Info
func (ctl *InfoController) Create(c *gin.Context) {
	// binding
	var form models.Info
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

// Update modifies a Info
func (ctl *InfoController) Update(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// find
	form := models.Info{ID: id}
	if err := models.DB().Select(&form); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	// binding
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, models.Error(err))
		return
	}

	// updates
	if err := models.DB().Update(&form); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	c.JSON(http.StatusOK, form)
}

// Delete removes a Info
func (ctl *InfoController) Delete(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// delete
	w := models.Info{ID: id}
	if err := models.DB().Delete(&w); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}
	c.JSON(http.StatusOK, models.Done)
}

// Route configures gin to route the controller
func (ctl *InfoController) Route(r gin.IRouter) {
	r.GET(models.InfoPath.String(), ctl.List)
	r.POST(models.InfoPath.String(), ctl.Create)
	r.GET(models.InfoPath.ID(), ctl.Get)
	r.PUT(models.InfoPath.ID(), ctl.Update)
	r.PATCH(models.InfoPath.ID(), ctl.Update)
	r.DELETE(models.InfoPath.ID(), ctl.Delete)
}
