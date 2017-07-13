package base

import (
	"github.com/gonutz/d3d9"
	"github.com/liuqi0826/seven/engine/display/resource"
)

type SubGeometry struct {
	ID        string
	UsedCount int32
	Uploaded  bool

	IndexBuffer  d3d9.IndexBuffer
	VertexBuffer d3d9.VertexBuffer

	geometryResource *resource.GeometryResource
}

func (this *SubGeometry) SubGeometry(geometryResource *resource.GeometryResource) {
	this.geometryResource = geometryResource
}
func (this *SubGeometry) Upload(context d3d9.Device) {
	//this.IndexBuffer = context.CreateIndexBuffer(Length, Usage, Format, Pool, pSharedHandle)
	//this.VertexBuffer = context.CreateVertexBuffer(Length, Usage, FVF, Pool, pSharedHandle)
	this.Uploaded = true
}
func (this *SubGeometry) Dispose() {

	this.Uploaded = false
}
