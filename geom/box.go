package geom

import ()

//  4----5
// /|   /|
//0----1 |
//| 7--|-6
//3----2

type Box struct {
	Mix Vector4
	Max Vector4
}

func (this *Box) Box(mix, max *Vector4) {
	this.Mix = mix.Clone()
	this.Max = max.Clone()
}
func (this *Box) GetCenter() *Vector4 {
	center := new(Vector4)
	center.X = this.Max.X - this.Mix.X
	center.Y = this.Max.Y - this.Mix.Y
	center.Z = this.Max.Z - this.Mix.Z
	center.W = 1.0
	return center
}
func (this *Box) GetRadius() float32 {
	return Vector4Distance(&this.Max, &this.Mix)
}
func (this *Box) GetVertex() [8]*Vector4 {
	var list [8]*Vector4
	list[0] = &Vector4{X: this.Mix.X, Y: this.Max.Y, Z: this.Mix.Z}
	list[1] = &Vector4{X: this.Max.X, Y: this.Max.Y, Z: this.Mix.Z}
	list[2] = &Vector4{X: this.Max.X, Y: this.Mix.Y, Z: this.Mix.Z}
	list[3] = &Vector4{X: this.Mix.X, Y: this.Mix.Y, Z: this.Mix.Z}
	list[4] = &Vector4{X: this.Mix.X, Y: this.Max.Y, Z: this.Max.Z}
	list[5] = &Vector4{X: this.Max.X, Y: this.Max.Y, Z: this.Max.Z}
	list[6] = &Vector4{X: this.Max.X, Y: this.Mix.Y, Z: this.Max.Z}
	list[7] = &Vector4{X: this.Mix.X, Y: this.Mix.Y, Z: this.Max.Z}
	return list
}
