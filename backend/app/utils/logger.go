package utils

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

var (
	logModes = map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"error": logrus.ErrorLevel,
	}
	defaultLogMode = logrus.ErrorLevel
)

func LoggerInit() {
	var LogLevelStr = strings.ToLower(os.Getenv("LOGLEVEL"))
	var logLevel, found = logModes[LogLevelStr]
	if !found {
		logLevel = defaultLogMode
	}
	Log = logrus.New()
	Log.SetLevel(logLevel)
}
