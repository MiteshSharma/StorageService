package logs

import (
	"github.com/Sirupsen/logrus"
	"log"
	"os"
)

type StorageLogger struct {
	Name string
	Log  *logrus.Logger
}

var Logger *StorageLogger

func NewStorageLogger(name string) *StorageLogger {
	var logger = StorageLogger{Name: name}
	logger.Log = logrus.New()
	logger.Log.Level = logrus.DebugLevel
	return &logger
}

func (l *StorageLogger) Debug(msg string) {
	l.Log.Debug(msg)
}

func (l *StorageLogger) Info(msg string) {
	l.Log.Info(msg)
}

func (l *StorageLogger) Error(msg string) {
	l.Log.Error(msg)
}