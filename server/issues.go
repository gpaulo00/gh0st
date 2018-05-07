package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gpaulo00/gh0st/models"
)

// IssueController is a HTTP controller to manage Issues
type IssueController struct{}

// List returns a list of all Issues
func (ctl *IssueController) List(c *gin.Context) {
	w := []models.Issue{}
	if err := models.DB().Model(&w).Select(); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Get return a single Issue
func (ctl *IssueController) Get(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// find Issue
	w := models.Issue{ID: id}
	if err := models.DB().Select(&w); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Create adds a new Issue
func (ctl *IssueController) Create(c *gin.Context) {
	// binding
	var form models.Issue
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

// Update modifies a Issue
func (ctl *IssueController) Update(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// find
	form := models.Issue{ID: id}
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

// Delete removes a Issue
func (ctl *IssueController) Delete(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// delete
	w := models.Issue{ID: id}
	if err := models.DB().Delete(&w); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}
	c.JSON(http.StatusOK, models.Done)
}

// Route configures gin to route the controller
func (ctl *IssueController) Route(r gin.IRouter) {
	r.GET(models.IssuePath.String(), ctl.List)
	r.POST(models.IssuePath.String(), ctl.Create)
	r.GET(models.IssuePath.ID(), ctl.Get)
	r.PUT(models.IssuePath.ID(), ctl.Update)
	r.PATCH(models.IssuePath.ID(), ctl.Update)
	r.DELETE(models.IssuePath.ID(), ctl.Delete)
}
