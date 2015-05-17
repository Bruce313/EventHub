package http

import (
	"fmt"
	core "github.com/Bruce313/EventHub/core"
	http "net/http"
)

type HttpHandler struct {
	regChan     chan<- core.RegisterConfig
	listenChan  chan<- core.ListenConfig
	triggerChan chan<- core.TriggerConfig
}

func NewHttpHandler() {
}

func (th *HttpHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	rw.Write([]byte("hello"))
}

func (th *HttpHandler) handleRegister(req *http.Request) {

}
