package http

import (
	"fmt"
)

type Listener struct {
	name    string
	port    int
	address string
}

func (l *Listener) Listen(ch <-chan string) {
	for {
		fmt.Println(l.name, <-ch)
	}
}
