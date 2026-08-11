package main

import (
	"fmt"
	"loggingformatter/logger"
	"loggingformatter/logger/decorator"
)

func main() {
	l := decorator.NewLogLevelDecorator(decorator.NewJsonFormatterDecorator(decorator.NewUpperCaseDecorator(logger.NewSimpleLogger())))
	l.Log("Hello")
	fmt.Println()
}
