package engine

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/vulkan-go/glfw/v3.3/glfw"

	"github.com/liuqi0826/seven/engine/display"
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform/es"
	//"github.com/liuqi0826/seven/engine/display/platform/vulkan"
	"github.com/liuqi0826/seven/engine/utils"
	"github.com/liuqi0826/seven/events"
)

func init() {
	runtime.LockOSThread()
}

func check(title string, err error) {
	if err != nil {
		fmt.Println(title, err)
	}
}

var Instance *Runtime

func Init(config *utils.Config) error {
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
	Context  core.IContext
	Resource ResourceManager

	config        *utils.Config
	instanceIndex uint32

	action func()
	render func()
}

func (this *Runtime) Runtime() {
	this.instanceIndex = 0
	this.Resource = ResourceManager{}
	this.Resource.ResourceManager()
}
func (this *Runtime) Setup(config *utils.Config) error {
	var err error
	this.config = config

	err = glfw.Init()
	check("glfw", err)

	switch this.config.API {
	case GLES:
		this.Context = new(es.Context)
	case VULKAN:
	case D3D9:
	case D3D12:
	}

	err = this.Context.Setup(this.config)
	check("context", err)

	err = this.Resource.Setup(this.Context)
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
			if this.Context.ShouldClose() {
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
func (this *Runtime) SetRender(render func()) {
	this.render = render
}
func (this *Runtime) defaultAction() {
}
func (this *Runtime) defaultRender() {
	this.ViewPort.Frame()
}
func (this *Runtime) frame() {
	begin := time.Now().UnixNano()

	//action
	if this.action != nil {
		this.action()
	}

	//render
	if this.render != nil {
		this.render()
	}

	this.Context.Present()

	end := time.Now().UnixNano()
	itv := time.Duration(end - begin)
	if itv < this.config.FrameInterval {
		time.Sleep(time.Nanosecond * time.Duration(this.config.FrameInterval-itv))
	}
	//fmt.Println(itv.Nanoseconds())
}
