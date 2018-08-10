package opengl

import (
	"fmt"

	"github.com/liuqi0826/seven/api/khronos/gl/gl"
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform"
)

const (
	TRANSFORM = "transform"
	TEXCOORD  = "texcoord"
	SKELETON  = "skeleton"
)

//Viewport render handle
func ForwordRender() {

}

//==================== Default ====================

//Renderer
type DefaultRender struct {
	ready   bool
	camera  core.ICamera
	program *Program3D

	matrixUBOIndex         uint32
	matrixUBOID            uint32
	matrixUBOBlockSize     int32
	matrixUBOBindingPoint  uint32
	matrixUBOValues        []string
	matrixUBOValuesIndices []uint32
	matrixUBOValuesOffset  []int32

	valueTransform []float32
}

func (this *DefaultRender) Setup(camera core.ICamera, program platform.IProgram3D) {
	if camera != nil {
		this.camera = camera
	}
	if shader, ok := program.(*Program3D); ok {
		this.program = shader
	}
	this.binding()
}
func (this *DefaultRender) SetCamera(camera core.ICamera) {
	if camera != nil {
		this.camera = camera
		this.binding()
	}
}
func (this *DefaultRender) SetProgram(program platform.IProgram3D) {
	if shader, ok := program.(*Program3D); ok {
		this.program = shader
		this.binding()
	}
}
func (this *DefaultRender) binding() {
	if this.camera != nil && this.program != nil && this.program.IsReady() && this.program.Index != 0 {
		this.ready = true

		this.matrixUBOBindingPoint = 0
		this.matrixUBOBlockSize = 3 * 16 * 4

		gl.UseProgram(this.program.Index)

		this.matrixUBOIndex = gl.GetUniformBlockIndex(this.program.Index, "MatrixBlock")
		gl.UniformBlockBinding(this.program.Index, this.matrixUBOIndex, this.matrixUBOBindingPoint)

		gl.GenBuffers(1, &this.matrixUBOID)
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, this.matrixUBOID)
		gl.BufferData(gl.GL_UNIFORM_BUFFER, int64(this.matrixUBOBlockSize), nil, gl.GL_DYNAMIC_DRAW)
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, 0)
		gl.BindBufferRange(gl.GL_UNIFORM_BUFFER, this.matrixUBOBindingPoint, this.matrixUBOID, 0, int64(this.matrixUBOBlockSize))
	} else {
		this.ready = false
	}
}
func (this *DefaultRender) SetValue(title string, value []float32) {
	switch title {
	case TRANSFORM:
		this.valueTransform = value
	}
}
func (this *DefaultRender) Render(target core.IRenderable) {
	if this.ready && target != nil {
		gl.UseProgram(this.program.Index)

		projection := this.camera.GetProjectionMatrix().GetRaw()
		camera := this.camera.GetTransformMatrix().GetRaw()
		var transform [16]float32
		for i := 0; i < 16; i++ {
			transform[i] = this.valueTransform[i]
		}

		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, this.matrixUBOID)
		gl.BufferSubData(gl.GL_UNIFORM_BUFFER, 0, 16*4, &projection[0])
		gl.BufferSubData(gl.GL_UNIFORM_BUFFER, 16*4, 16*4, &camera[0])
		gl.BufferSubData(gl.GL_UNIFORM_BUFFER, 16*4*2, 16*4, &transform[0])
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, 0)

		vlist := target.GetVertexBuffer()
		for _, vb := range *vlist {
			if vertexBuffer, ok := vb.(*VertexBuffer); ok {
				gl.BindBuffer(gl.GL_ARRAY_BUFFER, vertexBuffer.Index)

				positionAttrib := uint32(gl.GetAttribLocation(this.program.Index, "position"))
				gl.EnableVertexAttribArray(positionAttrib)
				gl.VertexAttribPointer(positionAttrib, 3, gl.GL_FLOAT, false, 8*4, 0)

				texcoordAttrib := uint32(gl.GetAttribLocation(this.program.Index, "texcoord"))
				gl.EnableVertexAttribArray(texcoordAttrib)
				gl.VertexAttribPointer(texcoordAttrib, 2, gl.GL_FLOAT, false, 8*4, 3*4)

				normalAttrib := uint32(gl.GetAttribLocation(this.program.Index, "normal"))
				gl.EnableVertexAttribArray(normalAttrib)
				gl.VertexAttribPointer(normalAttrib, 3, gl.GL_FLOAT, false, 8*4, 5*4)
			}
		}

		index := target.GetIndexBuffer()
		if indexBuffer, ok := index.(*IndexBuffer); ok {
			gl.BindBuffer(gl.GL_ELEMENT_ARRAY_BUFFER, indexBuffer.Index)
			gl.DrawElements(gl.GL_TRIANGLES, int32(indexBuffer.Length*4), gl.GL_UNSIGNED_SHORT, nil)
		}

		gl.UseProgram(0)
		gl.BindBuffer(gl.GL_ELEMENT_ARRAY_BUFFER, 0)
		gl.BindBuffer(gl.GL_ARRAY_BUFFER, 0)
	}
}

