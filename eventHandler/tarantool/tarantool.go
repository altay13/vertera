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

func NewTarantool(conf *Config) (*Tarantool, error) {
	c := &Tarantool{
		Config: *conf,
	}

	err := c.newConnection()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Tarantool) newConnection() error {
	opts := tarantool.Opts{
		Timeout:       time.Duration(c.Timeout) * time.Millisecond,
		Reconnect:     time.Duration(c.Reconnect) * time.Second,
		MaxReconnects: uint(c.MaxReconnects),
		User:          "guest",
		// Pass:          c.Pass,
	}
	var err error
	c.client, err = tarantool.Connect(c.Host, opts)
	if err != nil {
		return err
	}

	resp, err := c.client.Ping()
	if err != nil {
		return err
	}
	log.Println(resp.Code)
	log.Println(resp.Data)

	return nil
}

func (c *Tarantool) Disconnect() {
	c.client.Close()
}

func (c *Tarantool) Set(event *eventHandler.Event) *eventHandler.Event {
	revent := &eventHandler.Event{
		Key:   event.Key,
		Value: event.Value,
	}

	_, err := c.client.Insert(c.Space, []interface{}{event.Key, event.Value})
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

	resp, err := c.client.Select(c.Space, "primary", 0, 1, tarantool.IterEq, []interface{}{event.Key})

	if err != nil {
		revent.Err = fmt.Errorf("Failed to GET. %s", err.Error())
	} else {
		if len(resp.Data) == 0 {
			revent.Err = fmt.Errorf("Failed to GET. Nil returned.")
			return revent
		}
		revent.Key = event.Key
		revent.Value = resp.Data[0].([]interface{})[1].([]byte)
	}

	return revent
}

func (c *Tarantool) GetName() string {
	return eventHandler.TARANTOOL
}
