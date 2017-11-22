package input

import (
	//"fmt"

	"github.com/liuqi0826/seven/events"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type KeyboardEvent struct {
	altKey      bool
	ctrlKey     bool
	shiftKey    bool
	commandKey  bool
	scanCode    uint32
	keyCode     uint32
	keyLocation uint32
}

func (this *KeyboardEvent) AltKey() bool {
	return this.altKey
}
func (this *KeyboardEvent) CtrlKey() bool {
	return this.ctrlKey
}
func (this *KeyboardEvent) ShiftKey() bool {
	return this.shiftKey
}
func (this *KeyboardEvent) CommandKey() bool {
	return this.commandKey
}
func (this *KeyboardEvent) ScanCode() uint32 {
	return this.scanCode
}
func (this *KeyboardEvent) KeyCode() uint32 {
	return this.keyCode
}
func (this *KeyboardEvent) KeyLocation() uint32 {
	return this.keyLocation
}

type KeyboardManager struct {
	events.EventDispatcher

	altKey      bool
	ctrlKey     bool
	shiftKey    bool
	commandKey  bool
	scanCode    uint32
	keyCode     uint32
	keyLocation uint32
	window      *glfw.Window
}

func (this *KeyboardManager) Setup(window *glfw.Window) {
	this.EventDispatcher.EventDispatcher(this)

	this.window = window
	this.window.SetKeyCallback(this.keyboardCallback)
}
func (this *KeyboardManager) AltKey() bool {
	return this.altKey
}
func (this *KeyboardManager) CtrlKey() bool {
	return this.ctrlKey
}
func (this *KeyboardManager) ShiftKey() bool {
	return this.shiftKey
}
func (this *KeyboardManager) CommandKey() bool {
	return this.commandKey
}
func (this *KeyboardManager) keyboardCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Repeat {
		return
	}

	switch scancode {
	case 56:
		switch action {
		case 0:
			this.altKey = false
		case 1:
			this.altKey = true
		}
	case 312:
		switch action {
		case 0:
			this.altKey = false
		case 1:
			this.altKey = true
		}
	case 29:
		switch action {
		case 0:
			this.ctrlKey = false
		case 1:
			this.ctrlKey = true
		}
	case 285:
		switch action {
		case 0:
			this.ctrlKey = false
		case 1:
			this.ctrlKey = true
		}
	case 42:
		switch action {
		case 0:
			this.shiftKey = false
		case 1:
			this.shiftKey = true
		}
	case 54:
		switch action {
		case 0:
			this.shiftKey = false
		case 1:
			this.shiftKey = true
		}
	case 91:
		switch action {
		case 0:
			this.commandKey = false
		case 1:
			this.commandKey = true
		}
	case 93:
		switch action {
		case 0:
			this.commandKey = false
		case 1:
			this.commandKey = true
		}
	}

	this.keyCode = uint32(key)
	this.scanCode = uint32(scancode)

	keyboardEvent := new(KeyboardEvent)
	keyboardEvent.altKey = this.altKey
	keyboardEvent.ctrlKey = this.ctrlKey
	keyboardEvent.shiftKey = this.shiftKey
	keyboardEvent.commandKey = this.commandKey
	keyboardEvent.keyCode = this.keyCode
	keyboardEvent.scanCode = this.scanCode
	keyboardEvent.keyLocation = this.keyLocation

	event := new(events.Event)
	event.Data = keyboardEvent
	switch action {
	case 0:
		event.Type = events.KEY_UP
	case 1:
		event.Type = events.KEY_DOWN
	}

	//fmt.Println(keyboardEvent.altKey, keyboardEvent.ctrlKey, keyboardEvent.shiftKey, keyboardEvent.commandKey, keyboardEvent.keyCode, keyboardEvent.scanCode, event.Type)

	this.DispatchEvent(event)
}
