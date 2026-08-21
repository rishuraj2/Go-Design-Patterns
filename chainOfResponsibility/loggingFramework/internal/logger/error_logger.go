package logger

import (
	"fmt"
	logmessage "loggingframework/internal/logMessage"
)

type ErrorLogger struct {
	BaseLogger
}

func NewErrorLogger() ErrorLogger {
	return ErrorLogger{
		BaseLogger: NewBaseLogger(),
	}
}

func (this ErrorLogger) Log(msg logmessage.LogMessage) {
	if msg.GetLevel() >= 3 {
		fmt.Printf("[ERROR] %s\n", msg.GetMessage())
	}
	this.forward(msg)
}
