package geom

type Matrix3x3 struct {
	raw [9]float32
}

func (this *Matrix3x3) Matrix3x3() {
	this.Identity()
}

func (this *Matrix3x3) Identity() {
	this.raw = [9]float32{1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0}
}

func (this *Matrix3x3) Transpose() {

}

func (this *Matrix3x3) Invert() {

}

func (this *Matrix3x3) Append(lhs *Matrix3x3) {

}

func (this *Matrix3x3) Prepend(rhs *Matrix3x3) {

}

func (this *Matrix3x3) AppendRotation(angle float64, axis Vector4) {

}

func (this *Matrix3x3) TransformVector(v *Vector4) {

}

func (this *Matrix3x3) TransformVectorList(vl *[]Vector4) {

}

func (this *Matrix3x3) FromEulerAngles() {

}

func (this *Matrix3x3) FromQuaternion() {

}

func (this *Matrix3x3) Clone() Matrix3x3 {
	m := Matrix3x3{}
	return m
}

func (this *Matrix3x3) ToBinary() []byte {
	return make([]byte, 0, 0)
}

func (this *Matrix3x3) ToString() string {
	return ""
}
