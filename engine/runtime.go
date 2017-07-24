package engine

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/vulkan-go/asche"
	"github.com/vulkan-go/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"

	"github.com/liuqi0826/seven/engine/display"
	"github.com/liuqi0826/seven/events"
)

const FPS30 = time.Duration(33333)
const FPS60 = time.Duration(16666)

func init() {
	runtime.LockOSThread()
}

func check(err interface{}) {
	switch v := err.(type) {
	case error:
		if v != nil {
			panic(err)
		}
	case vulkan.Result:
		if err := vulkan.Error(v); err != nil {
			panic(err)
		}
	case bool:
		if !v {
			panic("condition failed: != true")
		}
	}
}

var Instance *Runtime

func Init(config *Config) error {
	var err error
	if Instance == nil {
		Instance = new(Runtime)
		Instance.Runtime()
		err = Instance.Setup(config)
	} else {
		err = errors.New("Engine instance is initialized.")
	}
	return err
}
func Dispose() error {
	var err error
	return err
}

//++++++++++++++++++++ Runtime ++++++++++++++++++++

type Runtime struct {
	events.EventDispatcher

	Alive bool
	Ready bool

	ViewPort display.Viewport

	Resource ResourceManager

	config        *Config
	instanceIndex int32
	window        *glfw.Window
	windowHandle  uintptr
	context       asche.Context

	action func()
	render func()
}

func (this *Runtime) Runtime() {
	this.instanceIndex = -1
	this.Resource = ResourceManager{}
	this.Resource.ResourceManager()
}
func (this *Runtime) Setup(config *Config) error {
	var err error
	this.config = config

	err = glfw.Init()
	check(err)
	err = vulkan.Init()
	check(err)
	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	this.window, err = glfw.CreateWindow(config.WindowWidth, config.WindowHeight, config.WindowTitle, nil, nil)
	check(err)
	this.windowHandle = this.window.GLFWWindow()
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
			this.config.FrameInterval = FPS60
		}
		this.ViewPort = display.Viewport{}
		this.ViewPort.Viewport(int32(this.config.WindowWidth), int32(this.config.WindowHeight))
		if this.action == nil {
			this.action = this.defaultAction
		}
		if this.render == nil {
			this.render = this.defaultRender
		}
		for this.Alive {
			if this.window.ShouldClose() {
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
func (this *Runtime) SetRender(render func()) {
	this.render = render
}
func (this *Runtime) defaultAction() {
}
func (this *Runtime) defaultRender() {
	this.ViewPort.Frame()
}
func (this *Runtime) frame() {
	bigin := time.Now().UnixNano()

	//action
	if this.action != nil {
		this.action()
	}

	//render
	if this.render != nil {
		this.render()
	}

	end := time.Now().UnixNano()
	itv := time.Duration(end - bigin)
	if itv < this.config.FrameInterval {
		time.Sleep(time.Nanosecond * time.Duration(this.config.FrameInterval-itv))
	}
	fmt.Println(itv.Nanoseconds())
}

//++++++++++++++++++++ Config ++++++++++++++++++++

type Config struct {
	Debug         bool
	WindowTitle   string
	WindowX       int
	WindowY       int
	WindowWidth   int
	WindowHeight  int
	FrameInterval time.Duration
}
