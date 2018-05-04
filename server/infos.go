package server

import (
	"net/http"

	"github.com/gpaulo00/gh0st/models"
	"github.com/gin-gonic/gin"
)

// InfoController is a HTTP controller to manage Infos
type InfoController struct{}

// List returns a list of all Infos
func (ctl *InfoController) List(c *gin.Context) {
	var w []models.Info
	if err := models.DB().Model(&w).Select(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": w})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": w})
}

// Create adds a new Info
func (ctl *InfoController) Create(c *gin.Context) {
	// binding
	var form models.Info
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": true})
}

// Route configures gin to route the controller
func (ctl *InfoController) Route(r gin.IRouter) {
	r.GET("/infos", ctl.List)
	r.GET("/infos/:id", ctl.Get)
	r.POST("/infos", ctl.Create)
	r.PUT("/infos/:id", ctl.Update)
	r.PATCH("/infos/:id", ctl.Update)
	r.DELETE("/infos/:id", ctl.Delete)
}
