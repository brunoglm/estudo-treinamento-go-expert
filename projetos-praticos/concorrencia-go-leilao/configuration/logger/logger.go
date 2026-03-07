package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logConfiguration := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	log, err = logConfiguration.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	Sync()
}

func Infof(template string, args ...any) {
	log.Info(fmt.Sprintf(template, args...))
	Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	Sync()
}

func Errorf(template string, args ...any) {
	log.Error(fmt.Sprintf(template, args...))
	Sync()
}

func Sync() {
	if err := log.Sync(); err != nil {
		fmt.Printf("Failed to sync logger: %v\n", err)
	}
}
