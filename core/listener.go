package core

import (
	"fmt"
)

type Listener struct {
	url  string
	port int
	name string
}

func (l *Listener) accept() {
	fmt.Println(l.name, "accept")
}
