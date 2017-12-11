package object

import (
	"github.com/liuqi0826/seven/engine/display/base"
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/display/render"
	"github.com/liuqi0826/seven/geom"
)

type Sprite struct {
	Mesh

	animation []core.IAnimation
}

func (this *Sprite) Sprite(geometry []*base.SubGeometry, material *base.Material, shader platform.IProgram3D) {
	this.DisplayObject.DisplayObject()

	this.geometry = geometry
	this.material = material
	this.shader = shader

	this.renderer = render.CreateRender("animation")
	if this.renderer != nil {
		this.renderer.Setup(this.GetCamera(), this.shader)
	}

	this.animation = make([]core.IAnimation, 0)
}
func (this *Sprite) AddAnimation(anim core.IAnimation) {
	this.animation = append(this.animation, anim)
}
func (this *Sprite) Update(transform *geom.Matrix4x4) {
	this.Mesh.Update(transform)

	for _, an := range this.animation {
		an.Update()
		switch an.GetType() {
		case "skeleton":
		case "texcoord":
		}
	}
}
