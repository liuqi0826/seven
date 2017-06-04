package geom

import ()

type Sphere struct {
	Center Vector4
	Radius float32
}

func (this *Sphere) Sphere(center *Vector4, radius float32) {
	this.Center = center.Clone()
	this.Radius = radius
}
func (this *Sphere) Clone() Sphere {
	c := this.Center.Clone()
	var s Sphere
	s.Sphere(&c, this.Radius)
	return s
}
