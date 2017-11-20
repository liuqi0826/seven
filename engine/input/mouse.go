package input

import (
	"fmt"

	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type MouseManager struct {
	window *glfw.Window
}

func (this *MouseManager) Setup(window *glfw.Window) {
	this.window = window
	this.window.SetMouseButtonCallback(this.mouseButtonCallback)
	this.window.SetCursorPosCallback(this.cursorPositionCallback)
}
func (this *MouseManager) mouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	fmt.Println(button, action, mod)
}
func (this *MouseManager) cursorPositionCallback(window *glfw.Window, x float64, y float64) {
	fmt.Println(x, y)
}
