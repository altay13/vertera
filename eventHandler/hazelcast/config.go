package hazelcast

import (
	"strings"
)

type Config struct {
	Host string
}

func DefaultConfig() *Config {
	config := &Config{
		Host: "127.0.0.1:5701",
	}

	return config
}

func (c *Config) SetByConfigString(configStr string) error {
	var rerr error
	confs := strings.Split(configStr, ";")
	for _, conf := range confs {
		vals := strings.Split(conf, "=")
		switch strings.ToLower(vals[0]) {
		case "host":
			c.Host = vals[1]
		default:
			// nothing yet. Later
		}
	}
	return rerr
}
