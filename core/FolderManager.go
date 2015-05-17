package core

type RegisterConfig struct {
	name string
}
type ListenConfig struct {
	name     string
	listener IListener
}
type TriggerConfig struct {
	name string
}
type FolderManager struct {
	RegChan     chan RegisterConfig
	ListenChan  chan ListenConfig
	TriggerChan chan TriggerConfig
	folder      Folder
}

func NewFolderManager() {
	return FolderManager{
		RegChan:     make(chan RegisterConfig),
		ListenChan:  make(chan ListenConfig),
		TriggerChan: make(chan TriggerConfig),
		folder:      NewFolder("root"),
	}
}

func (fm *FolderManager) Start() {
	select {
	case rc := <-fm.RegChan:
		fm.handleReg(rc)
	case lc := <-fm.ListenChan:
		fm.handleListen(lc)
	case tc := <-rm.TriggerChan:
		fm.handleTrigger(tc)
	}
}

func (fm *FolderManager) handleReg(rc RegConfig) {

}

func (fm *FolderManager) handleListen(lc ListenConfig) {
}

func (fm *FolderManager) handleTrigger(tc TriggerConfig) {
}
