export CONTAINER_NAME=local.test
export HOST_IP=127.0.0.1

.PNONY: start-redis start-cassandra start-rocksdb start-hazelcast help start-all

## Start redis docker container
start-redis:
	docker stop redis-${CONTAINER_NAME} || true && docker rm redis-${CONTAINER_NAME} || true
	docker run -d --name redis-${CONTAINER_NAME} -v `pwd`/redis-data:/data -p 6379:6379 redis:latest redis-server --appendonly yes

## Start cassandra docker container
start-cassandra:
	docker stop cassandra-${CONTAINER_NAME} || true && docker rm cassandra-${CONTAINER_NAME} || true
	docker run -d --name cassandra-${CONTAINER_NAME} -v `pwd`/cassandra-data:/var/lib/cassandra -p 9042:9042 cassandra:latest

## Start rocksdb docker container
start-rocksdb:
	# docker stop rocksdb-${CONTAINER_NAME} || true && docker rm rocksdb-${CONTAINER_NAME} || true
	echo rocksdb

## Start hazelcast docker container
start-hazelcast:
	docker stop hazelcast-${CONTAINER_NAME} || true && docker rm hazelcast-${CONTAINER_NAME} || true
	docker run -d --name hazecast-${CONTAINER_NAME} -p 5701:5701 hazelcast/hazelcast

## Start all DBs
start-all: start-redis start-cassandra start-rocksdb start-hazelcast

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
			print  $$1 "\t" helpMsg;                    \
	}                                               \
	{ helpMsg = $$0 }'                              \
	$(MAKEFILE_LIST)  |  grep --color '^[^: ]*'
