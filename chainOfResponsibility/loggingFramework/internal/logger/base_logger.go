package logger

import (
	logmessage "loggingframework/internal/logMessage"
)

type Logger interface {
	Log(msg logmessage.LogMessage)
}

type BaseLogger struct {
	next Logger
}

func NewBaseLogger() BaseLogger {
	return BaseLogger{}
}

func (this *BaseLogger) SetNext(next Logger) {
	this.next = next
}

func (this *BaseLogger) Forward(msg logmessage.LogMessage) {
	if this.next != nil {
		this.next.Log(msg)
	}
}
