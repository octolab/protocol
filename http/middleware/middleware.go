package middleware

import "net/http"

// Middleware defines the Handler wrapper.
type Middleware func(http.Handler) http.Handler

type Stack []Middleware

func (stack Stack) Apply(handler http.Handler) http.Handler {
	for i := len(stack) - 1; i >= 0; i-- {
		handler = stack[i](handler)
	}
	return handler
}
