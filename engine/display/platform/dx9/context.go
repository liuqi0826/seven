package dx9

import (
	//"github.com/gonutz/d3d9"
	//"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/engine/utils"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type Context struct {
	config *utils.Config
	window *glfw.Window

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
