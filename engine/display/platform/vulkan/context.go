package vulkan

import (
	"github.com/liuqi0826/seven/engine/utils"

	"github.com/vulkan-go/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"
)

type Context struct {
	config       *utils.Config
	window       *glfw.Window
	windowHandle uintptr

	device vulkan.Device

	ready bool
	debug bool

	clearRed   float32
	clearGreen float32
	clearBlue  float32
	clearAlpha float32

	depthMask       bool
	passCompareMode string
}

func (this *Context) Setup(config *utils.Config) error {
	var err error

	this.config = config
	this.debug = config.Debug

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	this.window, err = glfw.CreateWindow(this.config.WindowWidth, this.config.WindowHeight, this.config.WindowTitle, nil, nil)
	if err != nil {
		return err
	}
	this.window.MakeContextCurrent()

	this.ready = true

	return err
}
func (this *Context) GetWindow() *glfw.Window {
	return this.window
}
