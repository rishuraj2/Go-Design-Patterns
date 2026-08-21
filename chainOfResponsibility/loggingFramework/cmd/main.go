package main

import (
	"loggingframework/internal/enum"
	logmessage "loggingframework/internal/logMessage"
	"loggingframework/internal/logger"
	"loggingframework/internal/logger/decorator"
)

func main() {
	msg := logmessage.NewLogMessage(enum.INFO, "Hello World")

	debugLogger := logger.NewDebugLogger()
	infoLogger := logger.NewInfoLogger()
	warnLogger := logger.NewWarnLogger()
	errorLogger := logger.NewErrorLogger()

	debugLogger.SetNext(&infoLogger)
	infoLogger.SetNext(&warnLogger)
	warnLogger.SetNext(&errorLogger)

	logger := decorator.NewColorDecorator(&debugLogger)

	logger.Log(msg)

}
