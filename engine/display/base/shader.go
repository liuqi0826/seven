package base

import (
	"fmt"

	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/display/resource"
)

type ShaderProgram struct {
	ID        string
	UsedCount int32
	Uploaded  bool

	Program  platform.IProgram3D
	Resource *resource.ShaderResource
}

func (this *ShaderProgram) ShaderProgram(id string, Resource *resource.ShaderResource) {
	this.ID = id
	this.Resource = Resource
}
func (this *ShaderProgram) Upload(context core.IContext) {
	this.Program = context.CreateProgram()
	err := this.Program.Upload(this.Resource.Vertex, this.Resource.Fragment)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		this.Uploaded = true
	}
}
func (this *ShaderProgram) Dispose() {
	this.Program.Dispose()
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
