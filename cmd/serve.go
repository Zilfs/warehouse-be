package cmd

import (
	"fmt"
	"warehouse/internal/app"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server",
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetString("APP_PORT")
		fmt.Println("Starting server on port", port)
		app.RunServer()
	},
}
