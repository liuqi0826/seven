package base

import (
	"fmt"

	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/display/resource"
)

type SubGeometry struct {
	ID string

	IndexBuffer  platform.IIndexBuffer
	VertexBuffer [8]platform.IVertexBuffer

	usedCount int32
	uploaded  bool

	geometryResource *resource.GeometryResource
}

func (this *SubGeometry) SubGeometry(geometryResource *resource.GeometryResource) {
	if geometryResource != nil {
		this.geometryResource = geometryResource
	} else {
		fmt.Println("GeometryResource is nil.")
	}
}
func (this *SubGeometry) Upload(context core.IContext) error {
	var err error
	this.IndexBuffer = context.CreateIndexBuffer()
	//fmt.Println("upload", this.IndexBuffer)
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
	this.uploaded = true
	return err
}
func (this *SubGeometry) GetIndexBuffer() platform.IIndexBuffer {
	return this.IndexBuffer
}
func (this *SubGeometry) GetVertexBuffer() *[8]platform.IVertexBuffer {
	return &this.VertexBuffer
}
func (this *SubGeometry) AddCount() {
	this.usedCount++
}
func (this *SubGeometry) SubCount() {
	if this.usedCount > 0 {
		this.usedCount--
	}
}
func (this *SubGeometry) GetCount() int32 {
	return this.usedCount
}
func (this *SubGeometry) IsReady() bool {
	return this.uploaded
}
func (this *SubGeometry) Dispose() {
	this.uploaded = false
}
