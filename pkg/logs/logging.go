package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func GetNewLogger(env string) *zap.Logger {
	cfg := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		CallerKey:        "caller",
		MessageKey:       "msg",
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " | ",
	})

	var logLevel zapcore.Level

	switch env {
	case "production":
		logLevel = zap.InfoLevel
	case "development":
		logLevel = zap.DebugLevel
	case "local":
		logLevel = zap.DebugLevel
	}

	core := zapcore.NewCore(cfg, zapcore.AddSync(os.Stdout), logLevel)

	logger := zap.New(core, zap.AddCaller())

	return logger
}
