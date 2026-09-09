package logger

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger
var countTest int32

const (
	TRACE_ID     = "trace-id"
	SERVICE      = "service"
	SERVICE_NAME = "notication-service"
	LOG_LEVEL    = "info"
)

func GetUUID() string {
	return uuid.New().String()
}

func CreateLoggerWithCtx(ctx context.Context) *zap.SugaredLogger {
	if logger == nil {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		config := zap.Config{
			Level:             GetLevel(),
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: false,
			Sampling:          nil,
			Encoding:          "json",
			EncoderConfig:     encoderCfg,
			OutputPaths: []string{
				"stderr",
			},
			ErrorOutputPaths: []string{
				"stderr",
			},
		}
		logger = zap.Must(config.Build()).Sugar()
		countTest++
		logger.Debugf("Created logger, count : %d", countTest)
	}
	if ctx != nil && ctx.Value(TRACE_ID).(string) != "" {
		traceId := ctx.Value(TRACE_ID)
		return logger.WithOptions(zap.Fields(zap.String(string(TRACE_ID), traceId.(string)), zap.String(SERVICE, SERVICE_NAME)))
	}
	return logger
}
func CreateLogger() *zap.SugaredLogger {
	if logger == nil {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		config := zap.Config{
			Level:             GetLevel(),
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: false,
			Sampling:          nil,
			Encoding:          "json",
			EncoderConfig:     encoderCfg,
			OutputPaths: []string{
				"stderr",
			},
			ErrorOutputPaths: []string{
				"stderr",
			},
		}
		logger = zap.Must(config.Build()).Sugar()
		countTest++
		logger.Infof("Created logger, count : %d", countTest)
	}
	return logger
}

func GetLoggerWithCtx(ctx context.Context) *zap.SugaredLogger {
	if logger == nil {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		config := zap.Config{
			Level:             GetLevel(),
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: false,
			Sampling:          nil,
			Encoding:          "json",
			EncoderConfig:     encoderCfg,
			OutputPaths: []string{
				"stderr",
			},
			ErrorOutputPaths: []string{
				"stderr",
			},
		}
		logger = zap.Must(config.Build()).Sugar()
		countTest++
		logger.Infof("Created logger, count : %d", countTest)
	}
	if ctx != nil {
		traceId := ctx.Value("trace-id")
		logger = logger.With(TRACE_ID, traceId.(string))
	}
	return logger
}
func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		config := zap.Config{
			Level:             GetLevel(),
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: false,
			Sampling:          nil,
			Encoding:          "json",
			EncoderConfig:     encoderCfg,
			OutputPaths: []string{
				"stderr",
			},
			ErrorOutputPaths: []string{
				"stderr",
			},
		}
		logger = zap.Must(config.Build()).Sugar()
		countTest++
		logger.Infof("Created logger, count : %d", countTest)
	}
	return logger
}

func GetTraceId(ctx context.Context) zap.Field {
	traceId := ctx.Value("trace-id")
	return zap.String("trace-id", traceId.(string))
}

func GetLevel() zap.AtomicLevel {

	switch LOG_LEVEL {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "dpanic":
		return zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "panic":
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	}
	return zap.NewAtomicLevelAt(zap.InfoLevel)
}
