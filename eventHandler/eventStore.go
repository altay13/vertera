package eventHandler

type EventStore interface {
	Set(*Event) *Event
	Get(*Event) *Event
}
