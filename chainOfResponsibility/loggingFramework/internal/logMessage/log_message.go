package logmessage

import "loggingframework/internal/enum"

type LogMessage struct {
	level   enum.LogLevel
	message string
}

func NewLogMessage(lvl enum.LogLevel, msg string) LogMessage {
	return LogMessage{
		level:   lvl,
		message: msg,
	}
}

func (this LogMessage) GetLevel() enum.LogLevel {
	return this.level
}

func (this LogMessage) GetMessage() string {
	return this.message
}
