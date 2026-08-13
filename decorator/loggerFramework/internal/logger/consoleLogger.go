package logger

import (
	"fmt"
	"loggerframework/internal/model"
	"strings"
)

type ConsoleLogger struct{}

func NewConsoleLogger() ConsoleLogger {
	return ConsoleLogger{}
}

func (this ConsoleLogger) Log(level model.LogLevel, message string) {
	fmt.Printf("%s: %s", strings.ToUpper(level.String()), message)
}
