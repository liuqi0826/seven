package core

import (
	"github.com/liuqi0826/seven/events"
)

type ICamera interface {
}

type IController interface {
}

type IRender interface {
}

type IEntity interface {
}

type IDisplayObject interface {
	events.IEventDispatcher

	GetID() string
	GetName() string
	SetName(string)
	GetRoot() IContainer
	SetRoot(IContainer)
	GetParent() IContainer
	SetParent(IContainer)
	GetLayerMask() int32
	SetLayerMask(int32)
	Render()
}

type IContainer interface {
	events.IEventDispatcher

	AddChild(IDisplayObject)
	RemoveChild(IDisplayObject) IDisplayObject
	RemoveAllChildren()
	GetChildByName(string) IDisplayObject
	GetChildrenNumber() int32
}
