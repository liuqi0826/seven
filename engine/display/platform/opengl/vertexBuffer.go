package opengl

import (
	"github.com/go-gl/gl/v4.5-core/gl"
)

type VertexBuffer struct {
	Index uint32
}

func (this *VertexBuffer) Upload(data []float32) error {
	var err error
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	return err
}
func (this *VertexBuffer) Dispose() {
	gl.DisableVertexAttribArray(this.Index)
}
