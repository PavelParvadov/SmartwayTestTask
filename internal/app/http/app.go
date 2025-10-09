package http

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	delivery "github.com/pavelParvadov/SmartwayTask/internal/delivery/http"
)

type App struct {
	router *fiber.App
	port   string
	host   string
}

func NewApp(port, host string, handlers ...*delivery.Handler) *App {
	router := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})
	router.Use(logger.New(), cors.New())

	for _, h := range handlers {
		h.Register(router)
	}

	return &App{router: router, port: port, host: host}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	if err := a.router.Listen(a.host + ":" + a.port); err != nil {
		return err
	}
	return nil
}

func (a *App) Stop() {
	if err := a.router.Shutdown(); err != nil {
		panic(err)
	}
}
