package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func hello(value string) (result string) {
	return fmt.Sprintf("Hello %s", value)
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "says hello",
	Long:  "says hello to value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Output: %s", hello(args[0]))
	},
}
