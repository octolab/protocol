package header_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/protocol/http/header"
)

func TestNoCache(t *testing.T) {
	tests := map[string]struct {
		header   http.Header
		expected bool
	}{
		"exists in header": {
			header:   http.Header{CacheControl: []string{"no-cache"}},
			expected: true,
		},
		"exists in header, case insensitive": {
			header:   http.Header{CacheControl: []string{"No-Cache"}},
			expected: true,
		},
		"empty duration": {
			header:   nil,
			expected: false,
		},
		"another duration": {
			header:   http.Header{CacheControl: []string{"only-if-cached"}},
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, NoCache(test.header))
		})
	}
}

func TestStrict(t *testing.T) {
	tests := map[string]struct {
		header   http.Header
		expected bool
	}{
		"exists in header, string": {
			header:   http.Header{XStrict: []string{"true"}},
			expected: true,
		},
		"exists in header, numeric": {
			header:   http.Header{XStrict: []string{"1"}},
			expected: true,
		},
		"exists in header, case insensitive": {
			header:   http.Header{XStrict: []string{"True"}},
			expected: true,
		},
		"empty duration": {
			header:   nil,
			expected: false,
		},
		"invalid duration": {
			header:   http.Header{XStrict: []string{"invalid"}},
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, Strict(test.header))
		})
	}
}
