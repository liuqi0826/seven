package resource

import (
	"github.com/liuqi0826/seven/geom"
)

type Geometry struct {
	ID       string
	Shader   string
	Material string

	Bound geom.Box

	VertexIndex []uint16
	ColorIndex  []uint16
	Vertex      [8][]float32
}

func (this *Geometry) Geometry() {

}
