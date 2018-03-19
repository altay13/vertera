package tarantool

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
