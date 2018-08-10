package platform

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
	IsReady() bool
	Dispose()
}
type ITexture interface {
	Upload(source []byte, sourceType string) error
	UploadCompressedTexture(source []byte) error
	SetSlot(index int32)
	Dispose()
}
