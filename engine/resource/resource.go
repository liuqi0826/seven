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
	skeletonResource  map[string]*resource.SkeletonResource
	materialResource  map[string]*resource.MaterialResource
	animationResource map[string]*resource.AnimationClipResource
	shaderResource    map[string]*resource.ShaderResource

	//运行时资源
	geometryRuntime map[string]*base.SubGeometry
	materialRuntime map[string]*base.Material
	shaderRuntime   map[string]platform.IProgram3D

	shaderUploadQueue    []*ShaderUploader
	geometrieUploadQueue []*GeometrieUploader
	materialUploadQueue  []*MaterialUploader
}

func (this *ResourceManager) Setup(context core.IContext) error {
	var err error

	this.context3D = context
	if this.context3D == nil {
		err = errors.New("Context3D is nil.")
	}

	this.geometryResource = make(map[string]*resource.GeometryResource)
	this.skeletonResource = make(map[string]*resource.SkeletonResource)
	this.materialResource = make(map[string]*resource.MaterialResource)
	this.animationResource = make(map[string]*resource.AnimationClipResource)
	this.shaderResource = make(map[string]*resource.ShaderResource)

	this.geometryRuntime = make(map[string]*base.SubGeometry)
	this.materialRuntime = make(map[string]*base.Material)
	this.shaderRuntime = make(map[string]platform.IProgram3D)

	this.shaderUploadQueue = make([]*ShaderUploader, 0)
	this.geometrieUploadQueue = make([]*GeometrieUploader, 0)
	this.materialUploadQueue = make([]*MaterialUploader, 0)

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
func (this *ResourceManager) LoadSkeleton(file string) error {
	var err error
	res, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	} else {
		id := this.ParserSkeleton(res)
		if id != "" {
			fmt.Println("Load Skeleton " + id)
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
func (this *ResourceManager) LoadAnimationClip(file string) error {
	var err error
	res, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	} else {
		id := this.ParserAnimationClip(res)
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
func (this *ResourceManager) ParserSkeleton(value []byte) string {
	sk := new(resource.SkeletonResource)
	err := sk.Parser(value)
	if err == nil {
		this.AddSkeleton(sk)
		return sk.ID
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
func (this *ResourceManager) ParserAnimationClip(value []byte) string {
	am := new(resource.AnimationClipResource)
	err := am.Parser(value)
	if err == nil {
		this.AddAnimationClip(am)
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
func (this *ResourceManager) AddSkeleton(value *resource.SkeletonResource) {
	this.skeletonResource[value.ID] = value
}
func (this *ResourceManager) GetSkeleton(id string) *resource.SkeletonResource {
	if sk, ok := this.skeletonResource[id]; ok {
		return sk
	}
	return nil
}
func (this *ResourceManager) AddMaterial(value *resource.MaterialResource) {
	this.materialResource[value.ID] = value
}
func (this *ResourceManager) GetMaterial(id string) *resource.MaterialResource {
	if mt, ok := this.materialResource[id]; ok {
		return mt
	}
	return nil
}
func (this *ResourceManager) AddAnimationClip(value *resource.AnimationClipResource) {
	this.animationResource[value.ID] = value
}
func (this *ResourceManager) GetAnimationClip(id string) *resource.AnimationClipResource {
	if an, ok := this.animationResource[id]; ok {
		return an
	}
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
		geometry.AddCount()
		return geometry
	} else {
		resource := this.GetGeometrie(id)
		if resource != nil {
			subGeometry := new(base.SubGeometry)
			subGeometry.AddCount()
			this.geometryRuntime[id] = subGeometry

			this.Lock()
			gu := new(GeometrieUploader)
			gu.Target = subGeometry
			gu.Resource = resource
			this.geometrieUploadQueue = append(this.geometrieUploadQueue, gu)
			this.Unlock()

			evt := new(events.Event)
			evt.Type = static.RESOURCE_EVENT
			this.DispatchEvent(evt)

			return subGeometry
		}
	}
	return nil
}
func (this *ResourceManager) CreateMaterial(id string) *base.Material {
	if material, ok := this.materialRuntime[id]; ok {
		return material
	} else {
		resource := this.GetMaterial(id)
		if resource != nil {
			material := new(base.Material)
			material.AddCount()
			this.materialRuntime[id] = material

			this.Lock()
			mu := new(MaterialUploader)
			mu.Target = material
			mu.Resource = resource
			this.materialUploadQueue = append(this.materialUploadQueue, mu)
			this.Unlock()

			evt := new(events.Event)
			evt.Type = static.RESOURCE_EVENT
			this.DispatchEvent(evt)

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
func (this *ResourceManager) Upload() {
	this.Lock()
	defer this.Unlock()

	var su, gu, mu bool
	for _, s := range this.shaderUploadQueue {
		s.Target = this.context3D.CreateProgram()
		s.Target.Upload(s.Resource.Vertex, s.Resource.Fragment)
		su = true
	}
	for _, g := range this.geometrieUploadQueue {
		g.Target.SubGeometry(g.Resource)
		g.Target.Upload(this.context3D)
		gu = true
	}
	for _, m := range this.materialUploadQueue {
		m.Target.Material(m.Resource)
		m.Target.Upload()
		mu = true
	}
	if su {
		this.shaderUploadQueue = make([]*ShaderUploader, 0)
	}
	if gu {
		this.geometrieUploadQueue = make([]*GeometrieUploader, 0)
	}
	if mu {
		this.materialUploadQueue = make([]*MaterialUploader, 0)
	}
}
func (this *ResourceManager) pretreatment() {
	switch static.API {
	case static.GL:
		for _, v := range opengl.ShaderResource {
			this.AddShader(v)

			//现阶段所有shader在开始时统一创建上传
			shader := this.context3D.CreateProgram()
			shader.Upload(v.Vertex, v.Fragment)
			this.shaderRuntime[v.ID] = shader
		}
	case static.VULKAN:
	case static.D3D9:
	case static.D3D12:
	}
}

type ShaderUploader struct {
	Target   platform.IProgram3D
	Resource *resource.ShaderResource
}
type GeometrieUploader struct {
	Target   *base.SubGeometry
	Resource *resource.GeometryResource
}
type MaterialUploader struct {
	Target   *base.Material
	Resource *resource.MaterialResource
}
