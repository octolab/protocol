package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.octolab.org/toolkit/protocol/http/header"
	. "go.octolab.org/toolkit/protocol/http/middleware"
)

func TestDeadline(t *testing.T) {
	var timeout = func(timeout time.Duration) http.HandlerFunc {
		return func(rw http.ResponseWriter, req *http.Request) {
			timer := time.NewTimer(timeout)
			select {
			case <-req.Context().Done():
				assert.True(t, true)
			case <-timer.C:
				assert.True(t, false)
			}
			assert.True(t, timer.Stop())
		}
	}
	corrector := func(deadline time.Time, present bool) time.Time {
		correction := time.Millisecond
		if present && deadline.After(time.Now().Add(2*correction)) {
			return deadline.Add(-correction)
		}
		return deadline
	}

	tests := map[string]struct {
		fallback time.Duration
		request  *http.Request
		handler  http.Handler
	}{
		"exists in header": {
			fallback: time.Second,
			request: &http.Request{
				Header: http.Header{
					header.XDeadline: []string{
						time.Now().Add(10 * time.Millisecond).Format(time.RFC3339Nano),
					},
				},
			},
			handler: timeout(15 * time.Millisecond),
		},
		"fallback cause empty": {
			fallback: 10 * time.Millisecond,
			request:  &http.Request{},
			handler:  timeout(15 * time.Millisecond),
		},
		"fallback cause invalid": {
			fallback: 10 * time.Millisecond,
			request: &http.Request{
				Header: http.Header{
					header.XDeadline: []string{"bad"},
					header.XTimeout:  []string{"bad"},
				},
			},
			handler: timeout(15 * time.Millisecond),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			middleware := Deadline(test.fallback, corrector)
			recorder := httptest.NewRecorder()
			middleware(test.handler).ServeHTTP(recorder, test.request)
		})
	}
}

func TestTimeout(t *testing.T) {
	var timeout = func(timeout time.Duration) http.HandlerFunc {
		return func(rw http.ResponseWriter, req *http.Request) {
			timer := time.NewTimer(timeout)
			select {
			case <-req.Context().Done():
				assert.True(t, true)
			case <-timer.C:
				assert.True(t, false)
			}
			assert.True(t, timer.Stop())
		}
	}
	corrector := func(timeout time.Duration, present bool) time.Duration {
		correction := time.Millisecond
		if present && timeout > 2*correction {
			return timeout - correction
		}
		return timeout
	}

	tests := map[string]struct {
		fallback time.Duration
		request  *http.Request
		handler  http.Handler
	}{
		"exists in header": {
			fallback: time.Second,
			request: &http.Request{
				Header: http.Header{
					header.XTimeout: []string{"10ms"},
				},
			},
			handler: timeout(15 * time.Millisecond),
		},
		"fallback cause empty": {
			fallback: 10 * time.Millisecond,
			request:  &http.Request{},
			handler:  timeout(15 * time.Millisecond),
		},
		"fallback cause invalid": {
			fallback: 10 * time.Millisecond,
			request: &http.Request{
				Header: http.Header{
					header.XTimeout: []string{"bad"},
				},
			},
			handler: timeout(15 * time.Millisecond),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			middleware := Timeout(test.fallback, corrector)
			recorder := httptest.NewRecorder()
			middleware(test.handler).ServeHTTP(recorder, test.request)
		})
	}
}
