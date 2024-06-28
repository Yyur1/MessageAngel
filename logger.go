package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger = logrus.New()

// SetLogger can set output log file and basic log level.
func SetLogger(logFile string, level logrus.Level) {
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			logrus.Fatal("Cannot open or create %v file：%v", logFile, err)
		}
		logger.Out = file
	} else {
		logger.Out = os.Stdout
	}
	logger.SetLevel(level)
}
