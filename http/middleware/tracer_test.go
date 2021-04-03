package middleware_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kamilsk/tracer"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/protocol/http/middleware"
)

func TestTracer(t *testing.T) {
	var spy bool
	logger := func(tracer *tracer.Trace) {
		assert.Contains(t, tracer.String(), "call middleware_test.TestTracer.func2")
		spy = true
	}
	middleware := Tracer(1, logger)

	var handler http.Handler
	handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		defer tracer.Fetch(req.Context()).Start().Stop()
		time.Sleep(time.Millisecond)
	})
	handler = middleware(handler)

	t.Run("without log", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		handler.ServeHTTP(httptest.NewRecorder(), (&http.Request{}).WithContext(ctx))
		assert.False(t, spy)
	})
	t.Run("with log", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		handler.ServeHTTP(httptest.NewRecorder(), (&http.Request{}).WithContext(ctx))
		assert.True(t, spy)
	})
}
