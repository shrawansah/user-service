package logger

import (
	configService "gojek.com/config"
)

func Info(args ...interface{}) {
	var logger = getLogger()
	logger.Info(args)
}

func Error(args ...interface{}) {
	var logger = getLogger()
	logger.Error(args)
}

/*
 definitions
*/
type loggers interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

var consoleLogger = consoleLoggerImpl{}
var cloudwatchLogger = cloudwatchLoggerImpl{}

// factory function for logger objects
func getLogger() loggers {

	environment := configService.Configs.Environment
	switch environment {
	case "development":
		return &consoleLogger
	case "test":
		return &cloudwatchLogger
	default:
		return &consoleLogger
	}
}
