package decorator

import (
	"fmt"
	"loggerframework/internal/logger"
	"loggerframework/internal/model"
)

type ColorDecorator struct {
	Decorator
}

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func NewColorDecorator(inner logger.Logger) ColorDecorator {
	return ColorDecorator{
		Decorator: NewDecorator(inner),
	}
}

func (this ColorDecorator) Log(level model.LogLevel, message string) {
	switch level {
	case model.INFO:
		fmt.Print(Green)

	case model.WARN:
		fmt.Print(Yellow)

	case model.ERROR:
		fmt.Print(Red)

	default:
	}

	this.inner.Log(level, message)
	fmt.Print(Reset)
}
