package logger

import (
	"fmt"
	logmessage "loggingframework/internal/logMessage"
)

type WarnLogger struct {
	BaseLogger
}

func NewWarnLogger() WarnLogger {
	return WarnLogger{
		BaseLogger: NewBaseLogger(),
	}
}

func (this WarnLogger) Log(msg logmessage.LogMessage) {
	if msg.GetLevel() >= 2 {
		fmt.Printf("[WARN] %s\n", msg.GetMessage())
	}
	this.forward(msg)
}
