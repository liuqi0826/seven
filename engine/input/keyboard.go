package input

import (
	"fmt"

	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type KeyboardManager struct {
	window *glfw.Window
}

func (this *KeyboardManager) Setup(window *glfw.Window) {
	this.window = window
	this.window.SetKeyCallback(this.keyboardCallback)
}
func (this *KeyboardManager) keyboardCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Repeat {
		return
	}
	fmt.Println(key, scancode, action, mods)
}
