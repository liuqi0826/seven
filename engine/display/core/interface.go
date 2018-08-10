package core

import (
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/seven/geom"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type IContext interface {
	events.IEventDispatcher

	Setup(window *glfw.Window, debug bool) error
	Clear(color bool, depth bool, stencil bool)
	CreateProgram() platform.IProgram3D
	CreateIndexBuffer() platform.IIndexBuffer
	CreateVertexBuffer() platform.IVertexBuffer
	SetClearColor(red float32, green float32, blue float32, alpha float32)
	SetDepthTest(depthMask bool, passCompareMode string)
	SetProgram(program platform.IProgram3D)
	SetVertexBufferAt(value string, stride int32, bufferOffset int32, format string)
	DrawTriangles(indexBuffer platform.IIndexBuffer, firstIndex int32, numTriangles int32)
	Present()
}

type ICamera interface {
	GetProjectionMatrix() *geom.Matrix4x4
	GetTransformMatrix() *geom.Matrix4x4
}

type IController interface {
	Update()
}

type IRenderer interface {
	Setup(camera ICamera, program3D platform.IProgram3D)
	SetCamera(camera ICamera)
	SetProgram(program platform.IProgram3D)
	SetValue(title string, value []float32)
	Render(renderable IRenderable)
}

type IRenderable interface {
	GetIndexBuffer() platform.IIndexBuffer
	GetVertexBuffer() *[8]platform.IVertexBuffer
}

type IEntity interface {
}

type IDisplayObject interface {
	events.IEventDispatcher

	GetID() string
	GetName() string
	SetName(string)
	GetRoot() IContainer
	SetRoot(root IContainer)
	GetParent() IContainer
	SetParent(parent IContainer)
	GetCamera() ICamera
	SetCamera(camera ICamera)
	GetLayerMask() int32
	SetLayerMask(int32)
	GetRenderer() IRenderer
	SetRenderer(renderer IRenderer)
	Update(transform *geom.Matrix4x4)
	Render()
}

type IContainer interface {
	events.IEventDispatcher

	AddChild(displayObject IDisplayObject)
	RemoveChild(displayObject IDisplayObject) IDisplayObject
	RemoveAllChildren()
	GetChildByName(name string) IDisplayObject
	GetChildrenNumber() int32
}

type IAnimation interface {
	GetType() string
	Update()
}
