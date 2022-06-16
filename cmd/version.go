package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Current version of the api server",
	Long:  "This command will print the current version of the api server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v2.0.1")
	},
}
