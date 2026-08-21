package decorator

import (
	"fmt"
	"loggingframework/internal/enum"
	logmessage "loggingframework/internal/logMessage"
	"loggingframework/internal/logger"
)

type ColorDecorator struct {
	inner logger.Logger
}

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func NewColorDecorator(inner logger.Logger) ColorDecorator {
	return ColorDecorator{
		inner: inner,
	}
}

func (this *ColorDecorator) Log(msg logmessage.LogMessage) {
	switch msg.GetLevel() {
	case enum.DEBUG:
		fmt.Print(Green)

	case enum.INFO:
		fmt.Print(Green)

	case enum.WARN:
		fmt.Print(Yellow)

	case enum.ERROR:
		fmt.Print(Red)

	default:
	}

	this.inner.Log(msg)
	fmt.Print(Reset)
}
