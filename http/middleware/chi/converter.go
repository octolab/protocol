package chi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

// Converter returns the Middleware to convert chi' URL params into query params.
//
//  router := chi.NewRouter()
//  router.Route("/resources", func(router chi.Router) {
//  	router.With(middleware.Pagination).Get("/", resource.List)
//  	router.Post("/", resource.Create)
//  	router.Route("/{id}", func(router chi.Router) {
//  		router.Use(middleware.Converter("id", "resourceID"))
//  		router.Get("/", resource.Fetch)
//  		router.Put("/", resource.Update)
//  		router.Delete("/", resource.Delete)
//  	})
//  })
//  http.Handle("/", router)
//
func Converter(placeholders ...string) func(http.Handler) http.Handler {
	if count := len(placeholders); count == 0 || count%2 != 0 {
		panic(errors.Errorf("count of passed placeholders must be even: obtained %d", count))
	}
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			q := req.URL.Query()
			for i, count := 0, len(placeholders); i < count; i += 2 {
				from, to := placeholders[i], placeholders[i+1]
				q.Set(to, chi.URLParam(req, from))
			}
			req.URL.RawQuery = q.Encode()
			handler.ServeHTTP(rw, req)
		})
	}
}
