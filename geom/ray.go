package geom

import (
	"math"
)

type Ray struct {
	Origin    Vector4
	Direction Vector4
}

func (this *Ray) Ray(origin Vector4, direction Vector4) {
	this.Origin = origin
	this.Direction = direction
}
func (this *Ray) ClosestPoint(point *Vector4) *Vector4 {
	var p Vector4
	var v Vector4 = Vector4{X: point.X - this.Origin.X, Y: point.Y - this.Origin.Y, Z: point.Z - this.Origin.Z}
	var t float32 = this.Direction.DotProduct(&v)
	if t < 0 {
		p = this.Origin.Clone()
		return &p
	}
	if t > 1 {
		p = this.Direction.Clone()
		return &p
	}
	p = this.Direction.Clone()
	p.Mul(t)
	p.Add(&this.Origin)
	return &p
}
func (this *Ray) IntersectPlane(plane *Plane) bool {
	var n Vector4 = plane.GetNormal()
	var p Vector4 = plane.GetRandPoint()
	var t float32 = (n.DotProduct(&p) - n.DotProduct(&this.Origin)) / n.DotProduct(&this.Direction)
	if t >= 0.0 {
		return true
	}
	return false
}
func (this *Ray) IntersectPlaneHitPoint(plane *Plane) *Vector4 {
	var n Vector4 = plane.GetNormal()
	var p Vector4 = plane.GetRandPoint()
	var t float32 = (n.DotProduct(&p) - n.DotProduct(&this.Origin)) / n.DotProduct(&this.Direction)
	if t >= 0.0 {
		var point Vector4 = this.Direction.Clone()
		point.Scale(t)
		point.Add(&this.Origin)
		return &point
	}
	return nil
}
func (this *Ray) IntersectSphere(sphere *Sphere) bool {
	e := Vector4{X: sphere.Center.X - this.Origin.X, Y: sphere.Center.Y - this.Origin.Y, Z: sphere.Center.Z - this.Origin.Z}
	elen := e.Length()
	a := this.Direction.DotProduct(&e)
	f := float32(math.Sqrt(float64(sphere.Radius*sphere.Radius - elen*elen + a*a)))
	if f >= 0 {
		return true
	}
	return false
}
func (this *Ray) IntersectSphereWithHitPoint(sphere *Sphere) *Vector4 {
	e := Vector4{X: sphere.Center.X - this.Origin.X, Y: sphere.Center.Y - this.Origin.Y, Z: sphere.Center.Z - this.Origin.Z}
	elen := e.Length()
	a := this.Direction.DotProduct(&e)
	f := float32(math.Sqrt(float64(sphere.Radius*sphere.Radius - elen*elen + a*a)))
	if f >= 0 {
		t := a - f
		point := this.Direction.Clone()
		point.Scale(t)
		point.Add(&this.Origin)
		return &point
	}
	return nil
}
