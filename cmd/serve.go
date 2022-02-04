package cmd

import (
	"github.com/analog-substance/lithium/lib"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Long: `Listens on port 8000`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.Serve(8000)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
