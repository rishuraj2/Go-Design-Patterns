package decorator

import "fmt"

type JsonFormatterDecorator struct {
	LogDecorator
}

func NewJsonFormatterDecorator(inner Logger) JsonFormatterDecorator {
	return JsonFormatterDecorator{
		LogDecorator: NewLogDecorator(inner),
	}
}

func (this JsonFormatterDecorator) Log(message string) {
	fmt.Printf(`{"message": "`)
	this.inner.Log(message)
	fmt.Print(`"}`)
}
