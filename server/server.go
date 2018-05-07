package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gpaulo00/gh0st/models"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Controller is type that provides routes
type Controller interface {
	Route(r gin.IRouter)
}

func parseID(c *gin.Context) (uint64, error) {
	// parse id
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResult{Error: err.Error()})
		return 0, err
	}

	return id, nil
}

// Server creates a HTTP server
func Server() {
	// set mode
	if models.Version != "develop" {
		gin.SetMode(gin.ReleaseMode)
	} else if mode := viper.GetString("http.mode"); mode != "" {
		gin.SetMode(mode)
	}

	// router
	r := gin.Default()
	new(WorkspaceController).Route(r)
	new(HostController).Route(r)
	new(SourceController).Route(r)
	new(ServiceController).Route(r)
	new(NoteController).Route(r)
	new(IssueController).Route(r)
	new(ImportController).Route(r)

	// http server
	address := viper.GetString("http.address")
	srv := &http.Server{
		Addr:    address,
		Handler: r,
	}

	// listen for interrupt signal, and graceful shutdown
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.WithError(err).Fatal("shutdown error")
		}
		log.Info("shutdown server")
	}()

	// listen
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.WithError(err).Fatal("listen error")
	}
}
