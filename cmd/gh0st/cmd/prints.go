package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	errorMsg = color.RedString("[ERROR]")
	infoMsg  = color.BlueString("[INFO ]")
	okMsg    = color.HiGreenString("[ OK  ]")
)

func printErr(e error) {
	fmt.Printf(" %s %s\n", errorMsg, e)
	os.Exit(1)
}

func printOK(msg string) {
	fmt.Printf(" %s %s\n", okMsg, msg)
}

func printInfo(msg string) {
	fmt.Printf(" %s %s\n", infoMsg, msg)
}
