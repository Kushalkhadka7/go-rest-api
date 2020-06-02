// Package logger provides structured logging with logrus.
package logger

import (
	"github.com/sirupsen/logrus"
)

var (
	// Logger is a configured logrus.Logger.
	Logger *logrus.Logger
)

// NewLogger creates and configures a new logrus Logger.
func NewLogger() *logrus.Logger {
	Logger = logrus.New()

	// var file, err = os.OpenFile("logs/error.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	// if err != nil {
	// 	fmt.Println("Could Not Open Log File : " + err.Error())
	// }

	// defer file.Close()

	Logger.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}

	// Output to stderr instead of stdout, could also be a file.

	// Logger.SetOutput(io.(file, os.Stderr))

	Logger.SetLevel(logrus.PanicLevel)
	Logger.SetLevel(logrus.FatalLevel)
	Logger.SetLevel(logrus.ErrorLevel)
	Logger.SetLevel(logrus.WarnLevel)
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetLevel(logrus.DebugLevel)

	return Logger
}
