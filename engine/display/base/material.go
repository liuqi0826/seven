package base

import (
	"github.com/liuqi0826/seven/engine/display/platform"
)

type Material struct {
	ID        string
	UsedCount int32

	texture []platform.ITexture

	uploadTextureCount int32
	uploaded           bool
}

func (this *Material) Material() {
	this.texture = make([]platform.ITexture, 0)
}
func (this *Material) Upload() {

}
func (this *Material) Dispose() {

}
