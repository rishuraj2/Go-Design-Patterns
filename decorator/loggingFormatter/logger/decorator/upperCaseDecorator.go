package decorator

import (
	"strings"
)

type UpperCaseDecorator struct {
	LogDecorator
}

func NewUpperCaseDecorator(inner Logger) UpperCaseDecorator {
	return UpperCaseDecorator{
		LogDecorator: NewLogDecorator(inner),
	}
}

func (this UpperCaseDecorator) Log(message string) {
	message = strings.ToUpper(message)
	this.inner.Log(message)
}
