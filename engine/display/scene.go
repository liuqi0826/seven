package display

import (
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/events"
)

type Scene struct {
	events.EventDispatcher

	displayList []core.IDisplayObject
}

func (this *Scene) Scene() {
	this.EventDispatcher.EventDispatcher()
	this.displayList = make([]core.IDisplayObject, 0)
}
func (this *Scene) AddChild(displayObject core.IDisplayObject) {
	this.displayList = append(this.displayList, displayObject)

	displayObject.SetRoot(this)
	displayObject.SetParent(this)

	event := new(events.Event)
	event.Type = events.ADDED
	displayObject.DispatchEvent(event)
}
func (this *Scene) RemoveChild(displayObject core.IDisplayObject) core.IDisplayObject {
	for i, c := range this.displayList {
		if c == displayObject {
			this.displayList = append(this.displayList[:i], this.displayList[i+1:]...)
			c.SetRoot(nil)
			c.SetParent(nil)
			event := new(events.Event)
			event.Type = events.REMOVE
			c.DispatchEvent(event)
			return c
		}
	}
	return nil
}
func (this *Scene) RemoveAllChildren() {
	for _, c := range this.displayList {
		this.RemoveChild(c)
	}
	this.displayList = make([]core.IDisplayObject, 0)
}
func (this *Scene) GetChildByName(name string) core.IDisplayObject {
	for _, c := range this.displayList {
		if c.GetName() == name {
			return c
		}
	}
	return nil
}
func (this *Scene) GetChildrenNumber() int32 {
	return int32(len(this.displayList))
}
