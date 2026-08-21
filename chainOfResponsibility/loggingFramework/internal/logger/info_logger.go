package logger

import (
	"fmt"
	logmessage "loggingframework/internal/logMessage"
)

type InfoLogger struct {
	BaseLogger
}

func NewInfoLogger() InfoLogger {
	return InfoLogger{
		BaseLogger: NewBaseLogger(),
	}
}

func (this InfoLogger) Log(msg logmessage.LogMessage) {
	if msg.GetLevel() >= 1 {
		fmt.Printf("[INFO] %s\n", msg.GetMessage())
	}
	this.forward(msg)
}
