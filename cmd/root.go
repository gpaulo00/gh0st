package cmd

import (
	"fmt"
	"os"

	"github.com/gpaulo00/gh0st/models"
	"github.com/gpaulo00/gh0st/server"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gh0std",
	Short:   "Simple and lightweight reporting framework",
	Version: models.Version,
	Run: func(cmd *cobra.Command, args []string) {
		// connect to database
		if err := models.ConfigureDB(); err != nil {
			log.WithError(err).Fatal("cannot connect to database")
		}

		// start HTTP server
		server.Server()
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// config file
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh0st.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	log.Infof("gh0st server (version %s)", models.Version)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gh0std" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName(".gh0st")
	}

	// read env
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debugf("configuration file: %s", viper.ConfigFileUsed())
	}
}

// Execute initializes the cobra command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("cannot execute cobra command")
		os.Exit(1)
	}
}
