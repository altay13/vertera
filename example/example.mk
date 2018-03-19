export CONTAINER_NAME=local.test
export HOST_IP=127.0.0.1

## Starts redis docker container
start-redis-container:
	docker stop redis-${CONTAINER_NAME} || true && docker rm redis-${CONTAINER_NAME} || true
	docker run -d --name redis-${CONTAINER_NAME} -v `pwd`/example/redis-data:/data -p 6379:6379 redis:latest redis-server --appendonly yes

## Starts cassandra docker container
start-cassandra-container:
	docker stop cassandra-${CONTAINER_NAME} || true && docker rm cassandra-${CONTAINER_NAME} || true
	docker run -d --name cassandra-${CONTAINER_NAME} -v `pwd`/example/cassandra-data:/var/lib/cassandra -p 9042:9042 cassandra:latest

## Starts rocksdb docker container
start-rocksdb-container:
	# docker stop rocksdb-${CONTAINER_NAME} || true && docker rm rocksdb-${CONTAINER_NAME} || true
	echo rocksdb

## Starts hazelcast docker container
start-hazelcast-container:
	docker stop hazelcast-${CONTAINER_NAME} || true && docker rm hazelcast-${CONTAINER_NAME} || true
	docker run -d --name hazelcast-${CONTAINER_NAME} -v `pwd`/example/hazelcast-data:/var/lib/hazelcast -p 5701:5701 hazelcast/hazelcast

## Starts Tarantool docker CONTAINER_NAME
start-tarantool-container:
	docker stop tarantool-${CONTAINER_NAME} || true && docker rm tarantool-${CONTAINER_NAME} || true
	docker run --name tarantool-${CONTAINER_NAME} -v `pwd`/example/tarantool-data:/var/lib/tarantool -p3301:3301 -d tarantool/tarantool

## Starts all DB containers
start-all-container: start-redis-container start-cassandra-container start-rocksdb-container start-hazelcast-container start-tarantool-container

.PNONY: start-redis-container \
	start-cassandra-container \
	start-rocksdb-container \
	start-hazelcast-container \
	start-tarantool-container \
	start-all-container
