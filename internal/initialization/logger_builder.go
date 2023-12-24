package initialization

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level string) *zap.Logger {
	config := zap.Config{
		Level:       loggerLvl(level),
		Encoding:    "console",
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder, // or zapcore.LowercaseColorLevelEncoder for colored output
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, err := config.Build()
	if err != nil {
		log.Fatalln(err)
	}

	return logger
}

func loggerLvl(level string) zap.AtomicLevel {
	switch level {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "panic":
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	case "fatal":
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}
