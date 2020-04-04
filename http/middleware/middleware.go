package middleware

import "net/http"

// Middleware defines a Handler wrapper.
type Middleware func(http.Handler) http.Handler

// Stack provides the method to wrap a Handler.
type Stack []func(http.Handler) http.Handler

// Apply wraps the Handler by the Stack of another handlers.
func (stack Stack) Apply(handler http.Handler) http.Handler {
	for i := len(stack) - 1; i >= 0; i-- {
		handler = stack[i](handler)
	}
	return handler
}
