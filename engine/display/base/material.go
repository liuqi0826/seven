package base

import (
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/display/platform/opengl"
	"github.com/liuqi0826/seven/engine/display/resource"
	"github.com/liuqi0826/seven/engine/static"
)

type Material struct {
	ID string

	texture []platform.ITexture

	userCount          uint32
	uploadTextureCount int32
	uploaded           bool

	materialResource *resource.MaterialResource
}

func (this *Material) Material(materialResource *resource.MaterialResource) {
	if materialResource != nil {
		this.ID = materialResource.ID
		this.materialResource = materialResource
		this.texture = make([]platform.ITexture, 0)
	}
}
func (this *Material) Upload() {
	if this.materialResource != nil {
		for _, t := range this.materialResource.TextureList {
			switch static.API {
			case static.GL:
				text := new(opengl.Texture)
				text.Upload(t.Texture, t.Type)

				this.texture = append(this.texture, text)
			}
		}
		this.uploaded = true
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
