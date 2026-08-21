package main

import (
	"loggingframework/internal/enum"
	"loggingframework/internal/logMessage"
	"loggingframework/internal/logger"
)

func main() {
	msg := logmessage.NewLogMessage(enum.ERROR, "Hello World")

	debugLogger := logger.NewDebugLogger()
	infoLogger := logger.NewInfoLogger()
	warnLogger := logger.NewWarnLogger()
	errorLogger := logger.NewErrorLogger()

	debugLogger.SetNext(&infoLogger)
	infoLogger.SetNext(&warnLogger)
	warnLogger.SetNext(&errorLogger)

	debugLogger.Log(msg)

}
