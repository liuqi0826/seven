package base

import (
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/display/platform/opengl"
	"github.com/liuqi0826/seven/engine/display/resource"
	"github.com/liuqi0826/seven/engine/static"
)

type Material struct {
	ID string

	texture [8]platform.ITexture

	userCount          uint32
	uploadTextureCount int32
	uploaded           bool

	materialResource *resource.MaterialResource
}

func (this *Material) Material(materialResource *resource.MaterialResource) {
	if materialResource != nil {
		this.ID = materialResource.ID
		this.materialResource = materialResource
		this.texture = [8]platform.ITexture{}
	}
}
func (this *Material) Upload() {
	if this.materialResource != nil {
		for i, t := range this.materialResource.TextureList {
			switch static.API {
			case static.GL:
				if i < 8 {
					text := new(opengl.Texture)
					text.Upload(t.Texture, t.Type)
					this.texture[i] = text
				}
			}
		}
		this.uploaded = true
	}
}
func (this *Material) Bind() {
	for i, t := range this.texture {
		switch static.API {
		case static.GL:
			if t != nil {
				idx := int32(i)
				t.SetSlot(idx)
			}
		}
	}
}
func (this *Material) AddCount() {
	this.userCount++
}
func (this *Material) SubCount() {
	if this.userCount > 0 {
		this.userCount--
	}
}
func (this *Material) GetCount() uint32 {
	return this.userCount
}
func (this *Material) Dispose() {

}
