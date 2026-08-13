package middleware

import "requestprocessor/model"

type RequestHandler interface {
	SetNext(next RequestHandler)
	Handle(req model.Request)
}
