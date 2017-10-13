package display

import (
	"github.com/liuqi0826/seven/engine/global"
	"github.com/liuqi0826/seven/engine/static"
	"github.com/liuqi0826/seven/geom"
)

type Viewport struct {
	Camera *Camera
	Scene  *Scene

	render func()
	width  uint32
	height uint32
}

func (this *Viewport) Viewport(width uint32, height uint32, rendingType string) {
	this.width = width
	this.height = height

	this.Camera = new(Camera)
	this.Camera.Camera(this, nil)
	this.Camera.Position.X = 0.0
	this.Camera.Position.Y = 10.0
	this.Camera.Position.Z = 10.0
	this.Camera.LookAt(new(geom.Vector4), nil)

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
	this.Camera.Update()
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

	for _, do := range this.Camera.DisplayList {
		projection := this.Camera.GetProjectionMatrix()
		transform := this.Camera.GetTranformMatrix()

		if projection != nil && transform != nil {
			do.Update(transform)
		}

		do.Render(projection)
	}

	global.Context3D.Present()
}
func (this *Viewport) deferred() {
	global.Context3D.Clear(true, true, true)

	global.Context3D.Present()

}
