package object

import (
	"github.com/liuqi0826/seven/engine/display/base"
)

type Mesh struct {
	geometry []*base.SubGeometry
}

func (this *Mesh) Mesh() {
	this.geometry = make([]*base.SubGeometry, 0)
}
