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
	listeners []chan<- string
}

func NewEvent(name string) Event {
	return Event{
		name:      name,
		paras:     make([]int, 0),
		listeners: make([]chan<- string, 0),
	}
}

func (e *Event) AddListener(l *IListener) {
	ch := make(chan string)
	e.listeners = append(e.listeners, chan<- string(ch))
	go l.Listen((<-chan string)(ch))
}

func (e *Event) Trigger(content string) {
	for _, c := range e.listeners {
		c <- content
	}
}
