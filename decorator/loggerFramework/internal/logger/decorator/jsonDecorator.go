package decorator

import (
	"fmt"
	"loggerframework/internal/logger"
	"loggerframework/internal/model"
	"strings"
	"time"
)

type JsonDecorator struct {
	Decorator
}

func NewJsonDecorator(inner logger.Logger) JsonDecorator {
	return JsonDecorator{
		Decorator: NewDecorator(inner),
	}
}

func (this JsonDecorator) Log(level model.LogLevel, message string) {
	fmt.Printf(`{"timestamp": %s, "level": %s, "message": "`, time.Now().Format("2006-01-02 15:04:05"), strings.ToUpper(level.String()))
	this.inner.Log(level, message)
	fmt.Println(`"}`)
}
