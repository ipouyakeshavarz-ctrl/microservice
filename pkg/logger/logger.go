package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var globalLogger *zap.Logger
var once sync.Once

func InitLogger(serviceName string, isProduction bool, logFilePath string) {
	once.Do(func() {

		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder

		defaultEncoder := zapcore.NewJSONEncoder(config)

		if logFilePath == "" {
			logFilePath = "./logs/log.json"
		}

		fileWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    10, // megabytes
			MaxAge:     30, // days
			MaxBackups: 10,
			LocalTime:  false,
		})

		stdOutWriter := zapcore.AddSync(os.Stdout)

		defaultLogLevel := zap.InfoLevel
		if !isProduction {
			defaultLogLevel = zap.DebugLevel
		}

		core := zapcore.NewTee(
			zapcore.NewCore(defaultEncoder, fileWriter, defaultLogLevel),
			zapcore.NewCore(defaultEncoder, stdOutWriter, defaultLogLevel),
		)

		logger := zap.New(
			core,
			zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zapcore.ErrorLevel),
		)

		globalLogger = logger.With(zap.String("service", serviceName))

		zap.ReplaceGlobals(globalLogger)
	})
}

func Info(msg string, fields ...zap.Field) {
	if globalLogger == nil {
		return
	}
	globalLogger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	if globalLogger == nil {
		return
	}
	globalLogger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	if globalLogger == nil {
		return
	}
	globalLogger.Fatal(msg, fields...)
}

func Sync() error {
	if globalLogger == nil {
		return nil
	}
	return globalLogger.Sync()
}
