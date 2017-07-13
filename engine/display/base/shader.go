package base

import (
	"github.com/gonutz/d3d9"
)

type ShaderProgram struct {
	ID        string
	UsedCount int32
	Uploaded  bool

	VertexCode   []byte
	FragmentCode []byte

	VertexShader   d3d9.VertexShader
	FragmentShader d3d9.PixelShader
}

func (this *ShaderProgram) ShaderProgram(id string, vertex []byte, fragment []byte) {
	this.ID = id
	this.VertexCode = vertex
	this.FragmentCode = fragment
}
func (this *ShaderProgram) Upload(context d3d9.Device) {
	//this.VertexShader, err := context.CreateVertexShaderFromBytes(this.VertexCode)
	//this.FragmentShader, err := context.CreatePixelShaderFromBytes(this.FragmentCode)
	this.Uploaded = true
}
func (this *ShaderProgram) Dispose() {
	this.Uploaded = false
}

type ShaderValue struct {
	Shader *ShaderProgram

	VertexData   []byte
	FragmentData []byte
}

func (this *ShaderValue) ShaderValue() {
	this.VertexData = make([]byte, 0)
	this.FragmentData = make([]byte, 0)
}
func (this *ShaderValue) UploadValue() {
}
