package opengl

import (
	//"fmt"

	"github.com/go-gl/gl/v4.5-core/gl"

	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform"
)

//Viewport render handle
func ForwordRender() {

}

//Renderer
type DefaultRender struct {
	ready   bool
	program *Program3D
}

func (this *DefaultRender) Setup(program platform.IProgram3D) {
	if shader, ok := program.(*Program3D); ok {
		this.ready = true
		this.program = shader
	} else {
		this.ready = false
	}
}
func (this *DefaultRender) Render(target core.IRenderable) {
	if this.ready && target != nil {
		gl.UseProgram(this.program.Index)

		value := target.GetValueBuffer()
		if len(value) >= 32 {
			projectionUniform := gl.GetUniformLocation(this.program.Index, gl.Str("projection\x00"))
			gl.UniformMatrix4fv(projectionUniform, 1, false, &value[0])
			transformUniform := gl.GetUniformLocation(this.program.Index, gl.Str("transform\x00"))
			gl.UniformMatrix4fv(transformUniform, 1, false, &value[16])
		}

		vlist := target.GetVertexBuffer()
		for _, vb := range *vlist {
			if vertexBuffer, ok := vb.(*VertexBuffer); ok {
				gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer.Index)

				positionAttrib := uint32(gl.GetAttribLocation(this.program.Index, gl.Str("position\x00")))
				gl.EnableVertexAttribArray(positionAttrib)
				gl.VertexAttribPointer(positionAttrib, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))

				texcoordAttrib := uint32(gl.GetAttribLocation(this.program.Index, gl.Str("texcoord\x00")))
				gl.EnableVertexAttribArray(texcoordAttrib)
				gl.VertexAttribPointer(texcoordAttrib, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))

				normalAttrib := uint32(gl.GetAttribLocation(this.program.Index, gl.Str("normal\x00")))
				gl.EnableVertexAttribArray(normalAttrib)
				gl.VertexAttribPointer(normalAttrib, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(5*4))
			}
		}

		index := target.GetIndexBuffer()
		if indexBuffer, ok := index.(*IndexBuffer); ok {
			gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBuffer.Index)
			gl.DrawElements(gl.TRIANGLES, int32(indexBuffer.Length*4), gl.UNSIGNED_SHORT, nil)
		}

		gl.UseProgram(0)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
		gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	}
}
