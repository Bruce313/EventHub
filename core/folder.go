package core

import ()

type Folder struct {
	name    string
	events  map[string]Event
	folders map[string]Folder
}

const (
	pathSep = "."
)

//create Folder from path like org.golang.core
func NewFolder(paraName string) Folder {
	ret := Folder{name: paraName}
	ret.events = make(map[string]Event, 0)
	ret.folders = make(map[string]Folder, 0)
	return ret
}

func (f *Folder) GetEvents() []Event {
	evs := make([]Event, len(f.events))
	for _, v := range f.events {
		evs = append(evs, v)
	}
	return evs
}

func (f *Folder) HasEvent(name string) bool {
	v, ok := f.events[name]
	return ok
}

//func (f *Folder) AddPath
//func (f *Folder) AddEvent
