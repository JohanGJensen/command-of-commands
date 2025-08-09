package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "coc",
	Short: "'command of commands' is a cli tool for listing package.json script commands.",
	Long:  "'command of commands' is a cli tool for listing and performing package.json script commands without having parsing the file or opening an editor.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	root.AddCommand(helloCmd)
	root.AddCommand(scriptsCmd)
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error while executing coc '%s'\n", err)
		os.Exit(1)
	}
}
