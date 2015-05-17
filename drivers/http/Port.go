package http

import (
	core "github.com/Bruce313/EventHub/core"
	"net/http"
	"time"
)

type Port struct {
	fm *core.FolderManager
}

func NewPort(fm *core.FolderManager) Port {
	return Port{fm: fm}
}

func (p *Port) Serve(fm *core.FolderManager) {
	p.fm = fm
	p.start()
}

func (p *Port) start() {
	httpHandler := HttpHandler{
		regChan:     p.fm.RegChan,
		listenChan:  p.fm.ListenChan,
		triggerChan: p.fm.TriggerChan,
	}
	s := &http.Server{
		Addr:           ":7000",
		Handler:        &httpHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
