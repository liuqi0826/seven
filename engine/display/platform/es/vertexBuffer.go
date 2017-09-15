package es

import (
	gl "github.com/go-gl/gl/v3.1/gles2"
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
