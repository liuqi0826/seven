package es

import (
	"errors"

	gl "github.com/go-gl/gl/v3.1/gles2"
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
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, this.Length*2, gl.Ptr(data), gl.STATIC_DRAW)
	}
	return err
}
func (this *IndexBuffer) Dispose() {
	gl.DisableVertexAttribArray(this.Index)
}
