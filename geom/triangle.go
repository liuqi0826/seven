package geom

import ()

type Triangle struct {
	V0 Vector4
	V1 Vector4
	V2 Vector4
}

func (this *Triangle) Triangle(v0, v1, v2 *Vector4) {
	this.V0 = v0.Clone()
	this.V1 = v1.Clone()
	this.V2 = v2.Clone()
}
func (this *Triangle) GetNormal() *Vector4 {
	e1, e2 := this.V1.Clone(), this.V2.Clone()
	e1.Subtract(&this.V0)
	e2.Subtract(&this.V1)
	normal := e1.CrossProduct(&e2)
	normal.Normalize()
	return normal
}
