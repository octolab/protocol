package header_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/protocol/http/header"
)

func TestDeadline(t *testing.T) {
	type tuple struct {
		deadline time.Time
		present  bool
	}

	now := time.Now()

	tests := map[string]struct {
		header    http.Header
		fallback  time.Duration
		expected  tuple
		precision time.Duration
	}{
		"exists in header": {
			header: http.Header{
				XDeadline: []string{now.Format(time.RFC3339Nano)},
			},
			fallback:  time.Second,
			expected:  tuple{now, true},
			precision: 0,
		},
		"exists in header, inaccuracy": {
			header: http.Header{
				XDeadline: []string{now.Format(time.RFC3339)},
			},
			fallback:  time.Second,
			expected:  tuple{now, true},
			precision: time.Second,
		},
		"exists in header, timeout": {
			header:    http.Header{XTimeout: []string{"100ms"}},
			fallback:  time.Second,
			expected:  tuple{now.Add(100 * time.Millisecond), true},
			precision: 10 * time.Millisecond,
		},
		"fallback cause empty": {
			header:    nil,
			fallback:  time.Second,
			expected:  tuple{now.Add(time.Second), false},
			precision: 10 * time.Millisecond,
		},
		"fallback cause invalid": {
			header: http.Header{
				XDeadline: []string{"bad"},
				XTimeout:  []string{"bad"},
			},
			fallback:  time.Second,
			expected:  tuple{now.Add(time.Second), false},
			precision: 10 * time.Millisecond,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			deadline, present := Deadline(test.header, test.fallback)
			assert.Equal(t, test.expected.present, present)
			assert.WithinDuration(t, test.expected.deadline, deadline, test.precision)
		})
	}
}

func TestTimeout(t *testing.T) {
	type tuple struct {
		duration time.Duration
		present  bool
	}

	tests := map[string]struct {
		header   http.Header
		fallback time.Duration
		expected tuple
	}{
		"exists in header": {
			header:   http.Header{XTimeout: []string{"100ms"}},
			fallback: time.Second,
			expected: tuple{100 * time.Millisecond, true},
		},
		"fallback cause empty": {
			header:   nil,
			fallback: time.Second,
			expected: tuple{time.Second, false},
		},
		"fallback cause invalid": {
			header:   http.Header{XTimeout: []string{"bad"}},
			fallback: time.Second,
			expected: tuple{time.Second, false},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			duration, present := Timeout(test.header, test.fallback)
			assert.Equal(t, test.expected, tuple{duration, present})
		})
	}
}
