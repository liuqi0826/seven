package base

import (
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/seven/geom"
)

type DisplayObjectContainer struct {
	DisplayObject

	displayList []core.IDisplayObject
}

func (this *DisplayObjectContainer) DisplayObjectContainer() {
	this.DisplayObject.DisplayObject()
	this.displayList = make([]core.IDisplayObject, 0)
}
func (this *DisplayObjectContainer) AddChild(displayObject core.IDisplayObject) {
	this.displayList = append(this.displayList, displayObject)

	displayObject.SetRoot(this.GetRoot())
	displayObject.SetParent(this)

	event := new(events.Event)
	event.Type = events.ADDED
	displayObject.DispatchEvent(event)
}
func (this *DisplayObjectContainer) RemoveChild(displayObject core.IDisplayObject) core.IDisplayObject {
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
func (this *DisplayObjectContainer) RemoveAllChildren() {
	for _, c := range this.displayList {
		this.RemoveChild(c)
	}
	this.displayList = make([]core.IDisplayObject, 0)
}
func (this *DisplayObjectContainer) GetChildByName(name string) core.IDisplayObject {
	for _, c := range this.displayList {
		if c.GetName() == name {
			return c
		}
	}
	return nil
}
func (this *DisplayObjectContainer) SetRoot(root core.IContainer) {
	this.root = root
	for _, c := range this.displayList {
		c.SetRoot(this.root)
	}
}
func (this *DisplayObjectContainer) GetChildrenNumber() int32 {
	return int32(len(this.displayList))
}
func (this *DisplayObjectContainer) Render(projection *geom.Matrix4x4) {
	for _, c := range this.displayList {
		c.Render(projection)
	}
}
