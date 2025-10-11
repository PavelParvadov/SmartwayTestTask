// @title Employee Service API
// @version 1.0
// @description API for managing employees
// @host localhost:8081
// @BasePath /
// @schemes http
package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/pavelParvadov/SmartwayTask/docs"
	"github.com/pavelParvadov/SmartwayTask/internal/app"
)

func main() {
	application := app.NewApp()
	go application.HTTPServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	application.HTTPServer.Stop()
	application.Logger.Info("application stopped")
}
