package decorator

import "fmt"

type TimeStampDecorator struct {
	LogDecorator
}

func NewTimeStampDecorator(inner Logger) TimeStampDecorator {
	return TimeStampDecorator{
		LogDecorator: NewLogDecorator(inner),
	}
}

func (this TimeStampDecorator) Log(message string) {
	fmt.Print("[2024-01-15 10:30:00] ")
	this.inner.Log(message)
}
