package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func parseID(c *gin.Context) (uint64, error) {
	// parse id
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, err
	}

	return id, nil
}

// Server creates a HTTP server
func Server() {
	// set mode
	if mode := viper.GetString("http.mode"); mode != "" {
		gin.SetMode(mode)
	}

	// router
	r := gin.Default()
	new(WorkspaceController).Route(r)
	new(HostController).Route(r)
	new(SourceController).Route(r)
	new(ServiceController).Route(r)
	new(InfoController).Route(r)
	new(ImportController).Route(r)

	// http server
	srv := &http.Server{
		Addr:    viper.GetString("http.address"),
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
