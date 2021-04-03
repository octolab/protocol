package middleware

import (
	"net/http"

	"github.com/kamilsk/tracer"
)

// Tracer returns the Middleware to inject a simple tracer
// into the request context.
func Tracer(buffer int, logger func(*tracer.Trace)) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := tracer.Inject(req.Context(), make([]*tracer.Call, 0, buffer))
			handler.ServeHTTP(rw, req.WithContext(ctx))

			select {
			case <-req.Context().Done():
				logger(tracer.Fetch(ctx))
			default:
				return
			}
		})
	}
}
