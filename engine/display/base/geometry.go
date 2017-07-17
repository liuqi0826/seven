package base

import (
	"fmt"

	"github.com/gonutz/d3d9"
	"github.com/liuqi0826/seven/engine/display/resource"
)

type SubGeometry struct {
	ID        string
	UsedCount int32
	Uploaded  bool

	IndexBuffer  d3d9.IndexBuffer
	VertexBuffer [8]d3d9.VertexBuffer

	geometryResource *resource.GeometryResource
}

func (this *SubGeometry) SubGeometry(geometryResource *resource.GeometryResource) {
	if geometryResource != nil {
		this.geometryResource = geometryResource
	} else {
		fmt.Println("GeometryResource is nil.")
	}
}
func (this *SubGeometry) Upload(context d3d9.Device) error {
	var err error
	if len(this.geometryResource.Geometrie.Index) > 0 {
		this.IndexBuffer, err = context.CreateIndexBuffer(uint(len(this.geometryResource.Geometrie.Index)*2), d3d9.USAGE_WRITEONLY, d3d9.FMT_INDEX16, d3d9.POOL_MANAGED, nil)
		if err == nil {
			data, err := this.IndexBuffer.Lock(0, 0, d3d9.LOCK_DISCARD)
			if err == nil {
				data.SetUint16s(0, this.geometryResource.Geometrie.Index)
			} else {
				fmt.Println(err)
			}
			this.IndexBuffer.Unlock()
		} else {
			fmt.Println(err)
		}
	}
	if len(this.geometryResource.Geometrie.Vertex.Slot0) > 0 {
		this.VertexBuffer[0], err = context.CreateVertexBuffer(uint(len(this.geometryResource.Geometrie.Vertex.Slot0)*4), d3d9.USAGE_WRITEONLY, 0, d3d9.POOL_MANAGED, nil)
		if err == nil {
			data, err := this.VertexBuffer[0].Lock(0, 0, d3d9.LOCK_DISCARD)
			if err == nil {
				data.SetFloat32s(0, this.geometryResource.Geometrie.Vertex.Slot0)
			} else {
				fmt.Println(err)
			}
			this.VertexBuffer[0].Unlock()
		} else {
			fmt.Println(err)
		}
	}
	this.Uploaded = true
	return err
}
func (this *SubGeometry) Dispose() {

	this.Uploaded = false
}
