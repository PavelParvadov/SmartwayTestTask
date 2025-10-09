package logs

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
	EnvLocal       = "local"
)

func GetNewLogger(env string) *zap.Logger {
	var logLevel zapcore.Level

	switch env {
	case EnvLocal:
		logLevel = zapcore.DebugLevel
	case EnvDevelopment:
		logLevel = zapcore.DebugLevel
	case EnvProduction:
		logLevel = zapcore.InfoLevel
	default:
		logLevel = zapcore.DebugLevel
	}

	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		CallerKey:        "caller",
		MessageKey:       "msg",
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " | ",
	})

	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel)
	return zap.New(core, zap.AddCaller())
}
