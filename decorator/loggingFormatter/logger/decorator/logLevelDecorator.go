package decorator

import "fmt"

type LogLevelDecorator struct {
	LogDecorator
}

func NewLogLevelDecorator(inner Logger) LogLevelDecorator {
	return LogLevelDecorator{
		LogDecorator: NewLogDecorator(inner),
	}
}

func (this LogLevelDecorator) Log(message string) {
	fmt.Printf("[Level] ")
	this.inner.Log(message)
}
