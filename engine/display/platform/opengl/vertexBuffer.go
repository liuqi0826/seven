package opengl

import (
	"github.com/liuqi0826/seven/api/khronos/gl/gl"
)

type VertexBuffer struct {
	Index uint32
}

func (this *VertexBuffer) Upload(data []float32) error {
	var err error
	gl.BufferData(gl.GL_ARRAY_BUFFER, int64(len(data)*4), data, gl.GL_STATIC_DRAW)
	return err
}
func (this *VertexBuffer) Dispose() {
	gl.DisableVertexAttribArray(this.Index)
}
