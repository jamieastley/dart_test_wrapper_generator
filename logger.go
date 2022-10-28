package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
}

type ConsoleLogger struct {
	logger *zap.SugaredLogger
}

func NewConsoleLogger(enableVerbose bool) *ConsoleLogger {
	logLevel := zapcore.InfoLevel
	if enableVerbose {
		logLevel = zapcore.DebugLevel
	}

	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), logLevel),
	)
	z := zap.New(core, zap.AddCaller())

	return &ConsoleLogger{logger: z.Sugar()}
}

func (c *ConsoleLogger) Info(args ...interface{}) {
	c.logger.Info(args)
}

func (c *ConsoleLogger) Debug(args ...interface{}) {
	c.logger.Debug(args)
}

func (c *ConsoleLogger) Error(args ...interface{}) {
	c.logger.Error(args)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
