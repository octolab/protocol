package chi_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/toolkit/protocol/http/middleware/chi"
)

func TestConvertPlaceholders(t *testing.T) {
	echo := func(rw http.ResponseWriter, req *http.Request) {
		_, err := io.Copy(rw, strings.NewReader(req.URL.Query().Get("message")))
		require.NoError(t, err)
	}

	t.Run("invalid placeholders", func(t *testing.T) {
		assert.Panics(t, func() { Converter("msg", "message", "id") })
	})
	t.Run("valid placeholders", func(t *testing.T) {
		router := chi.NewRouter()
		router.With(Converter("msg", "message")).Get("/{msg}", echo)

		body := assert.HTTPBody(router.ServeHTTP, http.MethodGet, "/hello", nil)
		assert.Equal(t, "hello", body)
	})
	t.Run("restful api", func(t *testing.T) {
		router := chi.NewRouter()
		router.Route("/resources", func(router chi.Router) {
			router.Route("/{id}", func(router chi.Router) {
				router.Use(Converter("id", "message"))
				router.Get("/", echo)
				router.Put("/", echo)
				router.Delete("/", echo)
			})
		})

		for _, method := range []string{http.MethodGet, http.MethodPut, http.MethodDelete} {
			body := assert.HTTPBody(router.ServeHTTP, method, "/resources/uid", nil)
			assert.Equal(t, "uid", body)
		}
	})
}
