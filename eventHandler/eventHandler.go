package eventHandler

import (
	"fmt"
	"sync"
)

type EventHandler struct {
	RequestChan chan Request
	eventStore  EventStore

	stopChan chan bool
	routines sync.WaitGroup
}

type Event struct {
	Key   interface{}
	Value []byte
	Err   error
}

type Request struct {
	Cmd      command
	Event    *Event
	Response chan *Event
}

type command string

const (
	GET      command = "get"
	SET      command = "set"
	STOP     command = "stop"
	RESPONSE command = "response"
)

func NewEventHandler(eventStore EventStore) *EventHandler {
	r := &EventHandler{
		RequestChan: make(chan Request, 100),
		eventStore:  eventStore,
		stopChan:    make(chan bool, 1),
	}

	r.goFunc(r.startHandler)

	return r
}

func (h *EventHandler) startHandler() {
	for req := range h.RequestChan {
		switch req.Cmd {
		case GET:
			event := h.eventStore.Get(req.Event)
			req.Response <- event
		case SET:
			event := h.eventStore.Set(req.Event)
			req.Response <- event
		case STOP:
			h.stopChan <- true
		default:
			fmt.Println("Unknown command.", req.Cmd)
		}

		select {
		case <-h.stopChan:
			return
		default:
			continue
		}
	}
}

func (h *EventHandler) GetDBName() string {
	if h.eventStore == nil {
		return ""
	}
	return h.eventStore.GetName()
}

func (h *EventHandler) CloseDB() {
	h.eventStore.Disconnect()
}

func (h *EventHandler) goFunc(fn func()) {
	h.routines.Add(1)
	go func() {
		defer h.routines.Done()
		fn()
	}()
}
