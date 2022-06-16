package cmd

import (
	"github.com/pkbhowmick/go-rest-api/api"
	"github.com/spf13/cobra"
)

var port string
var auth bool

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "This flag sets the port of the server")
	startCmd.PersistentFlags().BoolVarP(&auth, "auth", "a", true, "This flag will impose/bypass authentication to API server")
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "This command will start the api server",
	Long:  "This command will start the go-rest-api server",
	Run: func(cmd *cobra.Command, args []string) {
		api.SetFlags(port, auth)
		api.StartServer()
	},
}
