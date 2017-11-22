package input

import (
	//"fmt"

	"github.com/liuqi0826/seven/events"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type MouseEvent struct {
	altKey     bool
	ctrlKey    bool
	shiftKey   bool
	commandKey bool
	buttonDown bool
	delta      int32
	movementX  float32
	movementY  float32
	stageX     float32
	stageY     float32
}

func (this *MouseEvent) AltKey() bool {
	return this.altKey
}
func (this *MouseEvent) CtrlKey() bool {
	return this.ctrlKey
}
func (this *MouseEvent) ShiftKey() bool {
	return this.shiftKey
}
func (this *MouseEvent) CommandKey() bool {
	return this.commandKey
}
func (this *MouseEvent) ButtonDown() bool {
	return this.buttonDown
}
func (this *MouseEvent) Delta() int32 {
	return this.delta
}
func (this *MouseEvent) MovementX() float32 {
	return this.movementX
}
func (this *MouseEvent) MovementY() float32 {
	return this.movementY
}
func (this *MouseEvent) StageX() float32 {
	return this.stageX
}
func (this *MouseEvent) StageY() float32 {
	return this.stageY
}

type MouseManager struct {
	events.EventDispatcher

	hasMouseDown       bool
	hasRigthMouseDown  bool
	hasMiddleMouseDown bool
	click              bool
	rightClick         bool
	middleClick        bool

	buttonDown bool
	lastMouseX float32
	mouseX     float32
	lastMouseY float32
	mouseY     float32
	window     *glfw.Window
}

func (this *MouseManager) Setup(window *glfw.Window) {
	this.EventDispatcher.EventDispatcher(this)

	this.window = window
	this.window.SetMouseButtonCallback(this.mouseButtonCallback)
	this.window.SetCursorPosCallback(this.cursorPositionCallback)
	this.window.SetScrollCallback(this.scrollCallback)
}
func (this *MouseManager) mouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if button == 0 {
		switch action {
		case 0:
			this.buttonDown = false
		case 1:
			this.buttonDown = true
		}
	}

	mouseEvent := new(MouseEvent)
	mouseEvent.buttonDown = this.buttonDown
	mouseEvent.movementX = this.lastMouseX - this.mouseX
	mouseEvent.movementY = this.lastMouseY - this.mouseY
	mouseEvent.stageX = this.mouseX
	mouseEvent.stageY = this.mouseY
	switch mod {
	case 1:
		mouseEvent.shiftKey = true
	case 2:
		mouseEvent.ctrlKey = true
		mouseEvent.commandKey = true
	case 3:
		mouseEvent.shiftKey = true
		mouseEvent.ctrlKey = true
		mouseEvent.commandKey = true
	case 4:
		mouseEvent.altKey = true
	case 5:
		mouseEvent.shiftKey = true
		mouseEvent.altKey = true
	case 6:
		mouseEvent.ctrlKey = true
		mouseEvent.commandKey = true
		mouseEvent.altKey = true
	case 7:
		mouseEvent.shiftKey = true
		mouseEvent.ctrlKey = true
		mouseEvent.commandKey = true
		mouseEvent.altKey = true
	}

	event := new(events.Event)
	event.Data = mouseEvent
	switch button {
	case 0:
		switch action {
		case 0:
			event.Type = events.MOUSE_UP
			if this.hasMouseDown {
				this.click = true
				this.hasMouseDown = false
			}
		case 1:
			event.Type = events.MOUSE_DOWN
			this.hasMouseDown = true
		}
	case 1:
		switch action {
		case 0:
			event.Type = events.RIGHT_MOUSE_UP
			if this.hasRigthMouseDown {
				this.rightClick = true
				this.hasRigthMouseDown = false
			}
		case 1:
			event.Type = events.RIGHT_MOUSE_DOWN
			this.hasRigthMouseDown = true
		}
	case 2:
		switch action {
		case 0:
			event.Type = events.MIDDLE_MOUSE_UP
			if this.hasMiddleMouseDown {
				this.middleClick = true
				this.hasMiddleMouseDown = false
			}
		case 1:
			event.Type = events.MIDDLE_MOUSE_DOWN
			this.hasMiddleMouseDown = true
		}
	case 3:
		switch action {
		case 0:
		case 1:
		}
	case 4:
		switch action {
		case 0:
		case 1:
		}
	}
	this.DispatchEvent(event)

	if this.click {
		event := new(events.Event)
		event.Data = mouseEvent
		event.Type = events.CLICK
		this.DispatchEvent(event)
		this.click = false
	}
	if this.rightClick {
		event := new(events.Event)
		event.Data = mouseEvent
		event.Type = events.RIGHT_CLICK
		this.DispatchEvent(event)
		this.rightClick = false
	}
	if this.middleClick {
		event := new(events.Event)
		event.Data = mouseEvent
		event.Type = events.MIDDLE_CLICK
		this.DispatchEvent(event)
		this.middleClick = false
	}
}
func (this *MouseManager) cursorPositionCallback(window *glfw.Window, x float64, y float64) {
	this.lastMouseX = this.mouseX
	this.lastMouseY = this.mouseY
	this.mouseX = float32(x)
	this.mouseY = float32(y)

	mouseEvent := new(MouseEvent)
	mouseEvent.buttonDown = this.buttonDown
	mouseEvent.movementX = this.lastMouseX - this.mouseX
	mouseEvent.movementY = this.lastMouseY - this.mouseY
	mouseEvent.stageX = this.mouseX
	mouseEvent.stageY = this.mouseY

	event := new(events.Event)
	event.Data = mouseEvent
	event.Type = events.MOUSE_MOVE
	this.DispatchEvent(event)
}
func (this *MouseManager) scrollCallback(window *glfw.Window, x float64, y float64) {
	mouseEvent := new(MouseEvent)
	mouseEvent.buttonDown = this.buttonDown
	mouseEvent.movementX = this.lastMouseX - this.mouseX
	mouseEvent.movementY = this.lastMouseY - this.mouseY
	mouseEvent.stageX = this.mouseX
	mouseEvent.stageY = this.mouseY
	mouseEvent.delta = int32(y)

	event := new(events.Event)
	event.Data = mouseEvent
	event.Type = events.MOUSE_WHEEL
	this.DispatchEvent(event)
}
