package server

import (
	"net/http"

	"github.com/gpaulo00/gh0st/models"
	"github.com/gin-gonic/gin"
)

// WorkspaceController is a HTTP controller to manage Workspaces
type WorkspaceController struct{}

// List returns a list of all Workspaces
func (ctl *WorkspaceController) List(c *gin.Context) {
	var w []models.Workspace
	if err := models.DB().Model(&w).Select(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": w})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": w})
}

// Create adds a new Workspace
func (ctl *WorkspaceController) Create(c *gin.Context) {
	// binding
	var form models.Workspace
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// binding
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// updates
	if err := models.DB().Update(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": form})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": true})
}

// Route configures gin to route the controller
func (ctl *WorkspaceController) Route(r gin.IRouter) {
	r.GET("/workspaces", ctl.List)
	r.GET("/workspaces/:id", ctl.Get)
	r.POST("/workspaces", ctl.Create)
	r.PUT("/workspaces/:id", ctl.Update)
	r.PATCH("/workspaces/:id", ctl.Update)
	r.DELETE("/workspaces/:id", ctl.Delete)
}
