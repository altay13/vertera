package eventHandler

type EventStore interface {
	Set(*Event) *Event
	Get(*Event) *Event
	GetName() string
	Disconnect()
}

const (
	REDIS     string = "redis"
	CASSANDRA string = "cassandra"
	ROCKSDB   string = "rocksdb"
	HAZELCAST string = "hazelcast"
	TARANTOOL string = "tarantool"
)

var (
	DBs map[string]bool = map[string]bool{
		REDIS:     true,
		CASSANDRA: true,
		ROCKSDB:   true,
		HAZELCAST: true,
		TARANTOOL: true,
	}
)
