package redis

import (
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	Host         string
	PoolIdleSize uint8
	IdleTimeout  uint16
}

func DefaultConfig() *Config {
	config := &Config{
		Host:         ":6379",
		PoolIdleSize: 5,
		IdleTimeout:  120,
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
		case "poolidlesize":
			if poolidlesize, err := strconv.Atoi(vals[1]); err != nil {
				rerr = fmt.Errorf("PoolIdleSize param should be integer.")
			} else {
				c.PoolIdleSize = uint8(poolidlesize)
			}
		case "idletimeout":
			if idletimeout, err := strconv.Atoi(vals[1]); err != nil {
				rerr = fmt.Errorf("PoolIdleSize param should be integer.")
			} else {
				c.IdleTimeout = uint16(idletimeout)
			}
			// c. = vals[1]
		default:
			// nothing yet. Later
		}
	}
	return rerr
}
