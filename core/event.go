package core

//para types
const (
	Int = iota
	Dle
	Str
	Flt
	Bol
)

type Event struct {
	name      string
	paras     []int
	listeners []Listener
}

func (e *Event) trigger() {
	for _, l := range e.listeners {
		l.accept()
	}
}
