package driver

import (
	FolderManager "github.com/Bruce313/EventHub/core/FolderManager"
)

type IPort interface {
	Serve(fm *FolderManager)
}
