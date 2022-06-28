package entities

import log "github.com/sirupsen/logrus"

type Logger struct {
}

func Construct() *Logger {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	logger := &Logger{}

	return logger
}

func (l *Logger) Panic(msg string) {
	log.Panic(msg)
}

func (l *Logger) Fatal(msg string) {
	log.Fatal(msg)
}

func (l *Logger) Error(msg string) {
	log.Error(msg)
}

func (l *Logger) Warn(msg string) {
	log.Warn(msg)
}

func (l *Logger) Info(msg string) {
	log.Info(msg)
}

func (l *Logger) Debug(msg string) {
	log.Debug(msg)
}

func (l *Logger) Trace(msg string) {
	log.Trace(msg)
}
