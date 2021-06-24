package config

import "strings"

type Upstream struct {
	PathPrefix string `toml:"path_prefix"`
	Url        string `toml:"url"`
}

type Config struct {
	MoxyPath        string     `toml:"moxy_path"`
	DefaultUpstream string     `toml:"default_upstream"`
	Upstreams       []Upstream `toml:"upstream"`
}

func (c *Config) GetDestination(path string) string {
	for _, p := range c.Upstreams {
		if p.PathPrefix == path || strings.HasPrefix(path, p.PathPrefix+"/") {
			return p.Url
		}
	}

	return c.DefaultUpstream
}
