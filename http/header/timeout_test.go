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
			http.Header{XDeadlineHeader: []string{now.Format(time.RFC3339Nano)}},
			time.Second,
			tuple{now, true},
			0,
		},
		"exists in header, inaccuracy": {
			http.Header{XDeadlineHeader: []string{now.Format(time.RFC3339)}},
			time.Second,
			tuple{now, true},
			time.Second,
		},
		"exists in header, timeout": {
			http.Header{XTimeoutHeader: []string{"100ms"}},
			time.Second,
			tuple{now.Add(100 * time.Millisecond), true},
			10 * time.Millisecond,
		},
		"fallback cause empty": {
			nil,
			time.Second,
			tuple{now.Add(time.Second), false},
			10 * time.Millisecond,
		},
		"fallback cause invalid": {
			http.Header{XDeadlineHeader: []string{"bad"}, XTimeoutHeader: []string{"bad"}},
			time.Second,
			tuple{now.Add(time.Second), false},
			10 * time.Millisecond,
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
			http.Header{XTimeoutHeader: []string{"100ms"}},
			time.Second,
			tuple{100 * time.Millisecond, true},
		},
		"fallback cause empty": {
			nil,
			time.Second,
			tuple{time.Second, false},
		},
		"fallback cause invalid": {
			http.Header{XTimeoutHeader: []string{"bad"}},
			time.Second,
			tuple{time.Second, false},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			duration, present := Timeout(test.header, test.fallback)
			assert.Equal(t,
				test.expected,
				tuple{duration, present},
			)
		})
	}
}
