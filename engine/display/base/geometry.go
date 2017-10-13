package base

import (
	"fmt"

	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/display/resource"
	"github.com/liuqi0826/seven/geom"
)

type SubGeometry struct {
	ID        string
	UsedCount int32
	Uploaded  bool

	IndexBuffer  platform.IIndexBuffer
	VertexBuffer [8]platform.IVertexBuffer
	ValueBuffer  []float32

	projection *geom.Matrix4x4
	transform  *geom.Matrix4x4

	geometryResource *resource.GeometryResource
}

func (this *SubGeometry) SubGeometry(geometryResource *resource.GeometryResource) {
	if geometryResource != nil {
		this.geometryResource = geometryResource
	} else {
		fmt.Println("GeometryResource is nil.")
	}
}
func (this *SubGeometry) Update(transform *geom.Matrix4x4) {
	this.transform = transform
}
func (this *SubGeometry) Upload(context core.IContext) error {
	var err error
	this.IndexBuffer = context.CreateIndexBuffer()
	if this.IndexBuffer != nil {
		err = this.IndexBuffer.Upload(this.geometryResource.Geometrie.Index)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("IndexBuffer is nil")
	}
	for i, v := range this.geometryResource.Geometrie.Vertex {
		if v != nil {
			this.VertexBuffer[i] = context.CreateVertexBuffer()
			if this.VertexBuffer[i] != nil {
				err = this.VertexBuffer[i].Upload(v)
				if err != nil {
					return err
				}
			}
		}
	}
	this.Uploaded = true
	return err
}
func (this *SubGeometry) SetProjection(projection *geom.Matrix4x4) {
	this.projection = projection
}
func (this *SubGeometry) GetIndexBuffer() platform.IIndexBuffer {
	return this.IndexBuffer
}
func (this *SubGeometry) GetVertexBuffer() *[8]platform.IVertexBuffer {
	return &this.VertexBuffer
}
func (this *SubGeometry) GetValueBuffer() []float32 {
	this.ValueBuffer = make([]float32, 0)
	if this.projection != nil {
		raw := this.projection.GetRaw()
		for _, v := range raw {
			this.ValueBuffer = append(this.ValueBuffer, v)
		}
	}
	if this.transform != nil {
		raw := this.transform.GetRaw()
		for _, v := range raw {
			this.ValueBuffer = append(this.ValueBuffer, v)
		}
	}
	return this.ValueBuffer
}
func (this *SubGeometry) Dispose() {
	this.Uploaded = false
}
