package object

import (
	"github.com/liuqi0826/seven/engine/display/base"
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/display/render"
	"github.com/liuqi0826/seven/geom"
)

type Mesh struct {
	base.DisplayObject

	renderer core.IRenderer
	geometry []*base.SubGeometry
	material *base.Material
	shader   platform.IProgram3D
}

func (this *Mesh) Mesh(geometry []*base.SubGeometry, material *base.Material, shader platform.IProgram3D) {
	this.DisplayObject.DisplayObject()

	this.geometry = geometry
	this.material = material
	this.shader = shader

	this.renderer = render.CreateRender("default")
	if this.renderer != nil {
		this.renderer.Setup(this.shader)
	}
}
func (this *Mesh) Update(transform *geom.Matrix4x4) {
	this.Object.Update()
	this.Object.GetTransform().Append(transform)
	for _, v := range this.geometry {
		v.Update(this.Object.GetTransform())
	}
}
func (this *Mesh) Render(projection *geom.Matrix4x4) {
	this.DisplayObject.Render(projection)
	if this.renderer != nil {
		for _, sg := range this.geometry {
			sg.SetProjection(projection)
			this.renderer.Render(sg)
		}
	}
}
