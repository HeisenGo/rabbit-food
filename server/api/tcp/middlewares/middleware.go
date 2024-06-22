package middleware

import (
	"server/api/tcp/handlers"
)

type Middleware func(handlers.Handler) handlers.Handler

func ApplyMiddleware(h handlers.Handler, middlewares ...Middleware) handlers.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
