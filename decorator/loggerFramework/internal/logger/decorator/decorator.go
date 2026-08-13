package decorator

import "loggerframework/internal/logger"

type Decorator struct {
	inner logger.Logger
}

func NewDecorator(inner logger.Logger) Decorator {
	return Decorator{
		inner: inner,
	}
}
