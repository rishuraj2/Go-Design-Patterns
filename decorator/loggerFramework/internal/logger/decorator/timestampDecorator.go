package decorator

import (
	"fmt"
	"loggerframework/internal/logger"
	"loggerframework/internal/model"
	"time"
)

type TimestampDecorator struct {
	Decorator
}

func NewTimestampDecorator(inner logger.Logger) TimestampDecorator {
	return TimestampDecorator{
		Decorator: NewDecorator(inner),
	}
}

func (this TimestampDecorator) Log(level model.LogLevel, message string) {
	fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
	this.inner.Log(level, message)
}
