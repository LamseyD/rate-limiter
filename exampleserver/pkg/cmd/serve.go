package cmd

import (
	"exampleserver/pkg/server"
	"exampleserver/pkg/service"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/vrischmann/envconfig"
	"go.uber.org/zap"
)

func serve(args []string) {
	logger, _ := zap.NewDevelopment()

	service, err := service.NewService(logger)
	if err != nil {
		logger.Fatal("Failed to create service")
	}

	var serverConfig server.Config
	if envconfig.Init(&serverConfig) != nil {
		logger.Fatal("Failed to load server config")
	}
	server := server.NewServer(service, logger, serverConfig)

	echo := echo.New()
	server.StartServer(echo)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts an echo server",
	Run: func(cmd *cobra.Command, args []string) {
		serve(args)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
