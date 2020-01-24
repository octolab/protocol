package http_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/protocol/http"
)

func TestHeader_NoCache(t *testing.T) {
	tests := map[string]struct {
		header   http.Header
		expected bool
	}{
		"exists in header": {
			http.Header{CacheControlHeader: []string{"no-cache"}},
			true,
		},
		"exists in header, case insensitive": {
			http.Header{CacheControlHeader: []string{"No-Cache"}},
			true,
		},
		"empty value": {
			nil,
			false,
		},
		"another value": {
			http.Header{CacheControlHeader: []string{"only-if-cached"}},
			false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, Header(test.header).NoCache())
		})
	}
}

func TestHeader_Strict(t *testing.T) {
	tests := map[string]struct {
		header   http.Header
		expected bool
	}{
		"exists in header, string": {
			http.Header{XStrictHeader: []string{"true"}},
			true,
		},
		"exists in header, numeric": {
			http.Header{XStrictHeader: []string{"1"}},
			true,
		},
		"exists in header, case insensitive": {
			http.Header{XStrictHeader: []string{"True"}},
			true,
		},
		"empty value": {
			nil,
			false,
		},
		"invalid value": {
			http.Header{XStrictHeader: []string{"invalid"}},
			false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, Header(test.header).Strict())
		})
	}
}

func TestHeader_Timeout(t *testing.T) {
	tests := map[string]struct {
		header     http.Header
		fallback   time.Duration
		percentage float64
		expected   time.Duration
	}{
		"exists in header": {
			http.Header{XTimeoutHeader: []string{"100ms"}},
			time.Second,
			0.9,
			90 * time.Millisecond,
		},
		"fallback cause empty": {
			nil,
			time.Second,
			0.9,
			900 * time.Millisecond,
		},
		"fallback cause invalid": {
			http.Header{XTimeoutHeader: []string{"invalid"}},
			time.Second,
			0.9,
			900 * time.Millisecond,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t,
				test.expected,
				Header(test.header).Timeout(test.fallback, test.percentage),
			)
		})
	}
}
