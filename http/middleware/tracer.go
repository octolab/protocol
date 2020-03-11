package middleware

import (
	"net/http"

	"github.com/kamilsk/tracer"
)

// Tracer returns the Middleware to inject a simple tracer
// into the request context.
func Tracer(buffer int) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			handler.ServeHTTP(rw, req.WithContext(
				tracer.Inject(req.Context(), make([]*tracer.Call, 0, buffer)),
			))

			select {
			case <-req.Context().Done():
				// TODO:log
			default:
				return
			}
		})
	}
}
