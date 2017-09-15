package es

import (
	gl "github.com/go-gl/gl/v3.1/gles2"

	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform"
)

type DefaultRender struct {
	program *Program3D
}

func (this *DefaultRender) Setup(program platform.IProgram3D) {
	if shader, ok := program.(*Program3D); ok {
		this.program = shader
	}
}
func (this *DefaultRender) Render(target core.IRenderable) {
	if target != nil {
		gl.UseProgram(this.program.Index)
		vlist := target.GetVertexBuffer()
		for _, vb := range *vlist {
			if vertexBuffer, ok := vb.(*VertexBuffer); ok {
				gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer.Index)
				positionAttrib := uint32(gl.GetAttribLocation(this.program.Index, gl.Str("position\x00")))
				gl.EnableVertexAttribArray(positionAttrib)
				gl.VertexAttribPointer(positionAttrib, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
			}
		}

		index := target.GetIndexBuffer()
		if indexBuffer, ok := index.(*IndexBuffer); ok {
			gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBuffer.Index)
			gl.DrawElements(gl.TRIANGLES, int32(indexBuffer.Length*4), gl.UNSIGNED_SHORT, nil)
		} else {
			return
		}

		gl.UseProgram(0)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
		gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	}
}