//==================== Stylize ====================

//Renderer
type StylizeRender struct {
	ready   bool
	camera  core.ICamera
	program *Program3D

	matrixUBOIndex         uint32
	matrixUBOID            uint32
	matrixUBOBlockSize     int32
	matrixUBOBindingPoint  uint32
	matrixUBOValues        []string
	matrixUBOValuesIndices []uint32
	matrixUBOValuesOffset  []int32

	valueTransform []float32
}

func (this *StylizeRender) Setup(camera core.ICamera, program platform.IProgram3D) {
	if camera != nil {
		this.camera = camera
	}
	if shader, ok := program.(*Program3D); ok {
		this.program = shader
	}
	this.binding()
}
func (this *StylizeRender) SetCamera(camera core.ICamera) {
	if camera != nil {
		this.camera = camera
		this.binding()
	}
}
func (this *StylizeRender) SetProgram(program platform.IProgram3D) {
	if shader, ok := program.(*Program3D); ok {
		this.program = shader
		this.binding()
	}
}
func (this *StylizeRender) binding() {
	if this.camera != nil && this.program != nil && this.program.IsReady() && this.program.Index != 0 {
		this.ready = true

		this.matrixUBOBindingPoint = 0
		this.matrixUBOBlockSize = 3 * 16 * 4

		gl.UseProgram(this.program.Index)

		this.matrixUBOIndex = gl.GetUniformBlockIndex(this.program.Index, "MatrixBlock")
		gl.UniformBlockBinding(this.program.Index, this.matrixUBOIndex, this.matrixUBOBindingPoint)

		gl.GenBuffers(1, &this.matrixUBOID)
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, this.matrixUBOID)
		gl.BufferData(gl.GL_UNIFORM_BUFFER, int64(this.matrixUBOBlockSize), nil, gl.GL_DYNAMIC_DRAW)
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, 0)
		gl.BindBufferRange(gl.GL_UNIFORM_BUFFER, this.matrixUBOBindingPoint, this.matrixUBOID, 0, int64(this.matrixUBOBlockSize))
	} else {
		this.ready = false
	}
}
func (this *StylizeRender) SetValue(title string, value []float32) {
	switch title {
	case TRANSFORM:
		this.valueTransform = value
	}
}
func (this *StylizeRender) Render(target core.IRenderable) {
	if this.ready && target != nil {
		gl.UseProgram(this.program.Index)

		projection := this.camera.GetProjectionMatrix().GetRaw()
		camera := this.camera.GetTransformMatrix().GetRaw()
		var transform [16]float32
		for i := 0; i < 16; i++ {
			transform[i] = this.valueTransform[i]
		}

		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, this.matrixUBOID)
		gl.BufferSubData(gl.GL_UNIFORM_BUFFER, 0, 16*4, &projection[0])
		gl.BufferSubData(gl.GL_UNIFORM_BUFFER, 16*4, 16*4, &camera[0])
		gl.BufferSubData(gl.GL_UNIFORM_BUFFER, 16*4*2, 16*4, &transform[0])
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, 0)

		vlist := target.GetVertexBuffer()
		for _, vb := range *vlist {
			if vertexBuffer, ok := vb.(*VertexBuffer); ok {
				gl.BindBuffer(gl.GL_ARRAY_BUFFER, vertexBuffer.Index)

				positionAttrib := uint32(gl.GetAttribLocation(this.program.Index, "position"))
				gl.EnableVertexAttribArray(positionAttrib)
				gl.VertexAttribPointer(positionAttrib, 3, gl.GL_FLOAT, false, 8*4, 0)

				texcoordAttrib := uint32(gl.GetAttribLocation(this.program.Index, "texcoord"))
				gl.EnableVertexAttribArray(texcoordAttrib)
				gl.VertexAttribPointer(texcoordAttrib, 2, gl.GL_FLOAT, false, 8*4, 3*4)

				normalAttrib := uint32(gl.GetAttribLocation(this.program.Index, "normal"))
				gl.EnableVertexAttribArray(normalAttrib)
				gl.VertexAttribPointer(normalAttrib, 3, gl.GL_FLOAT, false, 8*4, 5*4)
			}
		}

		index := target.GetIndexBuffer()
		if indexBuffer, ok := index.(*IndexBuffer); ok {
			gl.BindBuffer(gl.GL_ELEMENT_ARRAY_BUFFER, indexBuffer.Index)
			gl.DrawElements(gl.GL_TRIANGLES, int32(indexBuffer.Length*4), gl.GL_UNSIGNED_SHORT, nil)
		}

		gl.UseProgram(0)
		gl.BindBuffer(gl.GL_ELEMENT_ARRAY_BUFFER, 0)
		gl.BindBuffer(gl.GL_ARRAY_BUFFER, 0)
	}
}

