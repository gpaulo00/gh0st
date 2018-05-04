package models

import (
	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var db *pg.DB

// DB returns the database connection
func DB() *pg.DB {
	return db
}

// ConfigureDB configures the connection pool to PostgreSQL.
func ConfigureDB() error {
	// open connection
	db = pg.Connect(&pg.Options{
		Addr:     viper.GetString("database.address"),
		Database: viper.GetString("database.name"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		OnConnect: func(z *pg.DB) error {
			log.Info("successful connection to database")
			return nil
		},
	})

	// test connection
	_, err := db.Exec("SELECT 1")
	return err
}
