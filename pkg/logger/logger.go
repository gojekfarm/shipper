package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"source.golabs.io/core/shipper/pkg/config"
)

// Logger - structure to hold the logrus logger
type Logger struct {
	*logrus.Logger
}

// NewLogger - create a new logrus logger
func NewLogger(config *config.Config) *Logger {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &Logger{
		&logrus.Logger{
			Out:       os.Stdout,
			Hooks:     make(logrus.LevelHooks),
			Level:     level,
			Formatter: &logrus.TextFormatter{},
		},
	}
}
