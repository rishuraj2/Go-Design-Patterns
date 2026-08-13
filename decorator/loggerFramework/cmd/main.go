package main

import (
	"loggerframework/internal/logger"
	"loggerframework/internal/logger/decorator"
	"loggerframework/internal/model"
)

func main() {
	text := decorator.NewColorDecorator(decorator.NewJsonDecorator(logger.NewConsoleLogger()))
	text.Log(model.WARN, "Hello")
	text.Log(model.INFO, "Hello")
}