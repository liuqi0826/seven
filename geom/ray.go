package geom

import ()

type Ray struct {
	Origin    *Vector4
	Direction *Vector4
}

func (this *Ray) Ray(origin *Vector4, direction *Vector4) {
	this.Origin = origin
	this.Direction = direction
	if this.Origin == nil {
		this.Origin = new(Vector4)
		this.Origin.Vector4()
	}
	if this.Direction == nil {
		this.Direction = new(Vector4)
		this.Direction.Vector4()
	}
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
	p.Add(this.Origin)
	return &p
}
func (this *Ray) IntersectPlane(plane *Plane) bool {
	var n Vector4 = plane.GetNormal()
	var p Vector4 = plane.GetRandPoint()
	var t float32 = (n.DotProduct(&p) - n.DotProduct(this.Origin)) / n.DotProduct(this.Direction)
	if t >= 0.0 {
		return true
	}
	return false
}
func (this *Ray) IntersectPlaneHitPoint(plane *Plane) *Vector4 {
	var n Vector4 = plane.GetNormal()
	var p Vector4 = plane.GetRandPoint()
	var t float32 = (n.DotProduct(&p) - n.DotProduct(this.Origin)) / n.DotProduct(this.Direction)
	if t >= 0.0 {
		var point Vector4 = this.Direction.Clone()
		point.Scale(t)
		point.Add(this.Origin)
		return &point
	}
	return nil
}
