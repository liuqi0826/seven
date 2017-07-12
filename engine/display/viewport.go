package display

import (
	"github.com/liuqi0826/seven/engine/display/core"
)

type Viewport struct {
	Camera *Camera
	Scene  *Scene
	Render core.IRender

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
}
func (this *Viewport) Frame() {

}
