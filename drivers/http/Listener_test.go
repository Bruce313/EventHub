package http

import (
	core "github.com/Bruce313/EventHub/core"
	"testing"
	"time"
)

func TestListener(t *testing.T) {
	l := Listener{name: "test"}
	f := core.NewFolder("root")
	f.Register("hello")
	f.BindListener("hello", &l)
	f.Trigger("hello", "message from http driver")
	time.Sleep(1000 * time.Millisecond)
}
