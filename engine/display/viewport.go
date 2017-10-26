package display

import (
	//"fmt"

	"github.com/liuqi0826/seven/engine/global"
	"github.com/liuqi0826/seven/engine/static"
)

type Viewport struct {
	Scene *Scene

	render func()
	width  uint32
	height uint32
}

func (this *Viewport) Viewport(width uint32, height uint32, rendingType string) {
	this.width = width
	this.height = height

	this.Scene = new(Scene)
	this.Scene.Scene()

	switch rendingType {
	case static.FORWARD:
		this.render = this.forword
	case static.DEFERRED:
		this.render = this.deferred
	default:
		this.render = this.forword
	}
}
func (this *Viewport) Frame() {
	this.Scene.camera.Update()
	this.render()
}
func (this *Viewport) GetWidth() uint32 {
	return this.width
}
func (this *Viewport) SetWidth(width uint32) {
	this.width = width
}
func (this *Viewport) GetHeight() uint32 {
	return this.height
}
func (this *Viewport) SetHeight(height uint32) {
	this.height = height
}
func (this *Viewport) SetRender(render func()) {
	this.render = render
}
func (this *Viewport) forword() {
	global.Context3D.Clear(true, true, true)

	for _, do := range this.Scene.camera.DisplayList {
		do.Update(nil)
		do.Render()
	}

	global.Context3D.Present()
}
func (this *Viewport) deferred() {
	global.Context3D.Clear(true, true, true)

	global.Context3D.Present()

}
