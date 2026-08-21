package logger

import (
	"fmt"
	logmessage "loggingframework/internal/logMessage"
)

type DebugLogger struct {
	BaseLogger
}

func NewDebugLogger() DebugLogger {
	return DebugLogger{
		BaseLogger: NewBaseLogger(),
	}
}

func (this DebugLogger) Log(msg logmessage.LogMessage) {
	if msg.GetLevel() >= 0 {
		fmt.Printf("[DEBUG] %s\n", msg.GetMessage())
	}
	this.forward(msg)
}
