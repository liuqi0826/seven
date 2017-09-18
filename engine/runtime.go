package engine

import (
	"fmt"
	"runtime"
	"time"

	"github.com/vulkan-go/glfw/v3.3/glfw"

	"github.com/liuqi0826/seven/engine/display"
	"github.com/liuqi0826/seven/engine/display/platform/es"
	"github.com/liuqi0826/seven/engine/resource"
	//"github.com/liuqi0826/seven/engine/display/platform/vulkan"
	"github.com/liuqi0826/seven/engine/global"
	"github.com/liuqi0826/seven/engine/utils"
	"github.com/liuqi0826/seven/events"
)

var Instance *Runtime

func init() {
	runtime.LockOSThread()

	Instance = new(Runtime)
	Instance.Runtime()
}

func check(title string, err error) {
	if err != nil {
		fmt.Println(title, err)
	}
}

//++++++++++++++++++++ Runtime ++++++++++++++++++++

type Runtime struct {
	events.EventDispatcher

	Alive bool
	Ready bool

	ViewPort display.Viewport

	config        *utils.Config
	instanceIndex uint32

	action func()
}

func (this *Runtime) Runtime() {
	this.instanceIndex = 0
}
func (this *Runtime) Setup(config *utils.Config) error {
	var err error
	this.config = config

	global.Debug = this.config.Debug
	global.API = this.config.API
	global.WindowTitle = this.config.WindowTitle
	global.WindowWidth = this.config.WindowWidth
	global.WindowHeight = this.config.WindowHeight
	global.WindowX = this.config.WindowX
	global.WindowY = this.config.WindowY

	err = glfw.Init()
	check("glfw", err)

	switch this.config.API {
	case global.GLES:
		global.Context3D = new(es.Context)
	case global.VULKAN:
	case global.D3D9:
	case global.D3D12:
	}

	err = global.Context3D.Setup(this.config)
	check("context", err)

	global.ResourceManager = new(resource.ResourceManager)

	err = global.ResourceManager.Setup(global.Context3D)
	check("resource", err)

	this.Ready = true

	event := new(events.Event)
	event.Type = events.INIT
	this.DispatchEvent(event)
	return err
}
func (this *Runtime) GetNextInstanceID() string {
	this.instanceIndex++
	return fmt.Sprintf("%d", this.instanceIndex)
}
func (this *Runtime) Start() {
	if this.Ready {
		this.Alive = true
		if this.config.FrameInterval == 0 {
			this.config.FrameInterval = global.FPS60
		}
		this.ViewPort = display.Viewport{}
		this.ViewPort.Viewport(uint32(this.config.WindowWidth), uint32(this.config.WindowHeight), global.FORWARD)
		for this.Alive {
			if global.Context3D.ShouldClose() {
				this.Alive = false
			} else {
				this.frame()
				glfw.PollEvents()
			}
		}
	}
}
func (this *Runtime) Stop() {
	this.Alive = false
}
func (this *Runtime) Destroy() {
	glfw.Terminate()
}
func (this *Runtime) StageWidth() int32 {
	return int32(this.config.WindowWidth)
}
func (this *Runtime) StageHeight() int32 {
	return int32(this.config.WindowHeight)
}
func (this *Runtime) SetAction(action func()) {
	this.action = action
}
func (this *Runtime) frame() {
	begin := time.Now().UnixNano()

	//action
	if this.action != nil {
		this.action()
	}

	this.ViewPort.Frame()

	end := time.Now().UnixNano()
	itv := time.Duration(end - begin)
	if itv < this.config.FrameInterval {
		time.Sleep(time.Nanosecond * time.Duration(this.config.FrameInterval-itv))
	}
	//fmt.Println(itv.Nanoseconds())
}
