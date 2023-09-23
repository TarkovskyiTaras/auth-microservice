package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func logFields(handler string) logrus.Fields {
	return logrus.Fields{
		"handler": handler,
	}
}

func LogError(handler string, err error) {
	logrus.WithFields(logFields(handler)).Error(err)
}

func InitLogParams() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func Info(format string, args ...interface{}) {
	logrus.Info(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}
