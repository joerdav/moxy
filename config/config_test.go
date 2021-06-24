package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetDestinationFromConfig(t *testing.T) {
	var tests = []struct {
		name     string
		config   Config
		input    string
		expected string
	}{
		{
			name:     "GetDestinationFromConfig: No upstreams always returns default",
			config:   Config{DefaultUpstream: "http://default"},
			input:    "/any-path",
			expected: "http://default",
		},
		{
			name: "GetDestinationFromConfig: Exact match on path prefix returns path",
			config: Config{Upstreams: []Upstream{
				{PathPrefix: "/app", Host: "http://app"},
			}},
			input:    "/app",
			expected: "http://app",
		},
		{
			name: "GetDestinationFromConfig: Match on path trailing slash prefix returns path",
			config: Config{Upstreams: []Upstream{
				{PathPrefix: "/app", Host: "http://app"},
			}},
			input:    "/app/",
			expected: "http://app",
		},
		{
			name: "GetDestinationFromConfig: Match on sub-path prefix returns path",
			config: Config{Upstreams: []Upstream{
				{PathPrefix: "/app", Host: "http://app"},
			}},
			input:    "/app/resource",
			expected: "http://app",
		},
		{
			name: "GetDestinationFromConfig: Two matching paths, return first",
			config: Config{Upstreams: []Upstream{
				{PathPrefix: "/app", Host: "http://first"},
				{PathPrefix: "/app/resource", Host: "http://second"},
			}},
			input:    "/app/myresource/subresource",
			expected: "http://first",
		},
		{
			name: "GetDestinationFromConfig: No matches, return default",
			config: Config{
				Upstreams: []Upstream{
					{PathPrefix: "/app", Host: "http://first"},
					{PathPrefix: "/app/resource", Host: "http://second"},
				},
				DefaultUpstream: "http://default",
			},
			input:    "/another-resource",
			expected: "http://default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.config.GetDestination(tt.input)

			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Error(diff)
			}
		})
	}
}
