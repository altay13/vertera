package redis

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
	r.cleanupHook()

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

func (r *Redis) cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)

	go func() {
		<-c
		r.pool.Close()
		os.Exit(0)
	}()
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

	key := event.Key

	val, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		revent.Err = fmt.Errorf("Failed to GET. %s", err.Error())
	} else {
		revent.Key = key
		revent.Value = val
	}

	return revent
}
