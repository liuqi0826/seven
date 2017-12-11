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
		this.renderer.Setup(this.GetCamera(), this.shader)
	}
}
func (this *Mesh) SetCamera(camera core.ICamera) {
	this.DisplayObject.SetCamera(camera)
	if this.renderer != nil {
		this.renderer.SetCamera(camera)
	}
}
func (this *Mesh) Update(transform *geom.Matrix4x4) {
	this.Object.Update()
	this.Object.GetTransform().Append(transform)

	this.renderer.SetValue("transform", this.Object.GetTransform().GetRawSlice())
}
func (this *Mesh) Render() {
	if this.renderer != nil {
		this.material.Bind()
		for _, sg := range this.geometry {
			this.renderer.Render(sg)
		}
	}
}
