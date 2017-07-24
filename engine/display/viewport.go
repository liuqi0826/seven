package display

import (
	"github.com/liuqi0826/seven/api/khronos/es3/gl"
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
func (this *Viewport) Frame() {
	gl.Clear(0)

	this.Camera.Update()
	this.Renderer.Render()
}
