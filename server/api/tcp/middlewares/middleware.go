package middleware

type Middleware func(handlerFunc HandlerFunc) HandlerFunc

func ApplyMiddlewares(h HandlerFunc, middlewares ...Middleware) HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
