package cmd

import (
	"rate-limiter/pkg/repository/mongo"
	"rate-limiter/pkg/server"
	"rate-limiter/pkg/service"

	"github.com/labstack/echo"
	"github.com/spf13/cobra"
	"github.com/vrischmann/envconfig"
	"go.uber.org/zap"
)

func serve(args []string) {
	e := echo.New()

	logger, _ := zap.NewDevelopment()

	var mongoConf mongo.Config
	if envconfig.Init(&mongoConf) != nil {
		logger.Fatal("Failed to load mongo config")
	}

	mongoRepo, err := mongo.NewRepository(mongoConf, logger)
	if err != nil {
		logger.Fatal("Failed to create mongo repository")
	}

	service, err := service.NewService(logger, mongoRepo)
	if err != nil {
		logger.Fatal("Failed to create service")
	}

	server := server.NewServer(service)
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
