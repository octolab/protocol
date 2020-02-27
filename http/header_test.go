package http_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/protocol/http"
)

func TestHeader_Deadline(t *testing.T) {
	type tuple struct {
		deadline time.Time
		present  bool
	}

	now := time.Now()

	tests := map[string]struct {
		header    Header
		fallback  time.Duration
		expected  tuple
		precision time.Duration
	}{
		"exists in header": {
			Header{XDeadlineHeader: []string{now.Format(time.RFC3339Nano)}},
			time.Second,
			tuple{now, true},
			0,
		},
		"exists in header, inaccuracy": {
			Header{XDeadlineHeader: []string{now.Format(time.RFC3339)}},
			time.Second,
			tuple{now, true},
			time.Second,
		},
		"exists in header, timeout": {
			Header{XTimeoutHeader: []string{"100ms"}},
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
			Header{XDeadlineHeader: []string{"invalid"}, XTimeoutHeader: []string{"invalid"}},
			time.Second,
			tuple{now.Add(time.Second), false},
			10 * time.Millisecond,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			deadline, present := test.header.Deadline(test.fallback)
			assert.Equal(t, test.expected.present, present)
			assert.WithinDuration(t, test.expected.deadline, deadline, test.precision)
		})
	}
}

func TestHeader_NoCache(t *testing.T) {
	tests := map[string]struct {
		header   Header
		expected bool
	}{
		"exists in header": {
			Header{CacheControlHeader: []string{"no-cache"}},
			true,
		},
		"exists in header, case insensitive": {
			Header{CacheControlHeader: []string{"No-Cache"}},
			true,
		},
		"empty duration": {
			nil,
			false,
		},
		"another duration": {
			Header{CacheControlHeader: []string{"only-if-cached"}},
			false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.header.NoCache())
		})
	}
}

func TestHeader_Strict(t *testing.T) {
	tests := map[string]struct {
		header   Header
		expected bool
	}{
		"exists in header, string": {
			Header{XStrictHeader: []string{"true"}},
			true,
		},
		"exists in header, numeric": {
			Header{XStrictHeader: []string{"1"}},
			true,
		},
		"exists in header, case insensitive": {
			Header{XStrictHeader: []string{"True"}},
			true,
		},
		"empty duration": {
			nil,
			false,
		},
		"invalid duration": {
			Header{XStrictHeader: []string{"invalid"}},
			false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.header.Strict())
		})
	}
}

func TestHeader_Timeout(t *testing.T) {
	type tuple struct {
		duration time.Duration
		present  bool
	}

	tests := map[string]struct {
		header   Header
		fallback time.Duration
		expected tuple
	}{
		"exists in header": {
			Header{XTimeoutHeader: []string{"100ms"}},
			time.Second,
			tuple{100 * time.Millisecond, true},
		},
		"fallback cause empty": {
			nil,
			time.Second,
			tuple{time.Second, false},
		},
		"fallback cause invalid": {
			Header{XTimeoutHeader: []string{"invalid"}},
			time.Second,
			tuple{time.Second, false},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			duration, present := test.header.Timeout(test.fallback)
			assert.Equal(t,
				test.expected,
				tuple{duration, present},
			)
		})
	}
}
