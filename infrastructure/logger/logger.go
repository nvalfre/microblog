package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is the internal logrus instance
var Logger *logrus.Logger

// InitializeLogger configures the logger for the application
func InitializeLogger() {
	Logger = logrus.New()

	Logger.SetOutput(os.Stdout)

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "DEBUG" {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.InfoLevel)
	}

	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05Z07:00",
	})

	Logger.Info("Logger initialized successfully")
}

// Info logs an informational message
func Info(message string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		Logger.WithFields(fields[0]).Info(message)
	} else {
		Logger.Info(message)
	}
}

// Warn logs a warning message
func Warn(message string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		Logger.WithFields(fields[0]).Warn(message)
	} else {
		Logger.Warn(message)
	}
}

// Error logs an error message with an error object
func Error(message string, err error, fields ...logrus.Fields) {
	if len(fields) > 0 {
		Logger.WithFields(fields[0]).Error(message + ", error:" + err.Error())
	} else {
		Logger.Error(err.Error())
	}
}

// Debug logs a debug message
func Debug(message string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		Logger.WithFields(fields[0]).Debug(message)
	} else {
		Logger.Debug(message)
	}
}

func Fatal(message string, err error, fields ...logrus.Fields) {
	if len(fields) > 0 {
		Logger.WithFields(fields[0]).Fatal(message + ", error:" + err.Error())
	} else {
		Logger.Fatal(err.Error())
	}
}
