package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/protocol/http/middleware"
)

func TestStack_Apply(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		var (
			calls []int
			stack Stack
		)
		handler := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
			calls = append(calls, 1)
		})

		stack.Apply(handler).ServeHTTP(httptest.NewRecorder(), &http.Request{})
		assert.Equal(t, []int{1}, calls)
	})
	t.Run("full stack", func(t *testing.T) {
		var calls []int
		stack := Stack{
			func(handler http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					calls = append(calls, 1)
					handler.ServeHTTP(rw, req)
				})
			},
			func(handler http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					calls = append(calls, 2)
					handler.ServeHTTP(rw, req)
				})
			},
		}
		handler := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
			calls = append(calls, 3)
		})

		stack.Apply(handler).ServeHTTP(httptest.NewRecorder(), &http.Request{})
		assert.Equal(t, []int{1, 2, 3}, calls)
	})
}
