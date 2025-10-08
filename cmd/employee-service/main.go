package main

import (
	"github.com/pavelParvadov/SmartwayTask/pkg/logs"
	"go.uber.org/zap"
)

func main() {
	logger := logs.GetNewLogger("local")
	logger.Info("test logger", zap.Any("logger", logger))
}
