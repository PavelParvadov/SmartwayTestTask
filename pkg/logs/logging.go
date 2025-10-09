package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

type Env string

const (
	EnvProduction  Env = "production"
	EnvDevelopment Env = "development"
	EnvLocal       Env = "local"
)

func ParseEnv(s string) Env {
	switch Env(strings.ToLower(strings.TrimSpace(s))) {
	case EnvProduction:
		return EnvProduction
	case EnvDevelopment:
		return EnvDevelopment
	case EnvLocal:
		return EnvLocal
	default:
		return EnvLocal
	}
}

func GetNewLogger(envStr string) *zap.Logger {
	env := ParseEnv(envStr)

	var enc zapcore.Encoder
	var lvl zapcore.Level
	var stacktraceAt zapcore.Level

	switch env {
	case EnvProduction:
		cfg := zap.NewProductionEncoderConfig()
		cfg.TimeKey = "time"
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		enc = zapcore.NewJSONEncoder(cfg)
		lvl = zapcore.InfoLevel
		stacktraceAt = zapcore.ErrorLevel
	case EnvDevelopment, EnvLocal:
		cfg := zap.NewDevelopmentEncoderConfig()
		cfg.TimeKey = "time"
		cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
		enc = zapcore.NewConsoleEncoder(cfg)
		lvl = zapcore.DebugLevel
		stacktraceAt = zapcore.ErrorLevel
	}

	core := zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), lvl)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(stacktraceAt))
}
