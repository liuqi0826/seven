package base

import (
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/seven/geom"
)

//++++++++++++++++++++ Object ++++++++++++++++++++

type Object struct {
	events.EventDispatcher

	id   string
	name string

	Position geom.Vector4
	Rotation geom.Vector4
	Scale    geom.Vector4
}

func (this *Object) Object() {
	this.EventDispatcher.EventDispatcher()

	this.Position = geom.Vector4{}
	this.Rotation = geom.Vector4{}
	this.Scale = geom.Vector4{}
}
func (this *Object) GetID() string {
	return this.id
}
func (this *Object) GetName() string {
	return this.name
}
func (this *Object) SetName(name string) {
	this.name = name
}
func (this *Object) Update() {

}

//++++++++++++++++++++ DisplayObject ++++++++++++++++++++

type DisplayObject struct {
	Object

	root      core.IContainer
	parent    core.IContainer
	layerMask int32
}

func (this *DisplayObject) DisplayObject() {
	this.Object.Object()
}
func (this *DisplayObject) GetRoot() core.IContainer {
	return this.root
}
func (this *DisplayObject) SetRoot(root core.IContainer) {
	this.root = root
}
func (this *DisplayObject) GetParent() core.IContainer {
	return this.parent
}
func (this *DisplayObject) SetParent(parent core.IContainer) {
	this.parent = parent
}
func (this *DisplayObject) GetLayerMask() int32 {
	return this.layerMask
}
func (this *DisplayObject) SetLayerMask(mask int32) {
	this.layerMask = mask
}
func (this *DisplayObject) Render() {
}
