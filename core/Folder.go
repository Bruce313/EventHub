package core

import (
	"strings"
)

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

//trigger event, if not found, return false
func (f *Folder) Trigger(path string, content string) bool {
	strs := strings.Split(path, pathSep)
	first := strs[0]
	if len(strs) == 1 {
		e, ok := f.events[first]
		if ok {
			e.Trigger(content)
			return true
		}
		return false
	}
	cf, ok := f.folders[first]
	if ok {
		cf.Trigger(strings.Join(strs[1:], pathSep), content)
	}
	return false
}

//create event, if exsit, do nth
func (f *Folder) Register(path string) bool {
	strs := strings.Split(path, pathSep)
	first := strs[0]
	if len(strs) == 1 {
		_, ok := f.events[first]
		if ok {
			return false
		}
		f.events[first] = NewEvent(first)
		return true
	}
	cf, ok := f.folders[first]
	if ok {
		return cf.Register(strings.Join(strs[1:], pathSep))
	}
	return false
}

//add event, if not, create
//path: like org.com.core
func (f *Folder) BindEvent(path string, l IListener) bool {
	strs := strings.Split(path, pathSep)
	first := strs[0]
	if len(strs) == 1 {
		e, ok := f.events[first]
		if ok {
			e.AddListener(l)
			return true
		}
		return false
	}
	cf, ok := f.folders[first]
	if ok {
		return cf.BindEvent(strings.Join(strs[1:], pathSep), l)
	}
	return false
}

func (f *Folder) GetEvents() []Event {
	evs := make([]Event, len(f.events))
	for _, v := range f.events {
		evs = append(evs, v)
	}
	return evs
}

func (f *Folder) HasEvent(name string) bool {
	_, ok := f.events[name]
	return ok
}

//func (f *Folder) AddPath
//func (f *Folder) AddEvent
