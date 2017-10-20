package base

import (
	"fmt"

	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/seven/geom"
)

//++++++++++++++++++++ Object ++++++++++++++++++++

type Object struct {
	events.EventDispatcher

	id   string
	name string

	X         float32
	Y         float32
	Z         float32
	RotationX float32
	RotationY float32
	RotationZ float32
	ScaleX    float32
	ScaleY    float32
	ScaleZ    float32

	position *geom.Vector4
	rotation *geom.Vector4
	scale    *geom.Vector4

	transform *geom.Matrix4x4
	changed   bool
}

func (this *Object) Object() {
	this.EventDispatcher.EventDispatcher()

	this.X = 0
	this.Y = 0
	this.Z = 0
	this.RotationX = 0
	this.RotationY = 0
	this.RotationZ = 0
	this.ScaleX = 1
	this.ScaleY = 1
	this.ScaleZ = 1

	this.position = new(geom.Vector4)
	this.position.Vector4()
	this.rotation = new(geom.Vector4)
	this.rotation.Vector4()
	this.scale = new(geom.Vector4)
	this.scale.Vector4()
	this.scale.X = 1.0
	this.scale.Y = 1.0
	this.scale.Z = 1.0

	this.transform = new(geom.Matrix4x4)
	this.transform.Matrix4x4(nil)
}
func (this *Object) GetID() string {
	return this.id
}
func (this *Object) GetName() string {
	return this.name
}
func (this *Object) SetName(name string) {
	this.name = name
}

func (this *Object) GetPosition() *geom.Vector4 {
	this.position.X = this.X
	this.position.Y = this.Y
	this.position.Z = this.Z
	return this.position
}
func (this *Object) GetRotation() *geom.Vector4 {
	this.rotation.X = this.RotationX
	this.rotation.Y = this.RotationY
	this.rotation.Z = this.RotationZ
	return this.rotation
}
func (this *Object) GetScale() *geom.Vector4 {
	this.scale.X = this.ScaleX
	this.scale.Y = this.ScaleY
	this.scale.Z = this.ScaleZ
	return this.scale
}
func (this *Object) GetTransform() *geom.Matrix4x4 {
	return this.transform
}

func (this *Object) Update() {
	if this.position.X != this.X {
		fmt.Println("x", this.X)
		this.position.X = this.X
		this.changed = true
	}
	if this.position.Y != this.Y {
		fmt.Println("y", this.Y)
		this.position.Y = this.Y
		this.changed = true
	}
	if this.position.Z != this.Z {
		fmt.Println("z", this.Z)
		this.position.Z = this.Z
		this.changed = true
	}

	if this.rotation.X != this.RotationX {
		fmt.Println("rotationX", this.RotationX)
		this.rotation.X = this.RotationX
		this.changed = true
	}
	if this.rotation.Y != this.RotationY {
		fmt.Println("rotationY", this.RotationY)
		this.rotation.Y = this.RotationY
		this.changed = true
	}
	if this.rotation.Z != this.RotationZ {
		fmt.Println("rotationZ", this.RotationZ)
		this.rotation.Z = this.RotationZ
		this.changed = true
	}

	if this.scale.X != this.ScaleX {
		fmt.Println("scaleX", this.ScaleX)
		if this.ScaleX <= 0 {
			this.ScaleX = 0.00000000001
		}
		this.scale.X = this.ScaleX
		this.changed = true
	}
	if this.scale.Y != this.ScaleY {
		fmt.Println("scaleY", this.ScaleY)
		if this.ScaleY <= 0 {
			this.ScaleY = 0.00000000001
		}
		this.scale.Y = this.ScaleY
		this.changed = true
	}
	if this.scale.Z != this.ScaleZ {
		fmt.Println("scaleZ", this.ScaleZ)
		if this.ScaleZ <= 0 {
			this.ScaleZ = 0.00000000001
		}
		this.scale.Z = this.ScaleZ
		this.changed = true
	}

	if this.changed {
		components := [3]*geom.Vector4{}
		components[0] = this.position
		components[1] = this.rotation
		components[2] = this.scale
		this.transform.Recompose(components, geom.AXIS_ANGLE)
		this.changed = false
	}
}

//++++++++++++++++++++ DisplayObject ++++++++++++++++++++

type DisplayObject struct {
	Object

	root      core.IContainer
	parent    core.IContainer
	renderer  core.IRenderer
	layerMask int32

	projection *geom.Matrix4x4
}

func (this *DisplayObject) DisplayObject() {
	this.Object.Object()
}
func (this *DisplayObject) GetRoot() core.IContainer {
	return this.root
}
func (this *DisplayObject) SetRoot(root core.IContainer) {
	this.root = root
}
func (this *DisplayObject) GetParent() core.IContainer {
	return this.parent
}
func (this *DisplayObject) SetParent(parent core.IContainer) {
	this.parent = parent
}
func (this *DisplayObject) GetLayerMask() int32 {
	return this.layerMask
}
func (this *DisplayObject) SetLayerMask(mask int32) {
	this.layerMask = mask
}
func (this *DisplayObject) GetRenderer() core.IRenderer {
	return this.renderer
}
func (this *DisplayObject) SetRenderer(renderer core.IRenderer) {
	this.renderer = renderer
}
func (this *DisplayObject) Update(transform *geom.Matrix4x4) {
	if transform != nil {
		this.transform.Append(transform)
	}
}
func (this *DisplayObject) Render(projection *geom.Matrix4x4) {
	if projection != nil {
		this.projection = projection
	}
}
