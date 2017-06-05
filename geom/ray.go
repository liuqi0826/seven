package geom

import (
	"math"
)

type Ray struct {
	Origin    Vector4
	Direction Vector4
}

func (this *Ray) Ray(origin *Vector4, direction *Vector4) {
	this.Origin = origin.Clone()
	this.Direction = direction.Clone()
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
func (this *Ray) IntersectTriangle(triangle *Triangle) bool {
	e1, e2 := triangle.V1.Clone(), triangle.V2.Clone()
	e1.Subtract(&triangle.V0)
	e2.Subtract(&triangle.V0)
	p := this.Direction.CrossProduct(&e2)
	det := e1.DotProduct(p)
	var t Vector4
	if det > 0 {
		t = this.Origin.Clone()
		t.Subtract(&triangle.V0)
	} else {
		t = triangle.V0.Clone()
		t.Subtract(&this.Origin)
		det = -det
	}
	if det < 0.00000001 {
		return false
	}
	tu := t.DotProduct(p)
	if tu < 0 || tu > det {
		return false
	}
	q := t.Clone()
	q.CrossProduct(&e1)
	tv := this.Direction.DotProduct(&q)
	if tu < 0 || tv+tu > det {
		return false
	}
	return true
}
func (this *Ray) IntersectTriangleWithHitPoint(triangle *Triangle) *Vector4 {
	e1, e2 := triangle.V1.Clone(), triangle.V2.Clone()
	e1.Subtract(&triangle.V0)
	e2.Subtract(&triangle.V0)
	p := this.Direction.CrossProduct(&e2)
	det := e1.DotProduct(p)
	var t Vector4
	if det > 0 {
		t = this.Origin.Clone()
		t.Subtract(&triangle.V0)
	} else {
		t = triangle.V0.Clone()
		t.Subtract(&this.Origin)
		det = -det
	}
	if det < 0.00000001 {
		return nil
	}
	tu := t.DotProduct(p)
	if tu < 0 || tu > det {
		return nil
	}
	q := t.Clone()
	q.CrossProduct(&e1)
	tv := this.Direction.DotProduct(&q)
	if tu < 0 || tv+tu > det {
		return nil
	}
	rt := e2.DotProduct(&q)
	invDet := 1 / det
	rt = rt * invDet
	tu = tu * invDet
	tv = tv * invDet
	point := this.Direction.Clone()
	point.Mul(rt)
	point.Add(&this.Origin)
	return &point
}
func (this *Ray) IntersectAABB(aabb *Box) bool {
	testPlaneList := make([]*Plane, 0)
	if this.Origin.X < aabb.Max.X && this.Origin.Y < aabb.Max.Y && this.Origin.Z < aabb.Max.Z && this.Origin.X > aabb.Mix.X && this.Origin.Y > aabb.Mix.Y && this.Origin.Z > aabb.Mix.Z {
		return true
	}
	if this.Origin.X > aabb.Max.X {
		p0 := &Vector4{X: aabb.Max.X, Y: aabb.Max.Y, Z: aabb.Max.Z}
		p1 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
		p2 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Mix.Z}
		p := new(Plane)
		p.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, p)
	}
	if this.Origin.Y > aabb.Max.Y {
		p0 := &Vector4{X: aabb.Mix.X, Y: aabb.Max.Y, Z: aabb.Mix.Z}
		p1 := &Vector4{X: aabb.Mix.X, Y: aabb.Max.Y, Z: aabb.Max.Z}
		p2 := &Vector4{X: aabb.Max.X, Y: aabb.Max.Y, Z: aabb.Max.Z}
		p := new(Plane)
		p.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, p)
	}
	if this.Origin.Z > aabb.Max.Z {
		p0 := &Vector4{X: aabb.Mix.X, Y: aabb.Max.Y, Z: aabb.Max.Z}
		p1 := &Vector4{X: aabb.Mix.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
		p2 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
		p := new(Plane)
		p.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, p)
	}
	if this.Origin.X < aabb.Max.X {
		p0 := &Vector4{X: aabb.Mix.X, Y: aabb.Max.Y, Z: aabb.Mix.Z}
		p1 := &Vector4{X: aabb.Mix.X, Y: aabb.Mix.Y, Z: aabb.Mix.Z}
		p2 := &Vector4{X: aabb.Mix.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
		p := new(Plane)
		p.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, p)
	}
	if this.Origin.Y < aabb.Max.Y {
		p0 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Mix.Z}
		p1 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
		p2 := &Vector4{X: aabb.Mix.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
		p := new(Plane)
		p.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, p)
	}
	if this.Origin.Z < aabb.Max.Z {
		p0 := &Vector4{X: aabb.Max.X, Y: aabb.Max.Y, Z: aabb.Mix.Z}
		p1 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Mix.Z}
		p2 := &Vector4{X: aabb.Mix.X, Y: aabb.Mix.Y, Z: aabb.Mix.Z}
		p := new(Plane)
		p.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, p)
	}
	for _, p := range testPlaneList {
		tp := this.IntersectPlaneHitPoint(p)
		if tp != nil {
			if tp.X-aabb.Max.X >= 0 && tp.X-aabb.Max.X < 0.00000001 {
				if tp.Y >= aabb.Mix.X && tp.Y <= aabb.Max.Y && tp.Z >= aabb.Mix.Z && tp.Z <= aabb.Max.Z {
					return true
				}
			}
			if tp.Y-aabb.Max.Y >= 0 && tp.Y-aabb.Max.Y < 0.00000001 {
				if tp.X >= aabb.Mix.X && tp.X <= aabb.Max.X && tp.Z >= aabb.Mix.Z && tp.Z <= aabb.Max.Z {
					return true
				}
			}
			if tp.Z-aabb.Max.Z >= 0 && tp.Z-aabb.Max.Z < 0.00000001 {
				if tp.X >= aabb.Mix.X && tp.X <= aabb.Max.X && tp.Y >= aabb.Mix.Y && tp.Y <= aabb.Max.Y {
					return true
				}
			}
			if tp.X-aabb.Mix.X >= 0 && tp.X-aabb.Mix.X < 0.00000001 {
				if tp.Y >= aabb.Mix.X && tp.Y <= aabb.Max.Y && tp.Z >= aabb.Mix.Z && tp.Z <= aabb.Max.Z {
					return true
				}
			}
			if tp.Y-aabb.Mix.Y >= 0 && tp.Y-aabb.Mix.Y < 0.00000001 {
				if tp.X >= aabb.Mix.X && tp.X <= aabb.Max.X && tp.Z >= aabb.Mix.Z && tp.Z <= aabb.Max.Z {
					return true
				}
			}
			if tp.Z-aabb.Mix.Z >= 0 && tp.Z-aabb.Mix.Z < 0.00000001 {
				if tp.X >= aabb.Mix.X && tp.X <= aabb.Max.X && tp.Y >= aabb.Mix.Y && tp.Y <= aabb.Max.Y {
					return true
				}
			}
		}
	}
	return false
}
func (this *Ray) IntersectAABBWithHitPoint(aabb *Box) *Vector4 {
	testPlaneList := make([]*Plane, 0)
	if this.Origin.X < aabb.Max.X && this.Origin.Y < aabb.Max.Y && this.Origin.Z < aabb.Max.Z && this.Origin.X > aabb.Mix.X && this.Origin.Y > aabb.Mix.Y && this.Origin.Z > aabb.Mix.Z {
		var p0, p1, p2 *Vector4
		p0.X, p0.Y, p0.Z = aabb.Max.X, aabb.Max.Y, aabb.Max.Z
		p1.X, p1.Y, p1.Z = aabb.Max.X, aabb.Mix.Y, aabb.Max.Z
		p2.X, p2.Y, p2.Z = aabb.Max.X, aabb.Mix.Y, aabb.Mix.Z
		pl0 := new(Plane)
		pl0.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, pl0)
		p0.X, p0.Y, p0.Z = aabb.Mix.X, aabb.Max.Y, aabb.Mix.Z
		p1.X, p1.Y, p1.Z = aabb.Mix.X, aabb.Max.Y, aabb.Max.Z
		p2.X, p2.Y, p2.Z = aabb.Max.X, aabb.Max.Y, aabb.Max.Z
		pl1 := new(Plane)
		pl1.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, pl1)
		p0.X, p0.Y, p0.Z = aabb.Mix.X, aabb.Max.Y, aabb.Max.Z
		p1.X, p1.Y, p1.Z = aabb.Mix.X, aabb.Mix.Y, aabb.Max.Z
		p2.X, p2.Y, p2.Z = aabb.Max.X, aabb.Mix.Y, aabb.Max.Z
		pl2 := new(Plane)
		pl2.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, pl2)
		p0.X, p0.Y, p0.Z = aabb.Mix.X, aabb.Max.Y, aabb.Mix.Z
		p1.X, p1.Y, p1.Z = aabb.Mix.X, aabb.Mix.Y, aabb.Mix.Z
		p2.X, p2.Y, p2.Z = aabb.Mix.X, aabb.Mix.Y, aabb.Max.Z
		pl3 := new(Plane)
		pl3.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, pl3)
		p0.X, p0.Y, p0.Z = aabb.Max.X, aabb.Mix.Y, aabb.Mix.Z
		p1.X, p1.Y, p1.Z = aabb.Max.X, aabb.Mix.Y, aabb.Max.Z
		p2.X, p2.Y, p2.Z = aabb.Mix.X, aabb.Mix.Y, aabb.Max.Z
		pl4 := new(Plane)
		pl4.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, pl4)
		p0.X, p0.Y, p0.Z = aabb.Max.X, aabb.Max.Y, aabb.Mix.Z
		p1.X, p1.Y, p1.Z = aabb.Max.X, aabb.Mix.Y, aabb.Mix.Z
		p2.X, p2.Y, p2.Z = aabb.Mix.X, aabb.Mix.Y, aabb.Mix.Z
		pl5 := new(Plane)
		pl5.FromPoints(p0, p1, p2)
		testPlaneList = append(testPlaneList, pl5)
	} else {
		if this.Origin.X > aabb.Max.X {
			p0 := &Vector4{X: aabb.Max.X, Y: aabb.Max.Y, Z: aabb.Max.Z}
			p1 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
			p2 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Mix.Z}
			p := new(Plane)
			p.FromPoints(p0, p1, p2)
			testPlaneList = append(testPlaneList, p)
		}
		if this.Origin.Y > aabb.Max.Y {
			p0 := &Vector4{X: aabb.Mix.X, Y: aabb.Max.Y, Z: aabb.Mix.Z}
			p1 := &Vector4{X: aabb.Mix.X, Y: aabb.Max.Y, Z: aabb.Max.Z}
			p2 := &Vector4{X: aabb.Max.X, Y: aabb.Max.Y, Z: aabb.Max.Z}
			p := new(Plane)
			p.FromPoints(p0, p1, p2)
			testPlaneList = append(testPlaneList, p)
		}
		if this.Origin.Z > aabb.Max.Z {
			p0 := &Vector4{X: aabb.Mix.X, Y: aabb.Max.Y, Z: aabb.Max.Z}
			p1 := &Vector4{X: aabb.Mix.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
			p2 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
			p := new(Plane)
			p.FromPoints(p0, p1, p2)
			testPlaneList = append(testPlaneList, p)
		}
		if this.Origin.X < aabb.Max.X {
			p0 := &Vector4{X: aabb.Mix.X, Y: aabb.Max.Y, Z: aabb.Mix.Z}
			p1 := &Vector4{X: aabb.Mix.X, Y: aabb.Mix.Y, Z: aabb.Mix.Z}
			p2 := &Vector4{X: aabb.Mix.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
			p := new(Plane)
			p.FromPoints(p0, p1, p2)
			testPlaneList = append(testPlaneList, p)
		}
		if this.Origin.Y < aabb.Max.Y {
			p0 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Mix.Z}
			p1 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
			p2 := &Vector4{X: aabb.Mix.X, Y: aabb.Mix.Y, Z: aabb.Max.Z}
			p := new(Plane)
			p.FromPoints(p0, p1, p2)
			testPlaneList = append(testPlaneList, p)
		}
		if this.Origin.Z < aabb.Max.Z {
			p0 := &Vector4{X: aabb.Max.X, Y: aabb.Max.Y, Z: aabb.Mix.Z}
			p1 := &Vector4{X: aabb.Max.X, Y: aabb.Mix.Y, Z: aabb.Mix.Z}
			p2 := &Vector4{X: aabb.Mix.X, Y: aabb.Mix.Y, Z: aabb.Mix.Z}
			p := new(Plane)
			p.FromPoints(p0, p1, p2)
			testPlaneList = append(testPlaneList, p)
		}
	}
	hitPointList := make([]*Vector4, 0)
	for _, p := range testPlaneList {
		tp := this.IntersectPlaneHitPoint(p)
		if tp != nil {
			if tp.X-aabb.Max.X >= 0 && tp.X-aabb.Max.X < 0.00000001 {
				if tp.Y >= aabb.Mix.X && tp.Y <= aabb.Max.Y && tp.Z >= aabb.Mix.Z && tp.Z <= aabb.Max.Z {
					hitPointList = append(hitPointList, tp)
				}
			}
			if tp.Y-aabb.Max.Y >= 0 && tp.Y-aabb.Max.Y < 0.00000001 {
				if tp.X >= aabb.Mix.X && tp.X <= aabb.Max.X && tp.Z >= aabb.Mix.Z && tp.Z <= aabb.Max.Z {
					hitPointList = append(hitPointList, tp)
				}
			}
			if tp.Z-aabb.Max.Z >= 0 && tp.Z-aabb.Max.Z < 0.00000001 {
				if tp.X >= aabb.Mix.X && tp.X <= aabb.Max.X && tp.Y >= aabb.Mix.Y && tp.Y <= aabb.Max.Y {
					hitPointList = append(hitPointList, tp)
				}
			}
			if tp.X-aabb.Mix.X >= 0 && tp.X-aabb.Mix.X < 0.00000001 {
				if tp.Y >= aabb.Mix.X && tp.Y <= aabb.Max.Y && tp.Z >= aabb.Mix.Z && tp.Z <= aabb.Max.Z {
					hitPointList = append(hitPointList, tp)
				}
			}
			if tp.Y-aabb.Mix.Y >= 0 && tp.Y-aabb.Mix.Y < 0.00000001 {
				if tp.X >= aabb.Mix.X && tp.X <= aabb.Max.X && tp.Z >= aabb.Mix.Z && tp.Z <= aabb.Max.Z {
					hitPointList = append(hitPointList, tp)
				}
			}
			if tp.Z-aabb.Mix.Z >= 0 && tp.Z-aabb.Mix.Z < 0.00000001 {
				if tp.X >= aabb.Mix.X && tp.X <= aabb.Max.X && tp.Y >= aabb.Mix.Y && tp.Y <= aabb.Max.Y {
					hitPointList = append(hitPointList, tp)
				}
			}
		}
	}
	if len(hitPointList) > 0 {
		hit := &Vector4{X: 999999999, Y: 999999999, Z: 999999999}
		for _, p := range hitPointList {
			if hit.LengthSquared() < p.LengthSquared() {
				hit = p
			}
		}
		return hit
	}
	return nil
}
