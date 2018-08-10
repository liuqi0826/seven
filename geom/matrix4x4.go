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
func InterpolateMatrix4x4(m1 *Matrix4x4, m2 *Matrix4x4, percent float32) *Matrix4x4 {
	arr := [16]float32{}
	for i := 0; i < 16; i++ {
		arr[i] = m1.Raw[i] + (m2.Raw[i]-m1.Raw[i])*percent
	}
	mtx := new(Matrix4x4)
	mtx.Matrix4x4(&arr)
	return mtx
}

type Matrix4x4 struct {
	Raw [16]float32
}

func (this *Matrix4x4) Matrix4x4(raw *[16]float32) {
	if raw != nil {
		for idx, v := range raw {
			this.Raw[idx] = v
		}
	} else {
		this.Identity()
	}
}
func (this *Matrix4x4) Identity() {
	this.Raw = [16]float32{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0}
}
func (this *Matrix4x4) Transpose() {
	raw := [16]float32{this.Raw[0], this.Raw[4], this.Raw[8], this.Raw[12], this.Raw[1], this.Raw[5], this.Raw[9], this.Raw[13], this.Raw[2], this.Raw[6], this.Raw[10], this.Raw[14], this.Raw[3], this.Raw[7], this.Raw[11], this.Raw[15]}
	this.Matrix4x4(&raw)
}
func (this *Matrix4x4) Invert() bool {
	d := this.Determinant()
	if math.Abs(float64(d)) > 0.00000000001 {
		d = 1 / d
		data := this.GetRaw()
		this.Raw[0] = d * (data[5]*(data[10]*data[15]-data[14]*data[11]) - data[9]*(data[6]*data[15]-data[14]*data[7]) + data[13]*(data[6]*data[11]-data[10]*data[7]))
		this.Raw[1] = -d * (data[1]*(data[10]*data[15]-data[14]*data[11]) - data[9]*(data[2]*data[15]-data[14]*data[3]) + data[13]*(data[2]*data[11]-data[10]*data[3]))
		this.Raw[2] = d * (data[1]*(data[6]*data[15]-data[14]*data[7]) - data[5]*(data[2]*data[15]-data[14]*data[3]) + data[13]*(data[2]*data[7]-data[6]*data[3]))
		this.Raw[3] = -d * (data[1]*(data[6]*data[11]-data[10]*data[7]) - data[5]*(data[2]*data[11]-data[10]*data[3]) + data[9]*(data[2]*data[7]-data[6]*data[3]))
		this.Raw[4] = -d * (data[4]*(data[10]*data[15]-data[14]*data[11]) - data[8]*(data[6]*data[15]-data[14]*data[7]) + data[12]*(data[6]*data[11]-data[10]*data[7]))
		this.Raw[5] = d * (data[0]*(data[10]*data[15]-data[14]*data[11]) - data[8]*(data[2]*data[15]-data[14]*data[3]) + data[12]*(data[2]*data[11]-data[10]*data[3]))
		this.Raw[6] = -d * (data[0]*(data[6]*data[15]-data[14]*data[7]) - data[4]*(data[2]*data[15]-data[14]*data[3]) + data[12]*(data[2]*data[7]-data[6]*data[3]))
		this.Raw[7] = d * (data[0]*(data[6]*data[11]-data[10]*data[7]) - data[4]*(data[2]*data[11]-data[10]*data[3]) + data[8]*(data[2]*data[7]-data[6]*data[3]))
		this.Raw[8] = d * (data[4]*(data[9]*data[15]-data[13]*data[11]) - data[8]*(data[5]*data[15]-data[13]*data[7]) + data[12]*(data[5]*data[11]-data[9]*data[7]))
		this.Raw[9] = -d * (data[0]*(data[9]*data[15]-data[13]*data[11]) - data[8]*(data[1]*data[15]-data[13]*data[3]) + data[12]*(data[1]*data[11]-data[9]*data[3]))
		this.Raw[10] = d * (data[0]*(data[5]*data[15]-data[13]*data[7]) - data[4]*(data[1]*data[15]-data[13]*data[3]) + data[12]*(data[1]*data[7]-data[5]*data[3]))
		this.Raw[11] = -d * (data[0]*(data[5]*data[11]-data[9]*data[7]) - data[4]*(data[1]*data[11]-data[9]*data[3]) + data[8]*(data[1]*data[7]-data[5]*data[3]))
		this.Raw[12] = -d * (data[4]*(data[9]*data[14]-data[13]*data[10]) - data[8]*(data[5]*data[14]-data[13]*data[6]) + data[12]*(data[5]*data[10]-data[9]*data[6]))
		this.Raw[13] = d * (data[0]*(data[9]*data[14]-data[13]*data[10]) - data[8]*(data[1]*data[14]-data[13]*data[2]) + data[12]*(data[1]*data[10]-data[9]*data[2]))
		this.Raw[14] = -d * (data[0]*(data[5]*data[14]-data[13]*data[6]) - data[4]*(data[1]*data[14]-data[13]*data[2]) + data[12]*(data[1]*data[6]-data[5]*data[2]))
		this.Raw[15] = d * (data[0]*(data[5]*data[10]-data[9]*data[6]) - data[4]*(data[1]*data[10]-data[9]*data[2]) + data[8]*(data[1]*data[6]-data[5]*data[2]))
		return true
	}
	return false
}
func (this *Matrix4x4) Append(lhs *Matrix4x4) {
	if lhs == nil {
		return
	}
	var data [16]float32 = [16]float32{this.Raw[0], this.Raw[1], this.Raw[2], this.Raw[3], this.Raw[4], this.Raw[5], this.Raw[6], this.Raw[7], this.Raw[8], this.Raw[9], this.Raw[10], this.Raw[11], this.Raw[12], this.Raw[13], this.Raw[14], this.Raw[15]}
	this.Raw[0] = data[0]*lhs.Raw[0] + data[1]*lhs.Raw[4] + data[2]*lhs.Raw[8] + data[3]*lhs.Raw[12]
	this.Raw[1] = data[0]*lhs.Raw[1] + data[1]*lhs.Raw[5] + data[2]*lhs.Raw[9] + data[3]*lhs.Raw[13]
	this.Raw[2] = data[0]*lhs.Raw[2] + data[1]*lhs.Raw[6] + data[2]*lhs.Raw[10] + data[3]*lhs.Raw[14]
	this.Raw[3] = data[0]*lhs.Raw[3] + data[1]*lhs.Raw[7] + data[2]*lhs.Raw[11] + data[3]*lhs.Raw[15]
	this.Raw[4] = data[4]*lhs.Raw[0] + data[5]*lhs.Raw[4] + data[6]*lhs.Raw[8] + data[7]*lhs.Raw[12]
	this.Raw[5] = data[4]*lhs.Raw[1] + data[5]*lhs.Raw[5] + data[6]*lhs.Raw[9] + data[7]*lhs.Raw[13]
	this.Raw[6] = data[4]*lhs.Raw[2] + data[5]*lhs.Raw[6] + data[6]*lhs.Raw[10] + data[7]*lhs.Raw[14]
	this.Raw[7] = data[4]*lhs.Raw[3] + data[5]*lhs.Raw[7] + data[6]*lhs.Raw[11] + data[7]*lhs.Raw[15]
	this.Raw[8] = data[8]*lhs.Raw[0] + data[9]*lhs.Raw[4] + data[10]*lhs.Raw[8] + data[11]*lhs.Raw[12]
	this.Raw[9] = data[8]*lhs.Raw[1] + data[9]*lhs.Raw[5] + data[10]*lhs.Raw[9] + data[11]*lhs.Raw[13]
	this.Raw[10] = data[8]*lhs.Raw[2] + data[9]*lhs.Raw[6] + data[10]*lhs.Raw[10] + data[11]*lhs.Raw[14]
	this.Raw[11] = data[8]*lhs.Raw[3] + data[9]*lhs.Raw[7] + data[10]*lhs.Raw[11] + data[11]*lhs.Raw[15]
	this.Raw[12] = data[12]*lhs.Raw[0] + data[13]*lhs.Raw[4] + data[14]*lhs.Raw[8] + data[15]*lhs.Raw[12]
	this.Raw[13] = data[12]*lhs.Raw[1] + data[13]*lhs.Raw[5] + data[14]*lhs.Raw[9] + data[15]*lhs.Raw[13]
	this.Raw[14] = data[12]*lhs.Raw[2] + data[13]*lhs.Raw[6] + data[14]*lhs.Raw[10] + data[15]*lhs.Raw[14]
	this.Raw[15] = data[12]*lhs.Raw[3] + data[13]*lhs.Raw[7] + data[14]*lhs.Raw[11] + data[15]*lhs.Raw[15]
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
	mtx.Raw[0] = x2 + (y2+z2)*cos
	mtx.Raw[1] = x*y*ccos + z*sin
	mtx.Raw[2] = x*z*ccos - y*sin
	mtx.Raw[4] = x*y*ccos - z*sin
	mtx.Raw[5] = y2 + (x2+z2)*cos
	mtx.Raw[6] = y*z*ccos + x*sin
	mtx.Raw[8] = x*z*ccos + y*sin
	mtx.Raw[9] = y*z*ccos - x*sin
	mtx.Raw[10] = z2 + (x2+y2)*cos
	mtx.Raw[12] = (tx*(y2+z2)-x*(ty*y+tz*z))*ccos + (ty*z-tz*y)*sin
	mtx.Raw[13] = (ty*(x2+z2)-y*(tx*x+tz*z))*ccos + (tz*x-tx*z)*sin
	mtx.Raw[14] = (tz*(x2+y2)-z*(tx*x+ty*y))*ccos + (tx*y-ty*x)*sin
	this.Append(mtx)
}
func (this *Matrix4x4) ApendScale(xScale float32, yScale float32, zScale float32) {
	mtx := new(Matrix4x4)
	mtx.Matrix4x4(nil)
	mtx.Raw[0] = xScale
	mtx.Raw[5] = yScale
	mtx.Raw[10] = zScale
	this.Append(mtx)
}
func (this *Matrix4x4) AppendTranslation(x float32, y float32, z float32) {
	this.Raw[12] += x
	this.Raw[13] += y
	this.Raw[14] += z
}
func (this *Matrix4x4) Prepend(rhs *Matrix4x4) {
	if rhs == nil {
		return
	}
	var data [16]float32 = [16]float32{this.Raw[0], this.Raw[1], this.Raw[2], this.Raw[3], this.Raw[4], this.Raw[5], this.Raw[6], this.Raw[7], this.Raw[8], this.Raw[9], this.Raw[10], this.Raw[11], this.Raw[12], this.Raw[13], this.Raw[14], this.Raw[15]}
	this.Raw[0] = rhs.Raw[0]*data[0] + rhs.Raw[1]*data[4] + rhs.Raw[2]*data[8] + rhs.Raw[3]*data[12]
	this.Raw[1] = rhs.Raw[0]*data[1] + rhs.Raw[1]*data[5] + rhs.Raw[2]*data[9] + rhs.Raw[3]*data[13]
	this.Raw[2] = rhs.Raw[0]*data[2] + rhs.Raw[1]*data[6] + rhs.Raw[2]*data[10] + rhs.Raw[3]*data[14]
	this.Raw[3] = rhs.Raw[0]*data[3] + rhs.Raw[1]*data[7] + rhs.Raw[2]*data[11] + rhs.Raw[3]*data[15]
	this.Raw[4] = rhs.Raw[4]*data[0] + rhs.Raw[5]*data[4] + rhs.Raw[6]*data[8] + rhs.Raw[7]*data[12]
	this.Raw[5] = rhs.Raw[4]*data[1] + rhs.Raw[5]*data[5] + rhs.Raw[6]*data[9] + rhs.Raw[7]*data[13]
	this.Raw[6] = rhs.Raw[4]*data[2] + rhs.Raw[5]*data[6] + rhs.Raw[6]*data[10] + rhs.Raw[7]*data[14]
	this.Raw[7] = rhs.Raw[4]*data[3] + rhs.Raw[5]*data[7] + rhs.Raw[6]*data[11] + rhs.Raw[7]*data[15]
	this.Raw[8] = rhs.Raw[8]*data[0] + rhs.Raw[9]*data[4] + rhs.Raw[10]*data[8] + rhs.Raw[11]*data[12]
	this.Raw[9] = rhs.Raw[8]*data[1] + rhs.Raw[9]*data[5] + rhs.Raw[10]*data[9] + rhs.Raw[11]*data[13]
	this.Raw[10] = rhs.Raw[8]*data[2] + rhs.Raw[9]*data[6] + rhs.Raw[10]*data[10] + rhs.Raw[11]*data[14]
	this.Raw[11] = rhs.Raw[8]*data[3] + rhs.Raw[9]*data[7] + rhs.Raw[10]*data[11] + rhs.Raw[11]*data[15]
	this.Raw[12] = rhs.Raw[12]*data[0] + rhs.Raw[13]*data[4] + rhs.Raw[14]*data[8] + rhs.Raw[15]*data[12]
	this.Raw[13] = rhs.Raw[12]*data[1] + rhs.Raw[13]*data[5] + rhs.Raw[14]*data[9] + rhs.Raw[15]*data[13]
	this.Raw[14] = rhs.Raw[12]*data[2] + rhs.Raw[13]*data[6] + rhs.Raw[14]*data[10] + rhs.Raw[15]*data[14]
	this.Raw[15] = rhs.Raw[12]*data[3] + rhs.Raw[13]*data[7] + rhs.Raw[14]*data[11] + rhs.Raw[15]*data[15]
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
	mtx.Raw[0] = x2 + (y2+z2)*cos
	mtx.Raw[1] = x*y*ccos + z*sin
	mtx.Raw[2] = x*z*ccos - y*sin
	mtx.Raw[4] = x*y*ccos - z*sin
	mtx.Raw[5] = y2 + (x2+z2)*cos
	mtx.Raw[6] = y*z*ccos + x*sin
	mtx.Raw[8] = x*z*ccos + y*sin
	mtx.Raw[9] = y*z*ccos - x*sin
	mtx.Raw[10] = z2 + (x2+y2)*cos
	mtx.Raw[12] = (tx*(y2+z2)-x*(ty*y+tz*z))*ccos + (ty*z-tz*y)*sin
	mtx.Raw[13] = (ty*(x2+z2)-y*(tx*x+tz*z))*ccos + (tz*x-tx*z)*sin
	mtx.Raw[14] = (tz*(x2+y2)-z*(tx*x+ty*y))*ccos + (tx*y-ty*x)*sin
	this.Prepend(mtx)
}
func (this *Matrix4x4) PrependScale(xScale float32, yScale float32, zScale float32) {
	mtx := new(Matrix4x4)
	mtx.Matrix4x4(nil)
	mtx.Raw[0] = xScale
	mtx.Raw[5] = yScale
	mtx.Raw[10] = zScale
	this.Prepend(mtx)
}
func (this *Matrix4x4) PrependTranslation(x float32, y float32, z float32) {
	mtx := new(Matrix4x4)
	mtx.Matrix4x4(nil)
	mtx.Raw[12] = x
	mtx.Raw[13] = y
	mtx.Raw[14] = z
	this.Prepend(mtx)
}
func (this *Matrix4x4) DeltaTransformVector(v *Vector4) *Vector4 {
	x, y, z := v.X, v.Y, v.Z
	return &Vector4{X: x*this.Raw[0] + y*this.Raw[4] + z*this.Raw[8] + this.Raw[3], Y: x*this.Raw[1] + y*this.Raw[5] + z*this.Raw[9] + this.Raw[7], Z: x*this.Raw[2] + y*this.Raw[6] + z*this.Raw[10] + this.Raw[11], W: 1.0}
}
func (this *Matrix4x4) TransformVector(v *Vector4) *Vector4 {
	var target Vector4
	target.X = v.X*this.Raw[0] + v.Y*this.Raw[4] + v.Z*this.Raw[8]
	target.Y = v.X*this.Raw[1] + v.Y*this.Raw[5] + v.Z*this.Raw[9]
	target.Z = v.X*this.Raw[2] + v.Y*this.Raw[6] + v.Z*this.Raw[10]
	target.W = v.X*this.Raw[3] + v.Y*this.Raw[7] + v.Z*this.Raw[11]
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
	scale := [12]float32{
		components[2].X, components[2].X, components[2].X, 0,
		components[2].Y, components[2].Y, components[2].Y, 0,
		components[2].Z, components[2].Z, components[2].Z, 0}
	switch orientationStyle {
	case EULER_ANGLES:
		cx := float32(math.Cos(float64(components[1].X)))
		cy := float32(math.Cos(float64(components[1].Y)))
		cz := float32(math.Cos(float64(components[1].Z)))
		sx := float32(math.Sin(float64(components[1].X)))
		sy := float32(math.Sin(float64(components[1].Y)))
		sz := float32(math.Sin(float64(components[1].Z)))
		this.Raw[0] = cy * cz * scale[0]
		this.Raw[1] = cy * sz * scale[1]
		this.Raw[2] = -sy * scale[2]
		this.Raw[3] = 0
		this.Raw[4] = (sx*sy*cz - cx*sz) * scale[4]
		this.Raw[5] = (sx*sy*sz + cx*cz) * scale[5]
		this.Raw[6] = sx * cy * scale[6]
		this.Raw[7] = 0
		this.Raw[8] = (cx*sy*cz + sx*sz) * scale[8]
		this.Raw[9] = (cx*sy*sz - sx*cz) * scale[9]
		this.Raw[10] = cx * cy * scale[10]
		this.Raw[11] = 0
		this.Raw[12] = components[0].X
		this.Raw[13] = components[0].Y
		this.Raw[14] = components[0].Z
		this.Raw[15] = 1
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
		this.Raw[0] = (1 - 2*y*y - 2*z*z) * scale[0]
		this.Raw[1] = (2*x*y + 2*w*z) * scale[1]
		this.Raw[2] = (2*x*z - 2*w*y) * scale[2]
		this.Raw[3] = 0
		this.Raw[4] = (2*x*y - 2*w*z) * scale[4]
		this.Raw[5] = (1 - 2*x*x - 2*z*z) * scale[5]
		this.Raw[6] = (2*y*z + 2*w*x) * scale[6]
		this.Raw[7] = 0
		this.Raw[8] = (2*x*z + 2*w*y) * scale[8]
		this.Raw[9] = (2*y*z - 2*w*x) * scale[9]
		this.Raw[10] = (1 - 2*x*x - 2*y*y) * scale[10]
		this.Raw[11] = 0
		this.Raw[12] = components[0].X
		this.Raw[13] = components[0].Y
		this.Raw[14] = components[0].Z
		this.Raw[15] = 1
	}
	if components[2].X == 0 {
		this.Raw[0] = 1e-15
	}
	if components[2].Y == 0 {
		this.Raw[5] = 1e-15
	}
	if components[2].Z == 0 {
		this.Raw[10] = 1e-15
	}
	return !(components[2].X == 0 || components[2].Y == 0 || components[2].Y == 0)
}
func (this *Matrix4x4) Determinant() float32 {
	return ((this.Raw[0]*this.Raw[5]-this.Raw[4]*this.Raw[1])*(this.Raw[10]*this.Raw[15]-this.Raw[14]*this.Raw[11]) - (this.Raw[0]*this.Raw[9]-this.Raw[8]*this.Raw[1])*(this.Raw[6]*this.Raw[15]-this.Raw[14]*this.Raw[7]) + (this.Raw[0]*this.Raw[13]-this.Raw[12]*this.Raw[1])*(this.Raw[6]*this.Raw[11]-this.Raw[10]*this.Raw[7]) + (this.Raw[4]*this.Raw[9]-this.Raw[8]*this.Raw[5])*(this.Raw[2]*this.Raw[15]-this.Raw[14]*this.Raw[3]) - (this.Raw[4]*this.Raw[13]-this.Raw[12]*this.Raw[5])*(this.Raw[2]*this.Raw[11]-this.Raw[10]*this.Raw[3]) + (this.Raw[8]*this.Raw[13]-this.Raw[12]*this.Raw[9])*(this.Raw[2]*this.Raw[7]-this.Raw[6]*this.Raw[3]))
}
func (this *Matrix4x4) Position() *Vector4 {
	return &Vector4{X: this.Raw[12], Y: this.Raw[13], Z: this.Raw[14], W: 1.0}
}
func (this *Matrix4x4) GetRaw() [16]float32 {
	raw := [16]float32{this.Raw[0], this.Raw[1], this.Raw[2], this.Raw[3], this.Raw[4], this.Raw[5], this.Raw[6], this.Raw[7], this.Raw[8], this.Raw[9], this.Raw[10], this.Raw[11], this.Raw[12], this.Raw[13], this.Raw[14], this.Raw[15]}
	return raw
}
func (this *Matrix4x4) GetRawSlice() []float32 {
	raw := []float32{this.Raw[0], this.Raw[1], this.Raw[2], this.Raw[3], this.Raw[4], this.Raw[5], this.Raw[6], this.Raw[7], this.Raw[8], this.Raw[9], this.Raw[10], this.Raw[11], this.Raw[12], this.Raw[13], this.Raw[14], this.Raw[15]}
	return raw
}
func (this *Matrix4x4) Clone() Matrix4x4 {
	m := Matrix4x4{}
	m.Matrix4x4(&this.Raw)
	return m
}
func (this *Matrix4x4) ToArray() []float32 {
	array := this.Raw[:]
	return array
}
func (this *Matrix4x4) ToBinary() []byte {
	var byteArray []byte
	buff := bytes.NewBuffer(byteArray)
	for _, v := range this.Raw {
		binary.Write(buff, binary.BigEndian, v)
	}
	return byteArray
}
func (this *Matrix4x4) ToString() string {
	str := "Matrix4x4[ "
	for _, v := range this.Raw {
		str += fmt.Sprintf("%f ", v)
	}
	str += "]"
	return str
}
