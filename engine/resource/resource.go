package resource

import (
	//"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/liuqi0826/seven/engine/display/base"
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/display/platform/opengl"
	"github.com/liuqi0826/seven/engine/display/resource"
	"github.com/liuqi0826/seven/engine/static"
	"github.com/liuqi0826/seven/events"
)

type ResourceManager struct {
	events.EventDispatcher
	sync.Mutex

	context3D core.IContext
	//静态资源库
	geometryResource  map[string]*resource.GeometryResource
	materialResource  map[string]*resource.MaterialResource
	animationResource map[string]*resource.AnimationResource
	shaderResource    map[string]*resource.ShaderResource

	//运行时资源
	geometryRuntime map[string]*base.SubGeometry
	materialRuntime map[string]*base.Material
	shaderRuntime   map[string]platform.IProgram3D
}

func (this *ResourceManager) Setup(context core.IContext) error {
	var err error

	this.context3D = context
	if this.context3D == nil {
		err = errors.New("Context3D is nil.")
	}

	this.geometryResource = make(map[string]*resource.GeometryResource)
	this.materialResource = make(map[string]*resource.MaterialResource)
	this.animationResource = make(map[string]*resource.AnimationResource)
	this.shaderResource = make(map[string]*resource.ShaderResource)

	this.geometryRuntime = make(map[string]*base.SubGeometry)
	this.materialRuntime = make(map[string]*base.Material)
	this.shaderRuntime = make(map[string]platform.IProgram3D)

	this.pretreatment()

	return err
}
func (this *ResourceManager) LoadGeometrie(file string) error {
	var err error
	res, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	} else {
		id := this.ParserGeometrie(res)
		if id != "" {
			fmt.Println("Load Geometrie " + id)
		}
	}
	return err
}
func (this *ResourceManager) LoadMaterial(file string) error {
	var err error
	res, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	} else {
		id := this.ParserMaterial(res)
		if id != "" {
			fmt.Println("Load Material " + id)
		}
	}
	return err
}
func (this *ResourceManager) LoadAnimation(file string) error {
	var err error
	res, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	} else {
		id := this.ParserAnimation(res)
		if id != "" {
			fmt.Println("Load Animation " + id)
		}
	}
	return err
}
func (this *ResourceManager) LoadShader(file string) error {
	var err error
	res, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	} else {
		id := this.ParserShader(res)
		if id != "" {
			fmt.Println("Load Shader " + id)
		}
	}
	return err
}
func (this *ResourceManager) ParserGeometrie(value []byte) string {
	gr := new(resource.GeometryResource)
	err := gr.Parser(value)
	if err == nil {
		this.AddGeometrie(gr)
		return gr.ID
	} else {
		fmt.Println(err)
	}
	return ""
}
func (this *ResourceManager) ParserMaterial(value []byte) string {
	mt := new(resource.MaterialResource)
	err := mt.Parser(value)
	if err == nil {
		this.AddMaterial(mt)
		return mt.ID
	} else {
		fmt.Println(err)
	}
	return ""
}
func (this *ResourceManager) ParserAnimation(value []byte) string {
	am := new(resource.AnimationResource)
	err := am.Parser(value)
	if err == nil {
		this.AddAnimation(am)
		return am.ID
	} else {
		fmt.Println(err)
	}
	return ""
}
func (this *ResourceManager) ParserShader(value []byte) string {
	sr := new(resource.ShaderResource)
	err := sr.Parser(value)
	if err == nil {
		this.AddShader(sr)
		return sr.ID
	} else {
		fmt.Println(err)
	}
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
	fmt.Println("AddMaterial", value.ID)
	this.materialResource[value.ID] = value
}
func (this *ResourceManager) GetMaterial(id string) *resource.MaterialResource {
	if mt, ok := this.materialResource[id]; ok {
		return mt
	}
	return nil
}
func (this *ResourceManager) AddAnimation(value *resource.AnimationResource) {
}
func (this *ResourceManager) GetAnimation(id string) *resource.AnimationResource {
	return nil
}
func (this *ResourceManager) AddShader(value *resource.ShaderResource) {
	this.shaderResource[value.ID] = value
}
func (this *ResourceManager) GetShader(id string) *resource.ShaderResource {
	if resource, ok := this.shaderResource[id]; ok {
		return resource
	}
	return nil
}

func (this *ResourceManager) CreateSubgeometrie(id string) *base.SubGeometry {
	if geometry, ok := this.geometryRuntime[id]; ok {
		geometry.UsedCount++
		return geometry
	} else {
		resource := this.GetGeometrie(id)
		if resource != nil {
			subGeometry := new(base.SubGeometry)
			this.geometryRuntime[id] = subGeometry

			subGeometry.SubGeometry(resource)
			subGeometry.UsedCount++

			event := new(events.Event)
			event.Type = static.RESOURCE_EVENT
			event.Data = func() {
				subGeometry.Upload(this.context3D)
			}
			this.DispatchEvent(event)

			return subGeometry
		}
	}
	return nil
}
func (this *ResourceManager) CreateMaterial(id string) *base.Material {
	if material, ok := this.materialRuntime[id]; ok {
		return material
	} else {
		fmt.Println(id)
		resource := this.GetMaterial(id)
		if resource != nil {
			material := new(base.Material)
			this.materialRuntime[id] = material

			material.Material(resource)
			material.AddCount()

			event := new(events.Event)
			event.Type = static.RESOURCE_EVENT
			event.Data = func() {
				material.Upload()
			}
			this.DispatchEvent(event)

			return material
		} else {
			fmt.Println(resource)
		}
	}
	return nil
}
func (this *ResourceManager) CreateShaderProgram(id string) platform.IProgram3D {
	if programe, ok := this.shaderRuntime[id]; ok {
		programe.AddCount()
		return programe
	} else {
		resource := this.GetShader(id)
		if resource != nil {
			switch static.API {
			case static.GL:
				shader := this.context3D.CreateProgram()
				this.shaderRuntime[id] = shader
				shader.AddCount()

				event := new(events.Event)
				event.Type = static.RESOURCE_EVENT
				event.Data = func() {
					shader.Upload(resource.Vertex, resource.Fragment)
				}
				this.DispatchEvent(event)

				return shader
			case static.VULKAN:
			case static.D3D9:
			case static.D3D12:
			}
		} else {
			fmt.Println("no shader res.")
		}
	}
	return nil
}
func (this *ResourceManager) pretreatment() {
	switch static.API {
	case static.GL:
		for _, v := range opengl.ShaderResource {
			this.AddShader(v)
		}
	case static.VULKAN:
	case static.D3D9:
	case static.D3D12:
	}
}
