package engine

import (
	"encoding/json"

	"github.com/liuqi0826/seven/engine/display/base"
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/resource"
)

type ResourceManager struct {
	context core.IContext

	//静态资源库
	geometryResource  map[string]*resource.GeometryResource
	materialResource  map[string]*resource.MaterialResource
	animationResource map[string]*resource.AnimationResource
	shaderResource    map[string]*resource.ShaderResource

	//运行时资源
	geometryRuntime map[string]*base.SubGeometry
	materialRuntime map[string]*base.Material
	shaderRuntime   map[string]*base.ShaderProgram
}

func (this *ResourceManager) ResourceManager() {
	this.geometryResource = make(map[string]*resource.GeometryResource)
	this.materialResource = make(map[string]*resource.MaterialResource)
	this.animationResource = make(map[string]*resource.AnimationResource)
	this.shaderResource = make(map[string]*resource.ShaderResource)

	this.geometryRuntime = make(map[string]*base.SubGeometry)
	this.materialRuntime = make(map[string]*base.Material)
	this.shaderRuntime = make(map[string]*base.ShaderProgram)
}
func (this *ResourceManager) Setup(context core.IContext) error {
	var err error
	this.context = context
	return err
}
func (this *ResourceManager) ParserGeometrie(value []byte) string {
	gr := new(resource.GeometryResource)
	err := json.Unmarshal(value, gr)
	if err == nil {
		this.AddGeometrie(gr)
		return gr.ID
	}
	return ""
}
func (this *ResourceManager) ParserMaterial(value []byte) string {
	return ""
}
func (this *ResourceManager) ParserAnimation(value []byte) string {
	return ""
}
func (this *ResourceManager) ParserShader(value []byte) string {
	return ""
}

func (this *ResourceManager) AddGeometrie(value *resource.GeometryResource) {
	this.geometryResource[value.ID] = value
}
func (this *ResourceManager) GetGeometrie(id string) *resource.GeometryResource {
	if gr, ok := this.geometryResource[id]; ok {
		return gr
	}
	return nil
}
func (this *ResourceManager) AddMaterial(value *resource.MaterialResource) {
}
func (this *ResourceManager) GetMaterial(id string) *resource.MaterialResource {
	return nil
}
func (this *ResourceManager) AddAnimation(value *resource.AnimationResource) {
}
func (this *ResourceManager) GetAnimation(id string) *resource.AnimationResource {
	return nil
}
func (this *ResourceManager) AddShader(value *resource.ShaderResource) {
}
func (this *ResourceManager) GetShader(id string) *resource.ShaderResource {
	return nil
}

func (this *ResourceManager) CreateSubgeometrie(id string) *base.SubGeometry {
	gr := this.GetGeometrie(id)
	if gr != nil {
		sg := new(base.SubGeometry)
		sg.SubGeometry(gr)
		if !sg.Uploaded {
			sg.Upload(this.context)
		}
	}
	return nil
}
func (this *ResourceManager) CreateMaterial(id string) *base.Material {
	return nil
}
func (this *ResourceManager) CreateShaderProgram(id string) *base.ShaderProgram {
	return nil
}
