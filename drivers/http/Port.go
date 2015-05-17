package http

import (
    core "github.com/Bruce313/core"
)

type Port struct {
    fm core.FolderManager
}

func (p *Port) Serve(fm core.FolderManager) {
    p.fm = fm
    p.start()
}

func (p *Port) start() {
}
