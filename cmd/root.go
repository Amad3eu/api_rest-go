package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   ".",
	Short: "This is the main command",
	Long:  `This is the main command of go-rest-api server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cli works")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
