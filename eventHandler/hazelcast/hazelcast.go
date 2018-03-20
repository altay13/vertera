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

func NewHazelcast(conf *Config) *Hazelcast {
	h := &Hazelcast{
		Config: *conf,
	}

	h.newCluster()
	return h
}

func (h *Hazelcast) newCluster() {
	config := hazelcast.NewHazelcastConfig()
	config.ClientNetworkConfig().AddAddress(h.Host)

	h.client, _ = hazelcast.NewHazelcastClientWithConfig(config)

	h.mp, _ = h.client.GetMap("gohazelcast")
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
		revent.Key = event.Key
		revent.Value = val.([]byte)
	}

	return revent
}

func (h *Hazelcast) GetName() string {
	return eventHandler.HAZELCAST
}
