## display the help text
help:
	$(info Available targets)
	$(info ------------------)
	@awk '/^[a-zA-Z\-\_0-9]+:/ {                    \
		nb = sub( /^## /, "", helpMsg );              \
		if(nb == 0) {                                 \
			helpMsg = $$0;                              \
			nb = sub( /^[^:]*:.* ## /, "", helpMsg );   \
		}                                             \
		if (nb)                                       \
			print  $$1 "  " helpMsg;                    \
	}                                               \
	{ helpMsg = $$0 }'                              \
	$(MAKEFILE_LIST)  |  grep --color '^[^: ]*'

## Get necessary go dependencies
get-dependencies:
	go get github.com/go-redis/redis
	go get github.com/gocql/gocql
	go get github.com/hazelcast/hazelcast-go-client
	go get github.com/tarantool/go-tarantool

## Starts all available db containers and runs interactive console
run-example-interactive-all: start-all-container
	go run main.go interactive

## Starts redis db containers and runs interactive console
run-example-interactive-redis: start-redis-container
	go run main.go interactive --db=redis --config='localhost'

## Starts hazelcast db containers and runs interactive console
run-example-interactive-hazelcast: start-hazelcast-container
	go run main.go interactive --db=hazelcast --config='localhost'

## Starts tarantool db containers and runs interactive console
run-example-interactive-tarantool: start-tarantool-container
	go run main.go interactive --db=tarantool --config='localhost'

## Starts cassandra db containers and runs interactive console
run-example-interactive-cassandra: start-cassandra-container
	go run main.go interactive --db=cassandra --config='localhost'

## Starts rocksdb db containers and runs interactive console
run-example-interactive-rocksdb: start-rocksdb-container
	go run main.go interactive --db=rocksdb --config='localhost'

include example/*.mk

.DEFAULT_GOAL := help
.PHONY: help \
	get-dependencies \
	run-example-interactive \
	run-example-interactive-redis \
	run-example-interactive-hazelcast \
	run-example-interactive-tarantool \
	run-example-interactive-cassandra \
	run-example-interactive-rocksdb
