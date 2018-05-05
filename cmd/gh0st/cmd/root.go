package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gpaulo00/gh0st/sdk"
	"github.com/spf13/cobra"
)

var (
	api    string
	client *sdk.Client
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh0st",
	Short: "RESTful API client for gh0st",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(start)
	rootCmd.PersistentFlags().StringVarP(&api, "server", "s", "http://localhost:8080", "gh0st server url")
}

// start runs when initializing
func start() {
	printInfo(color.CyanString("gh0st client"))

	// connect to gh0st
	c, err := sdk.New(api)
	if err != nil {
		printErr(err)
	}
	client = c
}
