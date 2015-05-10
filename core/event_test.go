package core

import (
	"testing"
)

func TestTrigger(t *testing.T) {
	e := Event{name: "test"}
	e.listeners = make([]Listener, 1)
	e.listeners = append(e.listeners, Listener{name: "listener"})
	e.trigger()
}
