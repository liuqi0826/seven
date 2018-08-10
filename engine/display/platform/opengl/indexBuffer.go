package opengl

import (
	"errors"

	"github.com/liuqi0826/seven/api/khronos/gl/gl"
)

type IndexBuffer struct {
	Index  uint32
	Length int
}

func (this *IndexBuffer) Upload(data []uint16) error {
	var err error
	this.Length = len(data)
	if this.Length > 65535 {
		err = errors.New("Data is large than 65535.")
	} else {
		gl.BufferData(gl.GL_ELEMENT_ARRAY_BUFFER, int64(this.Length*2), data, gl.GL_STATIC_DRAW)
	}
	return err
}
func (this *IndexBuffer) Dispose() {
	gl.DisableVertexAttribArray(this.Index)
}
