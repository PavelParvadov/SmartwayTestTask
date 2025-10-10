// @title Employee Service API
// @version 1.0
// @description API for managing employees
// @host localhost:8081
// @BasePath /
// @schemes http
package main

import (
	_ "github.com/pavelParvadov/SmartwayTask/docs"
	"github.com/pavelParvadov/SmartwayTask/internal/app"
	"github.com/pavelParvadov/SmartwayTask/internal/config"
	"github.com/pavelParvadov/SmartwayTask/pkg/logs"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.GetConfig()
	log := logs.GetNewLogger(cfg.Env)

	application := app.NewApp()
	go application.HTTPServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	application.HTTPServer.Stop()
	log.Info("application stopped")
}
