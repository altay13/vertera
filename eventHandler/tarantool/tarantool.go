package tarantool

import (
	"fmt"
	"log"
	"time"

	"github.com/altay13/vertera/eventHandler"
	tarantool "github.com/tarantool/go-tarantool"
)

type Tarantool struct {
	Config
	client *tarantool.Connection
}

func NewTarantool(conf *Config) *Tarantool {
	c := &Tarantool{
		Config: *conf,
	}

	c.newConnection()
	return c
}

func (c *Tarantool) newConnection() {
	opts := tarantool.Opts{
		Timeout:       time.Duration(c.Timeout) * time.Millisecond,
		Reconnect:     time.Duration(c.Reconnect) * time.Second,
		MaxReconnects: uint(c.MaxReconnects),
		User:          "guest",
		// Pass: c.Pass,
	}
	c.client, _ = tarantool.Connect(c.Host, opts)

	resp, err := c.client.Ping()
	log.Println(resp.Code)
	log.Println(resp.Data)
	log.Println(err)
}

func (c *Tarantool) Disconnect() {
	c.client.Close()
}

func (c *Tarantool) Set(event *eventHandler.Event) *eventHandler.Event {
	revent := &eventHandler.Event{
		Err: fmt.Errorf("Key %s is set", event.Key.(string)),
	}

	_, err := c.client.Insert("golang", []interface{}{event.Key, event.Value})
	if err != nil {
		revent.Err = fmt.Errorf("Failed to SET. %s", err.Error())
	}

	return revent
}

func (c *Tarantool) Get(event *eventHandler.Event) *eventHandler.Event {
	revent := &eventHandler.Event{}

	if event.Key == nil {
		revent.Err = fmt.Errorf("Failed to GET. No Key provided.")
		return revent
	}

	resp, err := c.client.Select("golang", "primary", 0, 1, tarantool.IterGe, []interface{}{event.Key})
	if err != nil {
		revent.Err = fmt.Errorf("Failed to GET. %s", err.Error())
	} else {
		revent.Key = event.Key
		revent.Value = resp.Data[1].([]byte)
	}

	return revent
}

func (c *Tarantool) GetName() string {
	return eventHandler.TARANTOOL
}
