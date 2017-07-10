package display

import (
	"github.com/liuqi0826/seven/engine/display/core"
)

type Viewport struct {
	Camera core.ICamera
	Scene  core.IContainer
	Render core.IRender
}

func (this *Viewport) Viewport() {

}
func (this *Viewport) Frame() {

}
