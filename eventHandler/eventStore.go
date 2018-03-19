package eventHandler

const (
	REDIS     string = "redis"
	CASSANDRA string = "cassandra"
	ROCKSDB   string = "rocksdb"
	HAZELCAST string = "hazelcast"
	TARANTOOL string = "tarantool"
)

type EventStore interface {
	Set(*Event) *Event
	Get(*Event) *Event
	GetName() string
	Disconnect()
}
