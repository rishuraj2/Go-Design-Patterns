package logger

import "fmt"

type SimpleLogger struct{}

func NewSimpleLogger() SimpleLogger {
	return SimpleLogger{}
}

func (this SimpleLogger) Log(message string) {
	fmt.Print(message)
}
