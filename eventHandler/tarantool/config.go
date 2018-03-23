package tarantool

import (
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	Host          string
	Timeout       uint16
	Reconnect     uint8
	MaxReconnects uint8
	User          string
	Pass          string
	Space         string
}

func DefaultConfig() *Config {
	config := &Config{
		Host:          "127.0.0.1:3301",
		Timeout:       500,
		Reconnect:     1,
		MaxReconnects: 3,
		User:          "test",
		Pass:          "test",
		Space:         "tester",
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
		case "timeout":
			if timeout, err := strconv.Atoi(vals[1]); err != nil {
				rerr = fmt.Errorf("Timeout param should be integer.")
			} else {
				c.Timeout = uint16(timeout)
			}
		case "reconnect":
			if reconnect, err := strconv.Atoi(vals[1]); err != nil {
				rerr = fmt.Errorf("Reconnect param should be integer.")
			} else {
				c.Reconnect = uint8(reconnect)
			}
		case "maxreconnects":
			if maxreconnects, err := strconv.Atoi(vals[1]); err != nil {
				rerr = fmt.Errorf("MaxReconnects param should be integer.")
			} else {
				c.MaxReconnects = uint8(maxreconnects)
			}
		case "user":
			c.User = vals[1]
		case "pass":
			c.Pass = vals[1]
		case "space":
			c.Space = vals[1]
		default:
			// nothing yet. Later
		}
	}
	return rerr
}
