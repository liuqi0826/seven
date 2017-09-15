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
	if this.geometryResource.Geometrie.Vertex.Slot0 != nil {
		this.VertexBuffer[0] = context.CreateVertexBuffer()
		if this.VertexBuffer[0] != nil {
			err = this.VertexBuffer[0].Upload(this.geometryResource.Geometrie.Vertex.Slot0)
			if err != nil {
				return err
			}
		}
	}
	if this.geometryResource.Geometrie.Vertex.Slot1 != nil {
		this.VertexBuffer[1] = context.CreateVertexBuffer()
		if this.VertexBuffer[1] != nil {
			err = this.VertexBuffer[1].Upload(this.geometryResource.Geometrie.Vertex.Slot1)
			if err != nil {
				return err
			}
		}
	}
	if this.geometryResource.Geometrie.Vertex.Slot2 != nil {
		this.VertexBuffer[2] = context.CreateVertexBuffer()
		if this.VertexBuffer[2] != nil {
			err = this.VertexBuffer[2].Upload(this.geometryResource.Geometrie.Vertex.Slot2)
			if err != nil {
				return err
			}
		}
	}
	if this.geometryResource.Geometrie.Vertex.Slot3 != nil {
		this.VertexBuffer[3] = context.CreateVertexBuffer()
		if this.VertexBuffer[3] != nil {
			err = this.VertexBuffer[3].Upload(this.geometryResource.Geometrie.Vertex.Slot3)
			if err != nil {
				return err
			}
		}
	}
	if this.geometryResource.Geometrie.Vertex.Slot4 != nil {
		this.VertexBuffer[4] = context.CreateVertexBuffer()
		if this.VertexBuffer[4] != nil {
			err = this.VertexBuffer[4].Upload(this.geometryResource.Geometrie.Vertex.Slot4)
			if err != nil {
				return err
			}
		}
	}
	if this.geometryResource.Geometrie.Vertex.Slot5 != nil {
		this.VertexBuffer[5] = context.CreateVertexBuffer()
		if this.VertexBuffer[5] != nil {
			err = this.VertexBuffer[5].Upload(this.geometryResource.Geometrie.Vertex.Slot5)
			if err != nil {
				return err
			}
		}
	}
	if this.geometryResource.Geometrie.Vertex.Slot6 != nil {
		this.VertexBuffer[6] = context.CreateVertexBuffer()
		if this.VertexBuffer[6] != nil {
			err = this.VertexBuffer[6].Upload(this.geometryResource.Geometrie.Vertex.Slot6)
			if err != nil {
				return err
			}
		}
	}
	if this.geometryResource.Geometrie.Vertex.Slot7 != nil {
		this.VertexBuffer[7] = context.CreateVertexBuffer()
		if this.VertexBuffer[7] != nil {
			err = this.VertexBuffer[7].Upload(this.geometryResource.Geometrie.Vertex.Slot7)
			if err != nil {
				return err
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
