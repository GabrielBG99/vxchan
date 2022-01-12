package logger

import (
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

type Level string

const (
	DebugLevel Level = "debug"
	InfoLevel  Level = "info"
	WarnLevel  Level = "warn"
	ErrorLevel Level = "error"
)

// Config logger's configuration struct
type Config struct {
	Name     string
	Level    Level
	Filepath string
}

func (c Config) getLevel() zapcore.Level {
	switch c.Level {
	case DebugLevel:
		return zap.DebugLevel
	case WarnLevel:
		return zap.WarnLevel
	case ErrorLevel:
		return zap.ErrorLevel
	}

	return zap.InfoLevel
}

// Init creates a new logger instance
func Init(config Config) (*zap.SugaredLogger, error) {
	if logger != nil {
		return logger, nil
	}

	var w zapcore.WriteSyncer = os.Stdout
	if config.Filepath != "" {
		w = zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.Filepath,
			MaxSize:    100,
			MaxBackups: 0,
		})
	}
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, w, config.getLevel())

	log := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.PanicLevel),
	)

	logger = log.Sugar()

	return logger, nil
}

// GetLogger returns the default logger instance
func GetLogger() *zap.SugaredLogger {
	return logger
}
