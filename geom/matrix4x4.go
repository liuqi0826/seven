package geom

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func CreatePerspective(width float32, height float32, zNear float32, zFar float32) Matrix4x4 {
	var mtx Matrix4x4
	mtx.Matrix4x4(&[16]float32{2.0 * zNear / width, 0, 0, 0, 0, 2.0 * zNear / height, 0, 0, 0, 0, zFar / (zFar - zNear), zNear * zFar / (zNear - zFar), 0, 0, 1, 0})
	return mtx
}
func CreateOrtho(width float32, height float32, zNear float32, zFar float32) Matrix4x4 {
	var mtx Matrix4x4
	mtx.Matrix4x4(&[16]float32{2.0 / width, 0, 0, 0, 0, 2.0 / height, 0, 0, 0, 0, 1 / (zFar - zNear), zNear / (zNear - zFar), 0, 0, 0, 1})
	return mtx
}

type Matrix4x4 struct {
	raw [16]float32
}

func (this *Matrix4x4) Matrix4x4(raw *[16]float32) {
	if raw != nil {
		for idx, v := range raw {
			this.raw[idx] = v
		}
	} else {
		this.Identity()
	}
}
func (this *Matrix4x4) Identity() {
	this.raw = [16]float32{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0}
}
func (this *Matrix4x4) Transpose() {
	raw := [16]float32{this.raw[0], this.raw[4], this.raw[8], this.raw[12], this.raw[1], this.raw[5], this.raw[9], this.raw[13], this.raw[2], this.raw[6], this.raw[10], this.raw[14], this.raw[3], this.raw[7], this.raw[11], this.raw[15]}
	this.Matrix4x4(&raw)
}
func (this *Matrix4x4) Invert() bool {
	d := this.Determinant()
	if math.Abs(float64(d)) > 0.00000000001 {
		d = 1 / d
		data := this.GetRaw()
		this.raw[0] = d * (data[5]*(data[10]*data[15]-data[14]*data[11]) - data[9]*(data[6]*data[15]-data[14]*data[7]) + data[13]*(data[6]*data[11]-data[10]*data[7]))
		this.raw[1] = -d * (data[1]*(data[10]*data[15]-data[14]*data[11]) - data[9]*(data[2]*data[15]-data[14]*data[3]) + data[13]*(data[2]*data[11]-data[10]*data[3]))
		this.raw[2] = d * (data[1]*(data[6]*data[15]-data[14]*data[7]) - data[5]*(data[2]*data[15]-data[14]*data[3]) + data[13]*(data[2]*data[7]-data[6]*data[3]))
		this.raw[3] = -d * (data[1]*(data[6]*data[11]-data[10]*data[7]) - data[5]*(data[2]*data[11]-data[10]*data[3]) + data[9]*(data[2]*data[7]-data[6]*data[3]))
		this.raw[4] = -d * (data[4]*(data[10]*data[15]-data[14]*data[11]) - data[8]*(data[6]*data[15]-data[14]*data[7]) + data[12]*(data[6]*data[11]-data[10]*data[7]))
		this.raw[5] = d * (data[0]*(data[10]*data[15]-data[14]*data[11]) - data[8]*(data[2]*data[15]-data[14]*data[3]) + data[12]*(data[2]*data[11]-data[10]*data[3]))
		this.raw[6] = -d * (data[0]*(data[6]*data[15]-data[14]*data[7]) - data[4]*(data[2]*data[15]-data[14]*data[3]) + data[12]*(data[2]*data[7]-data[6]*data[3]))
		this.raw[7] = d * (data[0]*(data[6]*data[11]-data[10]*data[7]) - data[4]*(data[2]*data[11]-data[10]*data[3]) + data[8]*(data[2]*data[7]-data[6]*data[3]))
		this.raw[8] = d * (data[4]*(data[9]*data[15]-data[13]*data[11]) - data[8]*(data[5]*data[15]-data[13]*data[7]) + data[12]*(data[5]*data[11]-data[9]*data[7]))
		this.raw[9] = -d * (data[0]*(data[9]*data[15]-data[13]*data[11]) - data[8]*(data[1]*data[15]-data[13]*data[3]) + data[12]*(data[1]*data[11]-data[9]*data[3]))
		this.raw[10] = d * (data[0]*(data[5]*data[15]-data[13]*data[7]) - data[4]*(data[1]*data[15]-data[13]*data[3]) + data[12]*(data[1]*data[7]-data[5]*data[3]))
		this.raw[11] = -d * (data[0]*(data[5]*data[11]-data[9]*data[7]) - data[4]*(data[1]*data[11]-data[9]*data[3]) + data[8]*(data[1]*data[7]-data[5]*data[3]))
		this.raw[12] = -d * (data[4]*(data[9]*data[14]-data[13]*data[10]) - data[8]*(data[5]*data[14]-data[13]*data[6]) + data[12]*(data[5]*data[10]-data[9]*data[6]))
		this.raw[13] = d * (data[0]*(data[9]*data[14]-data[13]*data[10]) - data[8]*(data[1]*data[14]-data[13]*data[2]) + data[12]*(data[1]*data[10]-data[9]*data[2]))
		this.raw[14] = -d * (data[0]*(data[5]*data[14]-data[13]*data[6]) - data[4]*(data[1]*data[14]-data[13]*data[2]) + data[12]*(data[1]*data[6]-data[5]*data[2]))
		this.raw[15] = d * (data[0]*(data[5]*data[10]-data[9]*data[6]) - data[4]*(data[1]*data[10]-data[9]*data[2]) + data[8]*(data[1]*data[6]-data[5]*data[2]))
		return true
	}
	return false
}
func (this *Matrix4x4) Append(lhs *Matrix4x4) {
	var data [16]float32 = [16]float32{this.raw[0], this.raw[1], this.raw[2], this.raw[3], this.raw[4], this.raw[5], this.raw[6], this.raw[7], this.raw[8], this.raw[9], this.raw[10], this.raw[11], this.raw[12], this.raw[13], this.raw[14], this.raw[15]}
	this.raw[0] = data[0]*lhs.raw[0] + data[1]*lhs.raw[4] + data[2]*lhs.raw[8] + data[3]*lhs.raw[12]
	this.raw[1] = data[0]*lhs.raw[1] + data[1]*lhs.raw[5] + data[2]*lhs.raw[9] + data[3]*lhs.raw[13]
	this.raw[2] = data[0]*lhs.raw[2] + data[1]*lhs.raw[6] + data[2]*lhs.raw[10] + data[3]*lhs.raw[14]
	this.raw[3] = data[0]*lhs.raw[3] + data[1]*lhs.raw[7] + data[2]*lhs.raw[11] + data[3]*lhs.raw[15]
	this.raw[4] = data[4]*lhs.raw[0] + data[5]*lhs.raw[4] + data[6]*lhs.raw[8] + data[7]*lhs.raw[12]
	this.raw[5] = data[4]*lhs.raw[1] + data[5]*lhs.raw[5] + data[6]*lhs.raw[9] + data[7]*lhs.raw[13]
	this.raw[6] = data[4]*lhs.raw[2] + data[5]*lhs.raw[6] + data[6]*lhs.raw[10] + data[7]*lhs.raw[14]
	this.raw[7] = data[4]*lhs.raw[3] + data[5]*lhs.raw[7] + data[6]*lhs.raw[11] + data[7]*lhs.raw[15]
	this.raw[8] = data[8]*lhs.raw[0] + data[9]*lhs.raw[4] + data[10]*lhs.raw[8] + data[11]*lhs.raw[12]
	this.raw[9] = data[8]*lhs.raw[1] + data[9]*lhs.raw[5] + data[10]*lhs.raw[9] + data[11]*lhs.raw[13]
	this.raw[10] = data[8]*lhs.raw[2] + data[9]*lhs.raw[6] + data[10]*lhs.raw[10] + data[11]*lhs.raw[14]
	this.raw[11] = data[8]*lhs.raw[3] + data[9]*lhs.raw[7] + data[10]*lhs.raw[11] + data[11]*lhs.raw[15]
	this.raw[12] = data[12]*lhs.raw[0] + data[13]*lhs.raw[4] + data[14]*lhs.raw[8] + data[15]*lhs.raw[12]
	this.raw[13] = data[12]*lhs.raw[1] + data[13]*lhs.raw[5] + data[14]*lhs.raw[9] + data[15]*lhs.raw[13]
	this.raw[14] = data[12]*lhs.raw[2] + data[13]*lhs.raw[6] + data[14]*lhs.raw[10] + data[15]*lhs.raw[14]
	this.raw[15] = data[12]*lhs.raw[3] + data[13]*lhs.raw[7] + data[14]*lhs.raw[11] + data[15]*lhs.raw[15]
}
func (this *Matrix4x4) AppendRotation(degrees float32, axis *Vector4, pivotPoint *Vector4) {
	tx, ty, tz := float32(0), float32(0), float32(0)
	if pivotPoint != nil {
		tx, ty, tz = pivotPoint.X, pivotPoint.Y, pivotPoint.Z
	}
	radian := degrees * RADIANS_TO_DEGREES
	cos := float32(math.Cos(float64(radian)))
	sin := float32(math.Sin(float64(radian)))
	x, y, z := axis.X, axis.Y, axis.Z
	x2, y2, z2 := x*x, y*y, z*z
	ls := x2 + y2 + z2
	if ls != 0 {
		l := float32(math.Sqrt(float64(ls)))
		x, y, z, x2, y2, z2 = x/l, y/l, z/l, x2/ls, y2/ls, z2/ls
	}
	ccos := 1 - cos
	mtx := new(Matrix4x4)
	mtx.raw[0] = x2 + (y2+z2)*cos
	mtx.raw[1] = x*y*ccos + z*sin
	mtx.raw[2] = x*z*ccos - y*sin
	mtx.raw[4] = x*y*ccos - z*sin
	mtx.raw[5] = y2 + (x2+z2)*cos
	mtx.raw[6] = y*z*ccos + x*sin
	mtx.raw[8] = x*z*ccos + y*sin
	mtx.raw[9] = y*z*ccos - x*sin
	mtx.raw[10] = z2 + (x2+y2)*cos
	mtx.raw[12] = (tx*(y2+z2)-x*(ty*y+tz*z))*ccos + (ty*z-tz*y)*sin
	mtx.raw[13] = (ty*(x2+z2)-y*(tx*x+tz*z))*ccos + (tz*x-tx*z)*sin
	mtx.raw[14] = (tz*(x2+y2)-z*(tx*x+ty*y))*ccos + (tx*y-ty*x)*sin
	this.Append(mtx)
}
func (this *Matrix4x4) ApendScale(xScale float32, yScale float32, zScale float32) {
	mtx := new(Matrix4x4)
	mtx.Matrix4x4(nil)
	mtx.raw[0] = xScale
	mtx.raw[5] = yScale
	mtx.raw[10] = zScale
	this.Append(mtx)
}
func (this *Matrix4x4) AppendTranslation(x float32, y float32, z float32) {
	this.raw[12] += x
	this.raw[13] += y
	this.raw[14] += z
}
func (this *Matrix4x4) Prepend(rhs *Matrix4x4) {
	var data [16]float32 = [16]float32{this.raw[0], this.raw[1], this.raw[2], this.raw[3], this.raw[4], this.raw[5], this.raw[6], this.raw[7], this.raw[8], this.raw[9], this.raw[10], this.raw[11], this.raw[12], this.raw[13], this.raw[14], this.raw[15]}
	this.raw[0] = rhs.raw[0]*data[0] + rhs.raw[1]*data[4] + rhs.raw[2]*data[8] + rhs.raw[3]*data[12]
	this.raw[1] = rhs.raw[0]*data[1] + rhs.raw[1]*data[5] + rhs.raw[2]*data[9] + rhs.raw[3]*data[13]
	this.raw[2] = rhs.raw[0]*data[2] + rhs.raw[1]*data[6] + rhs.raw[2]*data[10] + rhs.raw[3]*data[14]
	this.raw[3] = rhs.raw[0]*data[3] + rhs.raw[1]*data[7] + rhs.raw[2]*data[11] + rhs.raw[3]*data[15]
	this.raw[4] = rhs.raw[4]*data[0] + rhs.raw[5]*data[4] + rhs.raw[6]*data[8] + rhs.raw[7]*data[12]
	this.raw[5] = rhs.raw[4]*data[1] + rhs.raw[5]*data[5] + rhs.raw[6]*data[9] + rhs.raw[7]*data[13]
	this.raw[6] = rhs.raw[4]*data[2] + rhs.raw[5]*data[6] + rhs.raw[6]*data[10] + rhs.raw[7]*data[14]
	this.raw[7] = rhs.raw[4]*data[3] + rhs.raw[5]*data[7] + rhs.raw[6]*data[11] + rhs.raw[7]*data[15]
	this.raw[8] = rhs.raw[8]*data[0] + rhs.raw[9]*data[4] + rhs.raw[10]*data[8] + rhs.raw[11]*data[12]
	this.raw[9] = rhs.raw[8]*data[1] + rhs.raw[9]*data[5] + rhs.raw[10]*data[9] + rhs.raw[11]*data[13]
	this.raw[10] = rhs.raw[8]*data[2] + rhs.raw[9]*data[6] + rhs.raw[10]*data[10] + rhs.raw[11]*data[14]
	this.raw[11] = rhs.raw[8]*data[3] + rhs.raw[9]*data[7] + rhs.raw[10]*data[11] + rhs.raw[11]*data[15]
	this.raw[12] = rhs.raw[12]*data[0] + rhs.raw[13]*data[4] + rhs.raw[14]*data[8] + rhs.raw[15]*data[12]
	this.raw[13] = rhs.raw[12]*data[1] + rhs.raw[13]*data[5] + rhs.raw[14]*data[9] + rhs.raw[15]*data[13]
	this.raw[14] = rhs.raw[12]*data[2] + rhs.raw[13]*data[6] + rhs.raw[14]*data[10] + rhs.raw[15]*data[14]
	this.raw[15] = rhs.raw[12]*data[3] + rhs.raw[13]*data[7] + rhs.raw[14]*data[11] + rhs.raw[15]*data[15]
}
func (this *Matrix4x4) PrependRotation(degrees float32, axis *Vector4, pivotPoint *Vector4) {
	tx, ty, tz := float32(0), float32(0), float32(0)
	if pivotPoint != nil {
		tx, ty, tz = pivotPoint.X, pivotPoint.Y, pivotPoint.Z
	}
	radian := degrees * RADIANS_TO_DEGREES
	cos := float32(math.Cos(float64(radian)))
	sin := float32(math.Sin(float64(radian)))
	x, y, z := axis.X, axis.Y, axis.Z
	x2, y2, z2 := x*x, y*y, z*z
	ls := x2 + y2 + z2
	if ls != 0 {
		l := float32(math.Sqrt(float64(ls)))
		x, y, z, x2, y2, z2 = x/l, y/l, z/l, x2/ls, y2/ls, z2/ls
	}
	ccos := 1 - cos
	mtx := new(Matrix4x4)
	mtx.raw[0] = x2 + (y2+z2)*cos
	mtx.raw[1] = x*y*ccos + z*sin
	mtx.raw[2] = x*z*ccos - y*sin
	mtx.raw[4] = x*y*ccos - z*sin
	mtx.raw[5] = y2 + (x2+z2)*cos
	mtx.raw[6] = y*z*ccos + x*sin
	mtx.raw[8] = x*z*ccos + y*sin
	mtx.raw[9] = y*z*ccos - x*sin
	mtx.raw[10] = z2 + (x2+y2)*cos
	mtx.raw[12] = (tx*(y2+z2)-x*(ty*y+tz*z))*ccos + (ty*z-tz*y)*sin
	mtx.raw[13] = (ty*(x2+z2)-y*(tx*x+tz*z))*ccos + (tz*x-tx*z)*sin
	mtx.raw[14] = (tz*(x2+y2)-z*(tx*x+ty*y))*ccos + (tx*y-ty*x)*sin
	this.Prepend(mtx)
}
func (this *Matrix4x4) PrependScale(xScale float32, yScale float32, zScale float32) {
	mtx := new(Matrix4x4)
	mtx.Matrix4x4(nil)
	mtx.raw[0] = xScale
	mtx.raw[5] = yScale
	mtx.raw[10] = zScale
	this.Prepend(mtx)
}
func (this *Matrix4x4) PrependTranslation(x float32, y float32, z float32) {
	mtx := new(Matrix4x4)
	mtx.Matrix4x4(nil)
	mtx.raw[12] = x
	mtx.raw[13] = y
	mtx.raw[14] = z
	this.Prepend(mtx)
}
func (this *Matrix4x4) DeltaTransformVector(v *Vector4) *Vector4 {
	x, y, z := v.X, v.Y, v.Z
	return &Vector4{X: x*this.raw[0] + y*this.raw[4] + z*this.raw[8] + this.raw[3], Y: x*this.raw[1] + y*this.raw[5] + z*this.raw[9] + this.raw[7], Z: x*this.raw[2] + y*this.raw[6] + z*this.raw[10] + this.raw[11], W: 1.0}
}
func (this *Matrix4x4) TransformVector(v *Vector4) *Vector4 {
	var target Vector4
	target.X = v.X*this.raw[0] + v.Y*this.raw[4] + v.Z*this.raw[8]
	target.Y = v.X*this.raw[1] + v.Y*this.raw[5] + v.Z*this.raw[9]
	target.Z = v.X*this.raw[2] + v.Y*this.raw[6] + v.Z*this.raw[10]
	target.W = v.X*this.raw[3] + v.Y*this.raw[7] + v.Z*this.raw[11]
	return &target
}
func (this *Matrix4x4) TransformVectorList(vl []*Vector4) []*Vector4 {
	var tl []*Vector4
	for _, v := range vl {
		tg := this.TransformVector(v)
		tl = append(tl, tg)
	}
	return tl
}
func (this *Matrix4x4) Decompose(orientationStyle string) [3]*Vector4 {
	vec := [3]*Vector4{}
	mr := this.GetRaw()
	pos := Vector4{X: mr[12], Y: mr[13], Z: mr[14], W: 1.0}
	mr[12], mr[13], mr[14] = 0, 0, 0
	scale := Vector4{}
	scale.X = float32(math.Sqrt(float64(mr[0]*mr[0] + mr[1]*mr[1] + mr[2]*mr[2])))
	scale.Y = float32(math.Sqrt(float64(mr[4]*mr[4] + mr[5]*mr[5] + mr[6]*mr[6])))
	scale.Z = float32(math.Sqrt(float64(mr[8]*mr[8] + mr[9]*mr[9] + mr[10]*mr[10])))
	if mr[0]*(mr[5]*mr[10]-mr[6]*mr[9])-mr[1]*(mr[4]*mr[10]-mr[6]*mr[8])+mr[2]*(mr[4]*mr[9]-mr[5]*mr[8]) < 0 {
		scale.Z = -scale.Z
	}
	ox := 1 / scale.X
	mr[0] *= ox
	mr[1] *= ox
	mr[2] *= ox
	oy := 1 / scale.Y
	mr[4] *= oy
	mr[5] *= oy
	mr[6] *= oy
	oz := 1 / scale.Z
	mr[8] *= oz
	mr[9] *= oz
	mr[10] *= oz
	rot := Vector4{}
	if orientationStyle != AXIS_ANGLE && orientationStyle != QUATERNION && orientationStyle != EULER_ANGLES {
		orientationStyle = EULER_ANGLES
	}
	switch orientationStyle {
	case AXIS_ANGLE:
		rot.W = float32(math.Acos(float64((mr[0] + mr[5] + mr[10] - 1) * .5)))
		var len = float32(math.Sqrt(float64((mr[6]-mr[9])*(mr[6]-mr[9]) + (mr[8]-mr[2])*(mr[8]-mr[2]) + (mr[1]-mr[4])*(mr[1]-mr[4]))))
		if len != 0 {
			ol := 1 / len
			rot.X = (mr[6] - mr[9]) * ol
			rot.Y = (mr[8] - mr[2]) * ol
			rot.Z = (mr[1] - mr[4]) * ol
		} else {
			rot.X, rot.Y, rot.Z = 0, 0, 0
		}
	case QUATERNION:
		var tr = mr[0] + mr[5] + mr[10]
		if tr > 0 {
			rot.W = float32(math.Sqrt(float64(1+tr) * .5))
			rot.X = (mr[6] - mr[9]) / (4 * rot.W)
			rot.Y = (mr[8] - mr[2]) / (4 * rot.W)
			rot.Z = (mr[1] - mr[4]) / (4 * rot.W)
		} else if (mr[0] > mr[5]) && (mr[0] > mr[10]) {
			rot.X = float32(math.Sqrt(float64(1+mr[0]-mr[5]-mr[10]) * .5))
			rot.W = (mr[6] - mr[9]) / (4 * rot.X)
			rot.Y = (mr[1] + mr[4]) / (4 * rot.X)
			rot.Z = (mr[8] + mr[2]) / (4 * rot.X)
		} else if mr[5] > mr[10] {
			rot.Y = float32(math.Sqrt(float64(1+mr[5]-mr[0]-mr[10]) * .5))
			rot.X = (mr[1] + mr[4]) / (4 * rot.Y)
			rot.W = (mr[8] - mr[2]) / (4 * rot.Y)
			rot.Z = (mr[6] + mr[9]) / (4 * rot.Y)
		} else {
			rot.Z = float32(math.Sqrt(float64(1+mr[10]-mr[0]-mr[5]) * .5))
			rot.X = (mr[8] + mr[2]) / (4 * rot.Z)
			rot.Y = (mr[6] + mr[9]) / (4 * rot.Z)
			rot.W = (mr[1] - mr[4]) / (4 * rot.Z)
		}
	case EULER_ANGLES:
		rot.Y = float32(math.Asin(float64(-mr[2])))
		if mr[2] != 1 && mr[2] != -1 {
			rot.X = float32(math.Atan2(float64(mr[6]), float64(mr[10])))
			rot.Z = float32(math.Atan2(float64(mr[1]), float64(mr[0])))
		} else {
			rot.Z = 0
			rot.X = float32(math.Atan2(float64(mr[4]), float64(mr[5])))
		}
	}
	vec[0] = &pos
	vec[1] = &rot
	vec[2] = &scale
	return vec
}
func (this *Matrix4x4) Recompose(components [3]*Vector4, orientationStyle string) bool {
	if components[2].X == 0 || components[2].Y == 0 || components[2].Z == 0 {
		return false
	}
	if orientationStyle != AXIS_ANGLE && orientationStyle != QUATERNION && orientationStyle != EULER_ANGLES {
		orientationStyle = EULER_ANGLES
	}
	this.Identity()
	scale := [12]float32{components[2].X, components[2].X, components[2].X, 0, components[2].Y, components[2].Y, components[2].Y, 0, components[2].Z, components[2].Z, components[2].Z, 0}
	switch orientationStyle {
	case EULER_ANGLES:
		cx := float32(math.Cos(float64(components[1].X)))
		cy := float32(math.Cos(float64(components[1].Y)))
		cz := float32(math.Cos(float64(components[1].Z)))
		sx := float32(math.Sin(float64(components[1].X)))
		sy := float32(math.Sin(float64(components[1].Y)))
		sz := float32(math.Sin(float64(components[1].Z)))
		this.raw[0] = cy * cz * scale[0]
		this.raw[1] = cy * sz * scale[1]
		this.raw[2] = -sy * scale[2]
		this.raw[3] = 0
		this.raw[4] = (sx*sy*cz - cx*sz) * scale[4]
		this.raw[5] = (sx*sy*sz + cx*cz) * scale[5]
		this.raw[6] = sx * cy * scale[6]
		this.raw[7] = 0
		this.raw[8] = (cx*sy*cz + sx*sz) * scale[8]
		this.raw[9] = (cx*sy*sz - sx*cz) * scale[9]
		this.raw[10] = cx * cy * scale[10]
		this.raw[11] = 0
		this.raw[12] = components[0].X
		this.raw[13] = components[0].Y
		this.raw[14] = components[0].Z
		this.raw[15] = 1
	default:
		x := components[1].X
		y := components[1].Y
		z := components[1].Z
		w := components[1].W
		if orientationStyle == AXIS_ANGLE {
			x *= float32(math.Sin(float64(w * .5)))
			y *= float32(math.Sin(float64(w * .5)))
			z *= float32(math.Sin(float64(w * .5)))
			w = float32(math.Cos(float64(w * .5)))
		}
		this.raw[0] = (1 - 2*y*y - 2*z*z) * scale[0]
		this.raw[1] = (2*x*y + 2*w*z) * scale[1]
		this.raw[2] = (2*x*z - 2*w*y) * scale[2]
		this.raw[3] = 0
		this.raw[4] = (2*x*y - 2*w*z) * scale[4]
		this.raw[5] = (1 - 2*x*x - 2*z*z) * scale[5]
		this.raw[6] = (2*y*z + 2*w*x) * scale[6]
		this.raw[7] = 0
		this.raw[8] = (2*x*z + 2*w*y) * scale[8]
		this.raw[9] = (2*y*z - 2*w*x) * scale[9]
		this.raw[10] = (1 - 2*x*x - 2*y*y) * scale[10]
		this.raw[11] = 0
		this.raw[12] = components[0].X
		this.raw[13] = components[0].Y
		this.raw[14] = components[0].Z
		this.raw[15] = 1
	}
	if components[2].X == 0 {
		this.raw[0] = 1e-15
	}
	if components[2].Y == 0 {
		this.raw[5] = 1e-15
	}
	if components[2].Z == 0 {
		this.raw[10] = 1e-15
	}
	return !(components[2].X == 0 || components[2].Y == 0 || components[2].Y == 0)
}
func (this *Matrix4x4) Determinant() float32 {
	return ((this.raw[0]*this.raw[5]-this.raw[4]*this.raw[1])*(this.raw[10]*this.raw[15]-this.raw[14]*this.raw[11]) - (this.raw[0]*this.raw[9]-this.raw[8]*this.raw[1])*(this.raw[6]*this.raw[15]-this.raw[14]*this.raw[7]) + (this.raw[0]*this.raw[13]-this.raw[12]*this.raw[1])*(this.raw[6]*this.raw[11]-this.raw[10]*this.raw[7]) + (this.raw[4]*this.raw[9]-this.raw[8]*this.raw[5])*(this.raw[2]*this.raw[15]-this.raw[14]*this.raw[3]) - (this.raw[4]*this.raw[13]-this.raw[12]*this.raw[5])*(this.raw[2]*this.raw[11]-this.raw[10]*this.raw[3]) + (this.raw[8]*this.raw[13]-this.raw[12]*this.raw[9])*(this.raw[2]*this.raw[7]-this.raw[6]*this.raw[3]))
}
func (this *Matrix4x4) Position() *Vector4 {
	return &Vector4{X: this.raw[12], Y: this.raw[13], Z: this.raw[14], W: 1.0}
}
func (this *Matrix4x4) GetRaw() [16]float32 {
	raw := [16]float32{this.raw[0], this.raw[1], this.raw[2], this.raw[3], this.raw[4], this.raw[5], this.raw[6], this.raw[7], this.raw[8], this.raw[9], this.raw[10], this.raw[11], this.raw[12], this.raw[13], this.raw[14], this.raw[15]}
	return raw
}
func (this *Matrix4x4) Clone() Matrix4x4 {
	m := Matrix4x4{}
	m.Matrix4x4(&this.raw)
	return m
}
func (this *Matrix4x4) ToBinary() []byte {
	var byteArray []byte
	buff := bytes.NewBuffer(byteArray)
	for _, v := range this.raw {
		binary.Write(buff, binary.BigEndian, v)
	}
	return byteArray
}
func (this *Matrix4x4) ToString() string {
	str := "Matrix4x4[ "
	for _, v := range this.raw {
		str += fmt.Sprintf("%f ", v)
	}
	str += "]"
	return str
}
