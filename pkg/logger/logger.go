// Provides a structured, leveled logging
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var err error
	log, err = config.Build()
	if err != nil {
		panic(err)
	}
}

// Get returns a initialized zap logger object
func Get() *zap.Logger {
	return log
}

// Info will use zap logger in Info level
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Debug will use zap logger in Debug level
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error will use zap logger in Error level
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

// Warn will use zap logger in Warn level
func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}

// Panic will use zap logger in Panic level
func Panic(message string, fields ...zap.Field) {
	log.Panic(message, fields...)
}

// Fatal will use zap logger in Fatal level
func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}
