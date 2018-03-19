package redis

import (
	"fmt"
	"time"

	"github.com/altay13/vertera/eventHandler"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	Config
	pool *redis.Pool
}

func NewRedis(conf *Config) *Redis {
	r := &Redis{
		Config: *conf,
	}

	r.newPool()
	return r
}

func (r *Redis) newPool() {
	r.pool = &redis.Pool{
		MaxIdle:     int(r.PoolIdleSize),
		IdleTimeout: time.Duration(r.IdleTimeout) * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", r.Host)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (r *Redis) Disconnect() {
	r.pool.Close()
}

func (r *Redis) Set(event *eventHandler.Event) *eventHandler.Event {
	conn := r.pool.Get()
	defer conn.Close()

	revent := &eventHandler.Event{}

	_, err := conn.Do("SET", event.Key, event.Value)
	if err != nil {
		revent.Err = fmt.Errorf("Failed to SET. %s", err.Error())
	}

	return revent
}

func (r *Redis) Get(event *eventHandler.Event) *eventHandler.Event {
	conn := r.pool.Get()
	defer conn.Close()

	revent := &eventHandler.Event{}

	if event.Key == nil {
		revent.Err = fmt.Errorf("Failed to GET. No Key provided.")
		return revent
	}

	val, err := redis.Bytes(conn.Do("GET", event.Key))
	if err != nil {
		revent.Err = fmt.Errorf("Failed to GET. %s", err.Error())
	} else {
		revent.Key = event.Key
		revent.Value = val
	}

	return revent
}

func (r *Redis) GetName() string {
	return eventHandler.REDIS
}
