package hazelcast

type Config struct {
	Host string
}

func DefaultConfig() *Config {
	config := &Config{
		Host: "127.0.0.1:5701",
	}

	return config
}
