package cmd

import (
	"fmt"
	"os"

	"github.com/gpaulo00/gh0st/external/nmap"

	"github.com/spf13/cobra"
)

var workspace uint64

func init() {
	rootCmd.AddCommand(nmapCmd)
	nmapCmd.Flags().Uint64VarP(&workspace, "workspace", "w", 1, "workspace to contain data")
	nmapCmd.MarkFlagRequired("workspace")
}

var nmapCmd = &cobra.Command{
	Use:   "nmap",
	Short: "Import a Nmap XML scan",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		printInfo(fmt.Sprintf("importing %d nmap scans", len(args)))

		for i := range args {
			// open file
			fp, err := os.Open(args[i])
			if err != nil {
				printErr(err)
			}
			defer fp.Close()

			// parse nmap scan
			imp, err := nmap.New().Parse(fp)
			if err != nil {
				printErr(err)
			}

			// import to gh0st
			res, err := client.Import(imp.Import(workspace))
			if err != nil {
				printErr(err)
			}
			printOK(res.String())
		}

		printOK("finish nmap import")
	},
}
