package core

import (
	"github.com/liuqi0826/seven/geom"
)

//++++++++++++++++++++ Object ++++++++++++++++++++

type Object struct {
	geom.Vector4

	ID   string
	Name string
}

func (this *Object) Object() {
	this.Vector4.Vector4()
}

//++++++++++++++++++++ DisplayObject ++++++++++++++++++++

type DisplayObject struct {
	Object
}

func (this *DisplayObject) DisplayObject() {
	this.Object.Object()
}
