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
	middleware := Tracer(1, func(tracer *tracer.Trace) {
		assert.Contains(t, tracer.String(), "call middleware_test.TestTracer.func2")
		spy = true
	})

	var handler http.Handler
	handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		defer tracer.Fetch(req.Context()).Start().Stop()
		time.Sleep(time.Millisecond)
	})
	handler = middleware(handler)

	ctx, cancel := context.WithCancel(context.Background())
	handler.ServeHTTP(httptest.NewRecorder(), (&http.Request{}).WithContext(ctx))
	assert.False(t, spy)

	cancel()
	handler.ServeHTTP(httptest.NewRecorder(), (&http.Request{}).WithContext(ctx))
	assert.True(t, spy)
}
