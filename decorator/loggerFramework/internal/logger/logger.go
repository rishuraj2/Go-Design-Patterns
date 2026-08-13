package logger

import "loggerframework/internal/model"

type Logger interface {
	Log(level model.LogLevel, message string)
}
