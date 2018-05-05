package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gpaulo00/gh0st/models"
)

// WorkspaceController is a HTTP controller to manage Workspaces
type WorkspaceController struct{}

// List returns a list of all Workspaces
func (ctl *WorkspaceController) List(c *gin.Context) {
	w := []models.Workspace{}
	if err := models.DB().Model(&w).Select(); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Get return a single Workspace
func (ctl *WorkspaceController) Get(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// find Workspace
	w := models.Workspace{ID: id}
	if err := models.DB().Select(&w); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}

	c.JSON(http.StatusOK, w)
}

// Create adds a new Workspace
func (ctl *WorkspaceController) Create(c *gin.Context) {
	// binding
	var form models.Workspace
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

// Update modifies a Workspace
func (ctl *WorkspaceController) Update(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// find workspace
	form := models.Workspace{ID: id}
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

// Delete removes a Workspace
func (ctl *WorkspaceController) Delete(c *gin.Context) {
	// parse id
	id, err := parseID(c)
	if err != nil {
		return
	}

	// delete
	w := models.Workspace{ID: id}
	if err := models.DB().Delete(&w); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error(err))
		return
	}
	c.JSON(http.StatusOK, models.Done)
}

// Route configures gin to route the controller
func (ctl *WorkspaceController) Route(r gin.IRouter) {
	r.GET(models.WorkspacePath.String(), ctl.List)
	r.POST(models.WorkspacePath.String(), ctl.Create)
	r.GET(models.WorkspacePath.ID(), ctl.Get)
	r.PUT(models.WorkspacePath.ID(), ctl.Update)
	r.PATCH(models.WorkspacePath.ID(), ctl.Update)
	r.DELETE(models.WorkspacePath.ID(), ctl.Delete)
}