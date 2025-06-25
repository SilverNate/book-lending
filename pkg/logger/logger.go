package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var sharedLogger *logrus.Logger

func NewLogger() *logrus.Logger {
	if sharedLogger != nil {
		return sharedLogger
	}

	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	sharedLogger = log
	return sharedLogger
}

func InfoWithUser(module, action string, userID int, format string, args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"module":  module,
		"action":  action,
		"user_id": userID,
	}).Infof(format, args...)
}

func Info(module string, action string, format string, args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"module": module,
		"action": action,
	}).Infof(format, args...)
}

func Error(module string, action string, format string, args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"module": module,
		"action": action,
	}).Errorf(format, args...)
}
