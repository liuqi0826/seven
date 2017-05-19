package geom

import (
	"math"
	"math/rand"
)

type Plane struct {
	A         float32
	B         float32
	C         float32
	D         float32
	alignment string
}

func (this *Plane) Plane(a float32, b float32, c float32, d float32) {
	if a == 0 && b == 0 && c == 0 {
		panic("Plane a,b,c all is 0.")
	}
	this.A, this.B, this.C, this.D = a, b, c, d
	if this.A == 0 && this.B == 0 {
		this.alignment = ALIGN_XY_AXIS
	} else if this.B == 0 && this.C == 0 {
		this.alignment = ALIGN_YZ_AXIS
	} else if this.A == 0 && this.C == 0 {
		this.alignment = ALIGN_XZ_AXIS
	} else {
		this.alignment = ALIGN_ANY
	}
}
func (this *Plane) FromPoints(p0 *Vector4, p1 *Vector4, p2 *Vector4) {
	d1x := p1.X - p0.X
	d1y := p1.Y - p0.Y
	d1z := p1.Z - p0.Z
	d2x := p2.X - p0.X
	d2y := p2.Y - p0.Y
	d2z := p2.Z - p0.Z
	this.A = d1y*d2z - d1z*d2y
	this.B = d1z*d2x - d1x*d2z
	this.C = d1x*d2y - d1y*d2x
	this.D = this.A*p0.X + this.B*p0.Y + this.C*p0.Z
	if this.A == 0 && this.B == 0 {
		this.alignment = ALIGN_XY_AXIS
	} else if this.B == 0 && this.C == 0 {
		this.alignment = ALIGN_YZ_AXIS
	} else if this.A == 0 && this.C == 0 {
		this.alignment = ALIGN_XZ_AXIS
	} else {
		this.alignment = ALIGN_ANY
	}
}
func (this *Plane) FromNormalAndPoint(normal *Vector4, point *Vector4) {
	this.A = normal.X
	this.B = normal.Y
	this.C = normal.Z
	this.D = this.A*point.X + this.B*point.Y + this.C*point.Z
	if this.A == 0 && this.B == 0 {
		this.alignment = ALIGN_XY_AXIS
	} else if this.B == 0 && this.C == 0 {
		this.alignment = ALIGN_YZ_AXIS
	} else if this.A == 0 && this.C == 0 {
		this.alignment = ALIGN_XZ_AXIS
	} else {
		this.alignment = ALIGN_ANY
	}
}
func (this *Plane) GetNormal() Vector4 {
	var normal Vector4
	normal.X = this.A
	normal.Y = this.B
	normal.Z = this.C
	return normal
}
func (this *Plane) GetRandPoint() Vector4 {
	var v Vector4
	rx := rand.Float32()
	ry := rand.Float32()
	rz := -(this.A*rx + this.B*ry + this.D) / this.C
	v.X, v.Y, v.Z, v.W = rx, ry, rz, 1.0
	return v
}
func (this *Plane) Normalize() {
	len := float32(1 / math.Sqrt(float64(this.A*this.A+this.B*this.B+this.C*this.C)))
	this.A *= len
	this.B *= len
	this.C *= len
	this.D *= len
}
func (this *Plane) Distance(v *Vector4) float32 {
	var r float32
	switch this.alignment {
	case ALIGN_YZ_AXIS:
		r = this.A*v.X - this.D
	case ALIGN_XZ_AXIS:
		r = this.B*v.Y - this.D
	case ALIGN_XY_AXIS:
		r = this.C*v.Z - this.D
	case ALIGN_ANY:
		r = this.A*v.X + this.B*v.Y + this.C*v.Z - this.D
	}
	return r
}
func (this *Plane) ClassifyPoint(v *Vector4, epsilon float32) string {
	var len float32
	switch this.alignment {
	case ALIGN_YZ_AXIS:
		len = this.A*v.X - this.D
	case ALIGN_XZ_AXIS:
		len = this.B*v.Y - this.D
	case ALIGN_XY_AXIS:
		len = this.C*v.Z - this.D
	case ALIGN_ANY:
		len = this.A*v.X + this.B*v.Y + this.C*v.Z - this.D
	}
	if len < -epsilon {
		return BACK
	} else if len > epsilon {
		return FRONT
	} else {
		return INTERSECT
	}
}
