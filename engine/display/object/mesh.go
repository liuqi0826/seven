package object

import (
	"fmt"

	"github.com/liuqi0826/seven/engine/display/base"
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/display/render"
)

type Mesh struct {
	base.DisplayObject

	renderer core.IRenderer
	geometry []*base.SubGeometry
	material *base.Material
	shader   platform.IProgram3D
}

func (this *Mesh) Mesh(geometry []*base.SubGeometry, material *base.Material, shader platform.IProgram3D) {
	this.geometry = geometry
	this.material = material
	this.shader = shader

	this.renderer = render.CreateRender("default")
	if this.renderer != nil {
		this.renderer.Setup(this.shader)
	}
}
func (this *Mesh) Render() {
	if this.renderer != nil {
		for _, sg := range this.geometry {
			fmt.Println(sg.ID)
			this.renderer.Render(sg)
		}
	}
}
