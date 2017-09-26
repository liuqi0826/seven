package platform

import ()

type IIndexBuffer interface {
	Upload(data []uint16) error
	Dispose()
}
type IVertexBuffer interface {
	Upload(data []float32) error
	Dispose()
}
type IProgram3D interface {
	AddCount()
	SubCount()
	GetCount() uint32
	Upload(vertexProgram string, fragmentProgram string) error
	Dispose()
}
type ITexture interface {
}
