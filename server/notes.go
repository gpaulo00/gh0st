package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gpaulo00/gh0st/models"
)

// NoteController is a HTTP controller to manage Notes
type NoteController struct{}

// List returns a list of all Notes
func (ctl *NoteController) List(c *gin.Context) {
	w := []models.Note{}
	if err := models.DB().Model(&w).Select(); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Get return a single Note
func (ctl *NoteController) Get(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// find Note
	w := models.Note{ID: id}
	if err := models.DB().Select(&w); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Create adds a new Note
func (ctl *NoteController) Create(c *gin.Context) {
	// binding
	var form models.Note
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

// Update modifies a Note
func (ctl *NoteController) Update(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// find
	form := models.Note{ID: id}
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

// Delete removes a Note
func (ctl *NoteController) Delete(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// delete
	w := models.Note{ID: id}
	if err := models.DB().Delete(&w); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}
	c.JSON(http.StatusOK, models.Done)
}

// Route configures gin to route the controller
func (ctl *NoteController) Route(r gin.IRouter) {
	r.GET(models.NotePath.String(), ctl.List)
	r.POST(models.NotePath.String(), ctl.Create)
	r.GET(models.NotePath.ID(), ctl.Get)
	r.PUT(models.NotePath.ID(), ctl.Update)
	r.PATCH(models.NotePath.ID(), ctl.Update)
	r.DELETE(models.NotePath.ID(), ctl.Delete)
}
