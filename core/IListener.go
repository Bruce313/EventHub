package core

type IListener interface {
	Listen(ch <-chan string)
}
