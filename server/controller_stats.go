package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gpaulo00/gh0st/models"
)

// StatController is a HTTP controller to manage charts data
type StatController struct{}

// Issues returns stats of the issues
func (ctl *StatController) Issues(c *gin.Context) {
	// initialize
	db := models.DB()
	var res models.IssuesSummary

	if err := db.Model(&res).Select(); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// Route configures gin to route the controller
func (ctl *StatController) Route(r gin.IRouter) {
	const path = "/stats"
	r.GET(path+"/issues", ctl.Issues)
}
