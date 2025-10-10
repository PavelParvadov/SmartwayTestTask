package app

import (
	"context"
	"strconv"

	httpapp "github.com/pavelParvadov/SmartwayTask/internal/app/http"
	"github.com/pavelParvadov/SmartwayTask/internal/config"
	deliveryhttp "github.com/pavelParvadov/SmartwayTask/internal/delivery/http"
	"github.com/pavelParvadov/SmartwayTask/internal/repository/postgres"
	"github.com/pavelParvadov/SmartwayTask/internal/services"
	"github.com/pavelParvadov/SmartwayTask/pkg/logs"
)

type App struct {
	HTTPServer *httpapp.App
}

func NewApp() *App {
	cfg := config.GetConfig()
	logger := logs.GetNewLogger(cfg.Env)

	db, err := postgres.NewDB(context.Background(), &config.DB{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
	})

	if err != nil {
		panic(err)
	}

	repo := postgres.NewEmployeeRepositoryImpl(db)

	svc := services.NewEmployeeService(logger, repo, repo, repo, repo)

	handler := deliveryhttp.NewHandler(svc)
	httpSrv := httpapp.NewApp(
		strconv.Itoa(cfg.App.Port),
		cfg.App.Host,
		handler,
	)

	return &App{HTTPServer: httpSrv}
}
