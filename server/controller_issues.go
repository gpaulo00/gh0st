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
		c.JSON(http.StatusInternalServerError, Error(err))
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
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Create adds a new Issue
func (ctl *IssueController) Create(c *gin.Context) {
	// binding
	var form models.Issue
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
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}

	// binding
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, Error(err))
		return
	}

	// updates
	if err := models.DB().Update(&form); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
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
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}
	c.JSON(http.StatusOK, Done)
}

// Route configures gin to route the controller
func (ctl *IssueController) Route(r gin.IRouter) {
	const path = "/issues"
	r.GET(path, ctl.List)
	r.POST(path, ctl.Create)
	r.GET(path+"/:id", ctl.Get)
	r.PUT(path+"/:id", ctl.Update)
	r.PATCH(path+"/:id", ctl.Update)
	r.DELETE(path+"/:id", ctl.Delete)
}
