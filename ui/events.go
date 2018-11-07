package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Event int

const (
	KeyEvent Event = iota
	ButtonEvent
	HoverEvent
	SelectEvent
	DeselectEvent
	TextEvent
	TextEditEvent
	TickEvent
)

type EventData struct {
	EventType   Event
	Key         sdl.Keysym
	KeyState    uint8
	Button      uint8
	ButtonState uint8
	MouseX      int32
	MouseY      int32
	Text        string
	Start       int32
	Length      int32
}

type EventManager struct {
	events map[Event]func(EventData)
}

func NewEventManager() *EventManager {
	return &EventManager{make(map[Event]func(EventData), 0)}
}

func (ev *EventManager) Register(event Event, f func(EventData)) {
	ev.events[event] = f
}

func (ev *EventManager) Trigger(event Event, data EventData) bool {
	if f, ok := ev.events[event]; ok {
		f(data)
		return true
	}
	return false
}
