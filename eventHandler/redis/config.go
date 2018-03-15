package redis

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
