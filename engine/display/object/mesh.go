package object

import (
	"github.com/liuqi0826/seven/engine/display/base"
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/events"
)

type Mesh struct {
	base.DisplayObject
	events.IEventDispatcher

	renderer core.IRenderer
	geometry []*base.SubGeometry
	material *base.Material
	shader   *base.ShaderProgram
}

func (this *Mesh) Mesh() {
	this.geometry = make([]*base.SubGeometry, 0)
}
func (this *Mesh) Render() {
	if this.renderer != nil {
		for _, sg := range this.geometry {
			this.renderer.Render(sg)
		}
	}
}
