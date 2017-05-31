package geom

import ()

type Sphere struct {
	Center Vector4
	Radius float32
}

func (this *Sphere) Sphere(center Vector4, radius float32) {
	this.Center = center
	this.Radius = radius
}
func (this *Sphere) Clone() Sphere {
	var s Sphere
	s.Sphere(this.Center.Clone(), this.Radius)
	return s
}
