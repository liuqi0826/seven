package display

import (
	"math"

	"github.com/liuqi0826/seven/engine/display/base"
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/pick"
	"github.com/liuqi0826/seven/geom"
)

const PROJECTION_TYPE_ORTHO = "projectionTypeOrtho"
const PROJECTION_TYPE_PERSPECTIVE = "projectionTypePerspective"

const COORDINATE_SYSTEM_LEFT_HAND = "coordinateSystemLeftHand"
const COORDINATE_SYSTEM_RIGHT_HAND = "coordinateSystemRightHand"

type ProjectionConfig struct {
	ProjectType      string
	NearClipping     float32
	FarClipping      float32
	Field            float32
	CoordinateSystem string
}

type Camera struct {
	base.Object

	Controller  core.IController
	MousePicker pick.Picker

	config *ProjectionConfig

	host *Viewport

	upAxis     geom.Vector4
	projection geom.Matrix4x4
	tranform   geom.Matrix4x4
}

func (this *Camera) Camera(host *Viewport, config *ProjectionConfig) {
	this.Object.Object()
	this.host = host
	this.config = config
	this.MousePicker.Picker()
	this.createProjectionMatrix()
}
func (this *Camera) Update() {
	//	this.Controller.Update()
}
func (this *Camera) LookAt(at *geom.Vector4, up *geom.Vector4) {
	if up == nil {
		this.upAxis = geom.VERCTOR4_Y_AXIS
	} else {
		this.upAxis = up.Clone()
	}
	switch this.config.CoordinateSystem {
	case COORDINATE_SYSTEM_LEFT_HAND:
		zAxis := at.Clone()
		zAxis.Subtract(&this.Position)
		zAxis.Normalize()
		xAxis := up.Clone()
		xAxis.CrossProduct(&zAxis)
		xAxis.Normalize()
		yAxis := zAxis.Clone()
		yAxis.CrossProduct(&xAxis)
		xm := -xAxis.DotProduct(&this.Position)
		ym := -yAxis.DotProduct(&this.Position)
		zm := -zAxis.DotProduct(&this.Position)
		raw := [16]float32{
			xAxis.X, yAxis.X, zAxis.X, 0.0,
			xAxis.Y, yAxis.Y, zAxis.Y, 0.0,
			xAxis.Z, yAxis.Z, zAxis.Z, 0.0,
			xm, ym, zm, 1.0}
		this.tranform.Matrix4x4(&raw)
	case COORDINATE_SYSTEM_RIGHT_HAND:
		zAxis := this.Position.Clone()
		zAxis.Subtract(at)
		zAxis.Normalize()
		xAxis := up.Clone()
		xAxis.CrossProduct(&zAxis)
		xAxis.Normalize()
		yAxis := zAxis.Clone()
		yAxis.CrossProduct(&xAxis)
		xm := -xAxis.DotProduct(&this.Position)
		ym := -yAxis.DotProduct(&this.Position)
		zm := -zAxis.DotProduct(&this.Position)
		raw := [16]float32{
			xAxis.X, yAxis.X, zAxis.X, 0.0,
			xAxis.Y, yAxis.Y, zAxis.Y, 0.0,
			xAxis.Z, yAxis.Z, zAxis.Z, 0.0,
			xm, ym, zm, 1.0}
		this.tranform.Matrix4x4(&raw)
	}
}
func (this *Camera) createProjectionMatrix() {
	var raw [16]float32
	switch this.config.ProjectType {
	case PROJECTION_TYPE_ORTHO:
		switch this.config.CoordinateSystem {
		case COORDINATE_SYSTEM_LEFT_HAND:
			raw = [16]float32{
				2.0 / float32(this.host.Width), 0.0, 0.0, 0.0,
				0.0, 2.0 / float32(this.host.Height), 0.0, 0.0,
				0.0, 0.0, 1.0 / (this.config.FarClipping - this.config.NearClipping), 0.0,
				0.0, 0.0, this.config.NearClipping / (this.config.NearClipping - this.config.FarClipping), 1.0}
		case COORDINATE_SYSTEM_RIGHT_HAND:
			raw = [16]float32{
				2.0 / float32(this.host.Width), 0.0, 0.0, 0.0,
				0.0, 2.0 / float32(this.host.Height), 0.0, 0.0,
				0.0, 0.0, 1.0 / (this.config.NearClipping - this.config.FarClipping), 0.0,
				0.0, 0.0, this.config.NearClipping / (this.config.NearClipping - this.config.FarClipping), 1.0}
		}
	case PROJECTION_TYPE_PERSPECTIVE:
		aspectRatio := float32(this.host.Width) / float32(this.host.Height)
		yScale := 1.0 / float32(math.Tan(float64(this.config.Field/2.0)))
		xScale := yScale / aspectRatio
		switch this.config.CoordinateSystem {
		case COORDINATE_SYSTEM_LEFT_HAND:
			raw = [16]float32{
				xScale, 0.0, 0.0, 0.0,
				0.0, yScale, 0.0, 0.0,
				0.0, 0.0, this.config.FarClipping / (this.config.FarClipping - this.config.NearClipping), 1.0,
				0.0, 0.0, (this.config.NearClipping * this.config.FarClipping) / (this.config.NearClipping - this.config.FarClipping), 0.0}
		case COORDINATE_SYSTEM_RIGHT_HAND:
			raw = [16]float32{
				xScale, 0.0, 0.0, 0.0,
				0.0, yScale, 0.0, 0.0,
				0.0, 0.0, this.config.FarClipping / (this.config.NearClipping - this.config.FarClipping), -1.0,
				0.0, 0.0, (this.config.NearClipping * this.config.FarClipping) / (this.config.NearClipping - this.config.FarClipping), 0.0}
		}
	}
	this.projection = geom.Matrix4x4{}
	this.projection.Matrix4x4(&raw)
}
