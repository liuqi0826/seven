package base

import (
	"fmt"

	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/display/resource"
)

type SubGeometry struct {
	ID        string
	UsedCount int32
	Uploaded  bool

	IndexBuffer  platform.IIndexBuffer
	VertexBuffer [8]platform.IVertexBuffer

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
func (this *SubGeometry) GetIndexBuffer() platform.IIndexBuffer {
	return this.IndexBuffer
}
func (this *SubGeometry) GetVertexBuffer() *[8]platform.IVertexBuffer {
	return &this.VertexBuffer
}
func (this *SubGeometry) Dispose() {

	this.Uploaded = false
}
