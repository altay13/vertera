package cassandra

import "github.com/gocql/gocql"

type Config struct {
	Host        []string
	Consistency gocql.Consistency
}

func DefaultConfig() *Config {
	config := &Config{
		Host:        []string{":9042"},
		Consistency: gocql.Quorum,
	}

	return config
}
