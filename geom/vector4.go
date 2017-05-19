package geom

import (
	"math"
)

var VERCTOR4_X_AXIS = Vector4{X: 1.0, Y: 0.0, Z: 0.0, W: 1.0}
var VERCTOR4_Y_AXIS = Vector4{X: 0.0, Y: 1.0, Z: 0.0, W: 1.0}
var VERCTOR4_Z_AXIS = Vector4{X: 0.0, Y: 0.0, Z: 1.0, W: 1.0}

func Vector4AngleBetween(a *Vector4, b *Vector4) float32 {
	la := a.Length()
	lb := b.Length()
	dot := a.DotProduct(b)
	if la != 0 {
		dot /= la
	}
	if lb != 0 {
		dot /= lb
	}
	return float32(math.Acos(float64(dot)))
}
func Vector4Distance(a *Vector4, b *Vector4) float32 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return float32(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}
func Vector4Equals(a *Vector4, b *Vector4) bool {
	if FloatEqual(a.X, b.X) && FloatEqual(a.Y, b.Y) && FloatEqual(a.Z, b.Z) && FloatEqual(a.W, b.W) {
		return true
	}
	return false
}

type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (this *Vector4) Vector4() {
	this.X = 0.0
	this.Y = 0.0
	this.Z = 0.0
	this.W = 1.0
}
func (this *Vector4) Zero() {
	this.X = 0.0
	this.Y = 0.0
	this.Z = 0.0
	this.W = 1.0
}
func (this *Vector4) Length() float32 {
	return float32(math.Sqrt(float64(this.LengthSquared())))
}
func (this *Vector4) LengthSquared() float32 {
	return float32(this.X*this.X + this.Y*this.Y + this.Z*this.Z)
}
func (this *Vector4) Add(v *Vector4) {
	this.X += v.X
	this.Y += v.Y
	this.Z += v.Z
}
func (this *Vector4) Subtract(v *Vector4) {
	this.X -= v.X
	this.Y -= v.Y
	this.Z -= v.Z
}
func (this *Vector4) Mul(value float32) {
	this.X *= value
	this.Y *= value
	this.Z *= value
}
func (this *Vector4) Div(value float32) {
	var n float32 = 1 / value
	this.X *= n
	this.Y *= n
	this.Z *= n
}
func (this *Vector4) Normalize() {
	magSQ := this.X*this.X + this.Y*this.Y + this.Z*this.Z
	if magSQ > 0.0 {
		oneOverMag := 1.0 / float32(math.Sqrt(float64(magSQ)))
		this.X = this.X * oneOverMag
		this.Y = this.Y * oneOverMag
		this.Z = this.Z * oneOverMag
	}
}
func (this *Vector4) CrossProduct(v *Vector4) Vector4 {
	return Vector4{X: this.Y*v.Z - this.Z*v.Y, Y: this.Z*v.X - this.X*v.Z, Z: this.X*v.Y - this.Y*v.X, W: 1.0}
}
func (this *Vector4) DotProduct(v *Vector4) float32 {
	return this.X*v.X + this.Y*v.Y + this.Z*v.Z
}
func (this *Vector4) Project() {
	oneOverW := 1.0 / this.W
	this.X = this.X * oneOverW
	this.Y = this.Y * oneOverW
	this.Z = this.Z * oneOverW
	this.W = 1.0
}
func (this *Vector4) Negate() Vector4 {
	return Vector4{X: -this.X, Y: -this.Y, Z: -this.Z, W: 1.0}
}
func (this *Vector4) Scale(v float32) {
	this.X = this.X * v
	this.Y = this.Y * v
	this.Z = this.Z * v
}
func (this *Vector4) Clone() Vector4 {
	var v Vector4
	v.X = this.X
	v.Y = this.Y
	v.Z = this.Z
	v.W = this.W
	return v
}
