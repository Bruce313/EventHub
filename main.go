package main

import (
	"fmt"
	core "github.com/Bruce313/EventHub/core"
	httpDriver "github.com/Bruce313/EventHub/drivers/http"
)

func main() {
	fmt.Println("main begin")
	fm := core.NewFolderManager()
	p := httpDriver.NewPort(&fm)
	p.Serve(&fm)
}
