package engine

import (
	"github.com/liuqi0826/seven/engine/display/base"
	"github.com/liuqi0826/seven/engine/display/resource"
)

type ResourceManager struct {
	geometryResource map[string]*resource.GeometryResource
	materialResource map[string]*resource.MaterialResource
	//animationResource map[string]*resource.MaterialResource
	shaderResource map[string]*resource.ShaderResource

	geometryRuntime map[string]*base.SubGeometry
	materialRuntime map[string]*base.Material
	shaderRuntime   map[string]*base.ShaderProgram
}

func (this *ResourceManager) ResourceManager() {
	this.geometryResource = make(map[string]*resource.GeometryResource)
	this.materialResource = make(map[string]*resource.MaterialResource)
	this.shaderResource = make(map[string]*resource.ShaderResource)

	this.geometryRuntime = make(map[string]*base.SubGeometry)
	this.materialRuntime = make(map[string]*base.Material)
	this.shaderRuntime = make(map[string]*base.ShaderProgram)
}
func (this *ResourceManager) ParserGeometrie(value []byte) string {
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
}
func (this *ResourceManager) GetGeometrie(id string) *resource.GeometryResource {
}
func (this *ResourceManager) AddMaterial(value *resource.MaterialResource) {
}
func (this *ResourceManager) GetMaterial(id string) *resource.MaterialResource {
}

//func (this *ResourceManager) AddAnimation(value *resource.MaterialResource) {
//}
//func (this *ResourceManager) GetAnimation(id string) *resource.MaterialResource {
//}
func (this *ResourceManager) AddShader(value *resource.ShaderResource) {
}
func (this *ResourceManager) GetShader(id string) *resource.ShaderResource {
}
