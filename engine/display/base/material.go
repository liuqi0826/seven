package base

import (
	"github.com/gonutz/d3d9"
)

type Material struct {
	ID        string
	UsedCount int32

	texture []d3d9.Texture

	uploadTextureCount int32
	uploaded           bool
}

func (this *Material) Material() {
	this.texture = make([]d3d9.Texture, 0)
}
func (this *Material) Upload() {

}
func (this *Material) Dispose() {

}