//==================== Animation Matrix4x4 ====================

type AnimationRender struct {
	ready   bool
	camera  core.ICamera
	program *Program3D

	matrixUBOIndex        uint32
	matrixUBOBindingPoint uint32
	jointUBOIndex         uint32
	jointUBOID            uint32

	valueTransform []float32
	valueTexcoord  []float32
	valueSkeleton  []float32
}

func (this *AnimationRender) Setup(camera core.ICamera, program platform.IProgram3D) {
	if camera != nil {
		this.camera = camera
	}
	if shader, ok := program.(*Program3D); ok {
		this.program = shader
	}
	fmt.Println(this.camera, this.program)
	if this.camera != nil && this.program != nil {
		this.ready = true

		gl.UseProgram(this.program.Index)

		this.matrixUBOIndex = gl.GetUniformBlockIndex(this.program.Index, "MatrixBlock")
		gl.UniformBlockBinding(this.program.Index, this.matrixUBOIndex, 0)

		gl.GenBuffers(1, &this.matrixUBOBindingPoint)
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, this.matrixUBOBindingPoint)
		gl.BufferData(gl.GL_UNIFORM_BUFFER, 2*16*4, nil, gl.GL_DYNAMIC_DRAW)
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, 0)
		gl.BindBufferRange(gl.GL_UNIFORM_BUFFER, 0, this.matrixUBOBindingPoint, 0, 2*16*4)

		gl.UseProgram(0)

		//this.jointUBOIndex = gl.GetUniformBlockIndex(this.program.Index, gl.Str("JointBlock\x00"))
		//gl.UniformBlockBinding(this.program.Index, this.jointUBOIndex, 0)
		//
		//gl.GenBuffers(1, &this.jointUBOID)
		//gl.BindBuffer(gl.UNIFORM_BUFFER, this.jointUBOID)
		//gl.BufferData(gl.UNIFORM_BUFFER, 100*16*4, nil, gl.DYNAMIC_DRAW)
		//gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
		//gl.BindBufferRange(gl.UNIFORM_BUFFER, 0, this.jointUBOID, 0, 100*16*4)

		fmt.Println("shader is ready!")
	} else {
		this.ready = false

		fmt.Println("shader is not ready!")
	}
}
func (this *AnimationRender) SetCamera(camera core.ICamera) {
	if camera != nil {
		this.camera = camera
	}
	fmt.Println(this.camera, this.program)
	if this.camera != nil && this.program != nil {
		this.ready = true

		gl.UseProgram(this.program.Index)

		this.matrixUBOIndex = gl.GetUniformBlockIndex(this.program.Index, "MatrixBlock")
		gl.UniformBlockBinding(this.program.Index, this.matrixUBOIndex, 0)

		gl.GenBuffers(1, &this.matrixUBOBindingPoint)
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, this.matrixUBOBindingPoint)
		gl.BufferData(gl.GL_UNIFORM_BUFFER, 2*16*4, nil, gl.GL_DYNAMIC_DRAW)
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, 0)
		gl.BindBufferRange(gl.GL_UNIFORM_BUFFER, 0, this.matrixUBOBindingPoint, 0, 2*16*4)

		gl.UseProgram(0)

		//this.jointUBOIndex = gl.GetUniformBlockIndex(this.program.Index, gl.Str("JointBlock\x00"))
		//gl.UniformBlockBinding(this.program.Index, this.jointUBOIndex, 0)
		//
		//gl.GenBuffers(1, &this.jointUBOID)
		//gl.BindBuffer(gl.UNIFORM_BUFFER, this.jointUBOID)
		//gl.BufferData(gl.UNIFORM_BUFFER, 100*16*4, nil, gl.DYNAMIC_DRAW)
		//gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
		//gl.BindBufferRange(gl.UNIFORM_BUFFER, 0, this.jointUBOID, 0, 100*16*4)

		fmt.Println("shader is ready!")
	} else {
		this.ready = false

		fmt.Println("shader is not ready!")
	}
	fmt.Println(this.ready)
}
func (this *AnimationRender) SetProgram(program platform.IProgram3D) {
	if shader, ok := program.(*Program3D); ok {
		this.program = shader
	}
	fmt.Println(this.camera, this.program)
	if this.camera != nil && this.program != nil {
		this.ready = true

		gl.UseProgram(this.program.Index)

		this.matrixUBOIndex = gl.GetUniformBlockIndex(this.program.Index, "MatrixBlock")
		gl.UniformBlockBinding(this.program.Index, this.matrixUBOIndex, 0)

		gl.GenBuffers(1, &this.matrixUBOBindingPoint)
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, this.matrixUBOBindingPoint)
		gl.BufferData(gl.GL_UNIFORM_BUFFER, 2*16*4, nil, gl.GL_DYNAMIC_DRAW)
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, 0)
		gl.BindBufferRange(gl.GL_UNIFORM_BUFFER, 0, this.matrixUBOBindingPoint, 0, 2*16*4)

		gl.UseProgram(0)

		//this.jointUBOIndex = gl.GetUniformBlockIndex(this.program.Index, gl.Str("JointBlock\x00"))
		//gl.UniformBlockBinding(this.program.Index, this.jointUBOIndex, 0)
		//
		//gl.GenBuffers(1, &this.jointUBOID)
		//gl.BindBuffer(gl.UNIFORM_BUFFER, this.jointUBOID)
		//gl.BufferData(gl.UNIFORM_BUFFER, 100*16*4, nil, gl.DYNAMIC_DRAW)
		//gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
		//gl.BindBufferRange(gl.UNIFORM_BUFFER, 0, this.jointUBOID, 0, 100*16*4)

		fmt.Println("shader is ready!")
	} else {
		this.ready = false

		fmt.Println("shader is not ready!")
	}
}
func (this *AnimationRender) SetValue(title string, value []float32) {
	switch title {
	case TRANSFORM:
		this.valueTransform = value
	case TEXCOORD:
		this.valueTexcoord = value
	case SKELETON:
		this.valueSkeleton = value
	}
}
func (this *AnimationRender) Render(target core.IRenderable) {
	if this.ready && target != nil {
		gl.UseProgram(this.program.Index)

		projection := this.camera.GetProjectionMatrix().ToArray()
		camera := this.camera.GetTransformMatrix().ToArray()
		matrixValue := make([]float32, 0)
		for _, p := range projection {
			matrixValue = append(matrixValue, p)
		}
		for _, c := range camera {
			matrixValue = append(matrixValue, c)
		}
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, this.matrixUBOBindingPoint)
		gl.BufferSubData(gl.GL_UNIFORM_BUFFER, 0, 2*16*4, matrixValue)
		gl.BindBuffer(gl.GL_UNIFORM_BUFFER, 0)

		if this.valueTransform != nil && len(this.valueTransform) >= 16 {
			transformUniform := gl.GetUniformLocation(this.program.Index, "transform")
			gl.UniformMatrix4fv(transformUniform, 1, false, &this.valueTransform[0])
		}

		if this.valueTexcoord != nil && len(this.valueTexcoord) >= 2 {

		}

		if this.valueSkeleton != nil && len(this.valueSkeleton) >= 2 {

		}

		vlist := target.GetVertexBuffer()
		for _, vb := range *vlist {
			if vertexBuffer, ok := vb.(*VertexBuffer); ok {
				gl.BindBuffer(gl.GL_ARRAY_BUFFER, vertexBuffer.Index)

				positionAttrib := uint32(gl.GetAttribLocation(this.program.Index, "position"))
				gl.EnableVertexAttribArray(positionAttrib)
				gl.VertexAttribPointer(positionAttrib, 3, gl.GL_FLOAT, false, 8*4, 0)

				texcoordAttrib := uint32(gl.GetAttribLocation(this.program.Index, "texcoord"))
				gl.EnableVertexAttribArray(texcoordAttrib)
				gl.VertexAttribPointer(texcoordAttrib, 2, gl.GL_FLOAT, false, 8*4, 3*4)

				normalAttrib := uint32(gl.GetAttribLocation(this.program.Index, "normal"))
				gl.EnableVertexAttribArray(normalAttrib)
				gl.VertexAttribPointer(normalAttrib, 3, gl.GL_FLOAT, false, 8*4, 5*4)
			}
		}

		index := target.GetIndexBuffer()
		if indexBuffer, ok := index.(*IndexBuffer); ok {
			gl.BindBuffer(gl.GL_ELEMENT_ARRAY_BUFFER, indexBuffer.Index)
			gl.DrawElements(gl.GL_TRIANGLES, int32(indexBuffer.Length*4), gl.GL_UNSIGNED_SHORT, nil)
		}

		gl.UseProgram(0)
		gl.BindBuffer(gl.GL_ELEMENT_ARRAY_BUFFER, 0)
		gl.BindBuffer(gl.GL_ARRAY_BUFFER, 0)
	}
}
