package geom

import (
	"math"
)

type Vector3 struct {
	X float32
	Y float32
	W float32
}

func (this *Vector3) Vector3() {
	this.X = 0.0
	this.Y = 0.0
	this.W = 1.0
}
func (this *Vector3) Length() float32 {
	return float32(math.Sqrt(float64(this.X*this.X + this.Y*this.Y)))
}
func (this *Vector3) Normalize() {
	magSQ := this.X*this.X + this.Y*this.Y
	if magSQ > 0.0 {
		oneOverMag := 1.0 / float32(math.Sqrt(float64(magSQ)))
		this.X = this.X * oneOverMag
		this.Y = this.Y * oneOverMag
	}
}
func (this *Vector3) Add(v *Vector3) {
	this.X += v.X
	this.Y += v.Y
}
func (this *Vector3) Sub(v *Vector3) {
	this.X -= v.X
	this.Y -= v.Y
}
