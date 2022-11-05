package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger    *zap.SugaredLogger
	ZapLogger *zap.Logger
)

func init() {
	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConf.LevelKey = "log_level"
	encoderConf.MessageKey = "message"
	encoderConf.StacktraceKey = "stacktrace"
	encoderConf.TimeKey = "timestamp_app"

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig:    encoderConf,
		DisableCaller:    true,
	}

	var err error
	ZapLogger, err = cfg.Build()

	if err != nil {
		panic(err)
	}

	logger = ZapLogger.Sugar()
	defer logger.Sync()
}

func Debug(message string, ctx ...*LogContext) {
	logger.Debugw(message, zap.Any("event", ctx))
}

func Info(message string, ctx ...LogContext) {
	logger.Infow(message, zap.Any("event", ctx))
	NewRelicLogger.Info(message, ctx)
}

func Warn(message string, ctx ...*LogContext) {
	logger.Warnw(message, zap.Any("event", ctx))
}

func Error(message string, err error, ctx ...LogContext) {
	logger.Errorw(message, zap.Error(err), zap.Any("event", ctx))
	NewRelicLogger.Error(message, err, ctx)
}

func Fatal(message string, err error, ctx ...*LogContext) {
	logger.Fatalw(message, zap.Error(err), zap.Any("event", ctx))
}
