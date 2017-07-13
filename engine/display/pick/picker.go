package pick

import (
	"github.com/liuqi0826/seven/geom"
)

type Picker struct {
	ray geom.Ray
}

func (this *Picker) Picker() {
	origin := new(geom.Vector4)
	origin.Vector4()
	direction := new(geom.Vector4)
	direction.Vector4()
	direction.Z = 1.0
	this.ray = geom.Ray{}
	this.ray.Ray(origin, direction)
}
