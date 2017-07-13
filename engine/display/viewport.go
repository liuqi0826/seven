package display

import (
	"github.com/gonutz/d3d9"
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/render"
)

type Viewport struct {
	Camera   *Camera
	Scene    *Scene
	Renderer core.IRenderer

	Width  int32
	Height int32
}

func (this *Viewport) Viewport(width int32, height int32) {
	this.Width = width
	this.Height = height
	cameraConfig := new(ProjectionConfig)
	cameraConfig.ProjectType = PROJECTION_TYPE_PERSPECTIVE
	cameraConfig.NearClipping = 0.1
	cameraConfig.FarClipping = 1000.0
	cameraConfig.Field = 45.0
	cameraConfig.CoordinateSystem = COORDINATE_SYSTEM_LEFT_HAND
	this.Camera = new(Camera)
	this.Camera.Camera(this, cameraConfig)

	this.Scene = new(Scene)
	this.Scene.Scene()

	this.Renderer = new(render.DefaultRenderer)
}
func (this *Viewport) Frame(context d3d9.Device) {
	context.Clear(nil, d3d9.CLEAR_TARGET, 128, 1, 0)
	context.BeginScene()

	this.Camera.Update()
	this.Renderer.Render()

	context.EndScene()
	context.Present(nil, nil, nil, nil)
}
