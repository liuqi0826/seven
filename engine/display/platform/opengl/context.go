package opengl

import (
	"fmt"

	"github.com/liuqi0826/seven/api/khronos/gl/gl"
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/events"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

const (
	ALWAYS        = "always"
	EQUAL         = "equal"
	GREATER       = "greater"
	GREATER_EQUAL = "greaterEqual"
	LESS          = "less"
	LESS_EQUAL    = "lessEqual"
	NEVER         = "never"
	NOT_EQUAL     = "notEqual"
)

const (
	BYTES_4 = "bytes4"
	FLOAT_1 = "float1"
	FLOAT_2 = "float2"
	FLOAT_3 = "float3"
	FLOAT_4 = "float4"
)

type Context struct {
	events.EventDispatcher

	window               *glfw.Window
	currentShaderProgram *Program3D

	ready bool
	debug bool

	clearRed   float32
	clearGreen float32
	clearBlue  float32
	clearAlpha float32

	depthMask       bool
	passCompareMode string
}

func (this *Context) Setup(window *glfw.Window, debug bool) error {
	var err error

	this.EventDispatcher.EventDispatcher(this)

	this.window = window
	this.debug = debug

	err = gl.Init()
	if err != nil {
		fmt.Println(err)
		return err
	}

	version := gl.GetString(gl.GL_VERSION)
	fmt.Println(version)

	gl.Enable(gl.GL_DEPTH_TEST)
	gl.DepthFunc(gl.GL_LESS)
	gl.CullFace(gl.GL_BACK)

	this.ready = true

	return err
}
func (this *Context) Clear(color bool, depth bool, stencil bool) {
	var mask uint32
	if color {
		mask = mask | gl.GL_COLOR_BUFFER_BIT
	}
	if depth {
		mask = mask | gl.GL_DEPTH_BUFFER_BIT
	}
	if stencil {
		mask = mask | gl.GL_STENCIL_BUFFER_BIT
	}
	if this.ready {
		gl.Clear(mask)
	}
}
func (this *Context) ConfigureBackBuffer() {
}
func (this *Context) CreateCubeTexture() {
}
func (this *Context) CreateProgram() platform.IProgram3D {
	return new(Program3D)
}
func (this *Context) CreateTexture() {
}
func (this *Context) CreateIndexBuffer() platform.IIndexBuffer {
	indexBuffer := new(IndexBuffer)
	gl.GenBuffers(1, &indexBuffer.Index)
	gl.BindBuffer(gl.GL_ELEMENT_ARRAY_BUFFER, indexBuffer.Index)
	//fmt.Println("CreateIndexBuffer", indexBuffer.Index, indexBuffer.Length)
	return indexBuffer
}
func (this *Context) CreateVertexBuffer() platform.IVertexBuffer {
	vertexBuffer := new(VertexBuffer)
	gl.GenBuffers(1, &vertexBuffer.Index)
	gl.BindBuffer(gl.GL_ARRAY_BUFFER, vertexBuffer.Index)
	return vertexBuffer
}
func (this *Context) Dispose() {
}
func (this *Context) Present() {
	this.window.SwapBuffers()
}
func (this *Context) DrawTriangles(indexBuffer platform.IIndexBuffer, firstIndex int32, numTriangles int32) {
	if indexBuf, ok := indexBuffer.(*IndexBuffer); ok {
		gl.BindVertexArray(indexBuf.Index)
		gl.DrawElements(gl.GL_TRIANGLES, numTriangles, gl.GL_UNSIGNED_SHORT, nil)
		gl.BindVertexArray(0)
	}
}
func (this *Context) SetBlendFactors() {
}
func (this *Context) SetClearColor(red float32, green float32, blue float32, alpha float32) {
	this.clearRed = red
	this.clearGreen = green
	this.clearBlue = blue
	this.clearAlpha = alpha
	if this.ready {
		gl.ClearColor(this.clearRed, this.clearGreen, this.clearBlue, this.clearAlpha)
	} else {
		fmt.Println("GL is not ready.")
	}
}
func (this *Context) SetColorMask() {
}
func (this *Context) SetCulling() {
}
func (this *Context) SetDepthTest(depthMask bool, passCompareMode string) {
	this.depthMask = depthMask
	this.passCompareMode = passCompareMode
	if this.ready {
		if this.depthMask {
			gl.Enable(gl.GL_DEPTH_TEST)
		} else {

		}
		switch this.passCompareMode {
		case ALWAYS:
			gl.DepthFunc(gl.GL_ALWAYS)
		case EQUAL:
			gl.DepthFunc(gl.GL_EQUAL)
		case GREATER:
			gl.DepthFunc(gl.GL_GREATER)
		case GREATER_EQUAL:
		case LESS:
			gl.DepthFunc(gl.GL_LESS)
		case LESS_EQUAL:
		case NEVER:
			gl.DepthFunc(gl.GL_NEVER)
		case NOT_EQUAL:
		}
	}
}
func (this *Context) SetProgram(program platform.IProgram3D) {
	if program3D, ok := program.(*Program3D); ok {
		this.currentShaderProgram = program3D
		gl.UseProgram(program3D.Index)
		fmt.Println(program3D.Index)
	}
}
func (this *Context) SetProgramConstantsFromByteArray() {
}
func (this *Context) SetProgramConstantsFromMatrix() {
}
func (this *Context) SetProgramConstantsFromVector() {
}
func (this *Context) SetRenderToBackBuffer() {
}
func (this *Context) SetRenderToTexture() {
}
func (this *Context) SetScissorRectangle() {
}
func (this *Context) SetStencilActions() {
}
func (this *Context) SetStencilReferenceValue() {
}
func (this *Context) SetTextureAt() {
}
func (this *Context) SetVertexBufferAt(value string, stride int32, bufferOffset int32, format string) {
	if this.currentShaderProgram != nil {
		var size int32
		var xtype int32

		switch format {
		case FLOAT_1:
			size = 1
			xtype = gl.GL_FLOAT
		case FLOAT_2:
			size = 2
			xtype = gl.GL_FLOAT
		case FLOAT_3:
			size = 3
			xtype = gl.GL_FLOAT
		case FLOAT_4:
			size = 4
			xtype = gl.GL_FLOAT
		case BYTES_4:
			size = 4
			xtype = gl.GL_BYTE
		}

		attrib := uint32(gl.GetAttribLocation(this.currentShaderProgram.Index, value))
		gl.EnableVertexAttribArray(attrib)
		gl.VertexAttribPointer(attrib, size, xtype, false, stride, bufferOffset)
	}
}
