package decorator

type Logger interface {
	Log(message string)
}

type LogDecorator struct {
	inner Logger
}

func NewLogDecorator(inner Logger) LogDecorator {
	return LogDecorator{
		inner: inner,
	}
}
