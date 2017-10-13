package core

import (
	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/utils"
	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/seven/geom"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type IContext interface {
	Setup(config *utils.Config) error
	Clear(color bool, depth bool, stencil bool)
	CreateProgram() platform.IProgram3D
	CreateIndexBuffer() platform.IIndexBuffer
	CreateVertexBuffer() platform.IVertexBuffer
	SetClearColor(red float32, green float32, blue float32, alpha float32)
	SetDepthTest(depthMask bool, passCompareMode string)
	SetProgram(program platform.IProgram3D)
	SetVertexBufferAt(value string, stride int32, bufferOffset int, format string)
	GetWindow() *glfw.Window
	DrawTriangles(indexBuffer platform.IIndexBuffer, firstIndex int32, numTriangles int32)
	Present()
	ShouldClose() bool
}

type ICamera interface {
}

type IController interface {
	Update()
}

type IRenderer interface {
	Setup(program3D platform.IProgram3D)
	Render(renderable IRenderable)
}

type IRenderable interface {
	GetIndexBuffer() platform.IIndexBuffer
	GetVertexBuffer() *[8]platform.IVertexBuffer
	GetValueBuffer() []float32
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
	GetLayerMask() int32
	SetLayerMask(int32)
	GetRenderer() IRenderer
	SetRenderer(renderer IRenderer)
	Update(transform *geom.Matrix4x4)
	Render(projection *geom.Matrix4x4)
}

type IContainer interface {
	events.IEventDispatcher

	AddChild(displayObject IDisplayObject)
	RemoveChild(displayObject IDisplayObject) IDisplayObject
	RemoveAllChildren()
	GetChildByName(name string) IDisplayObject
	GetChildrenNumber() int32
}
