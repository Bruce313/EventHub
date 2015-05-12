package http

import (
	core "github.com/Bruce313/EventHub/core"
	"testing"
)

func TestListener(t *testing.T) {
	l := Listener{name: "test"}
	e := core.NewFolder("root")
	e.Register("hello")
	e.Listen("hello", l)
	e.Trigger("hello")
}
