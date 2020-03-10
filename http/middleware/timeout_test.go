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
			time.Second,
			&http.Request{
				Header: http.Header{
					header.XDeadlineHeader: []string{
						time.Now().Add(10 * time.Millisecond).Format(time.RFC3339Nano),
					},
				},
			},
			http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				timer := time.NewTimer(20 * time.Millisecond)
				select {
				case <-req.Context().Done():
					assert.True(t, true)
				case <-timer.C:
					assert.True(t, false)
				}
				assert.True(t, timer.Stop())
			}),
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
			time.Second,
			&http.Request{
				Header: http.Header{
					header.XTimeoutHeader: []string{
						(10 * time.Millisecond).String(),
					},
				},
			},
			http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				timer := time.NewTimer(20 * time.Millisecond)
				select {
				case <-req.Context().Done():
					assert.True(t, true)
				case <-timer.C:
					assert.True(t, false)
				}
				assert.True(t, timer.Stop())
			}),
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
