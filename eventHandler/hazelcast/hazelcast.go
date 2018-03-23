package hazelcast

import (
	"fmt"

	"github.com/altay13/vertera/eventHandler"
	hazelcast "github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/core"
)

type Hazelcast struct {
	Config
	client hazelcast.IHazelcastInstance
	mp     core.IMap
}

func NewHazelcast(conf *Config) (*Hazelcast, error) {
	h := &Hazelcast{
		Config: *conf,
	}

	err := h.newCluster()
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (h *Hazelcast) newCluster() error {
	config := hazelcast.NewHazelcastConfig()
	config.ClientNetworkConfig().AddAddress(h.Host)
	var err error
	h.client, err = hazelcast.NewHazelcastClientWithConfig(config)
	if err != nil {
		return err
	}

	h.mp, _ = h.client.GetMap("gohazelcast")
	return nil
}

func (h *Hazelcast) Disconnect() {
	h.client.Shutdown()
}

func (h *Hazelcast) Set(event *eventHandler.Event) *eventHandler.Event {
	revent := &eventHandler.Event{
		Key:   event.Key,
		Value: event.Value,
	}

	_, err := h.mp.TryPut(event.Key, event.Value)
	if err != nil {
		revent.Err = fmt.Errorf("Failed to SET. %s", err.Error())
	}

	return revent
}

func (h *Hazelcast) Get(event *eventHandler.Event) *eventHandler.Event {
	revent := &eventHandler.Event{}

	if event.Key == nil {
		revent.Err = fmt.Errorf("Failed to GET. No Key provided.")
		return revent
	}

	val, err := h.mp.Get(event.Key)
	if err != nil {
		revent.Err = fmt.Errorf("Failed to GET. %s", err.Error())
	} else {
		if val == nil {
			revent.Err = fmt.Errorf("Failed to GET. Nil returned.")
			return revent
		}
		revent.Key = event.Key
		revent.Value = val.([]byte)
	}

	return revent
}

func (h *Hazelcast) GetName() string {
	return eventHandler.HAZELCAST
}
