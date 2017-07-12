package engine

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/gonutz/d3d9"
	"github.com/liuqi0826/seven/engine/display"
	"github.com/liuqi0826/seven/events"
	"github.com/veandco/go-sdl2/sdl"
)

const FPS30 = time.Duration(33333)
const FPS60 = time.Duration(16666)

func init() {
	runtime.LockOSThread()
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

	config        *Config
	instanceIndex int32
	window        *sdl.Window
	context       d3d9.Device
}

func (this *Runtime) Runtime() {
	this.instanceIndex = -1
}
func (this *Runtime) Setup(config *Config) error {
	var err error
	this.config = config
	err = sdl.Init(0)
	this.window, err = sdl.CreateWindow(config.WindowTitle, config.WindowX, config.WindowY, config.WindowWidth, config.WindowHeight, 0)
	info, err := this.window.GetWMInfo()
	windowHandle := info.GetWindowsInfo().Window
	err = d3d9.Init()
	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	this.context, _, err = d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		windowHandle,
		d3d9.CREATE_HARDWARE_VERTEXPROCESSING,
		d3d9.PRESENT_PARAMETERS{
			Windowed:      true,
			SwapEffect:    d3d9.SWAPEFFECT_DISCARD,
			HDeviceWindow: windowHandle,
		},
	)
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
		for this.Alive {
			this.frame()
		}
	}
}
func (this *Runtime) Stop() {
	this.Alive = false
}
func (this *Runtime) StageWidth() int32 {
	return int32(this.config.WindowWidth)
}
func (this *Runtime) StageHeight() int32 {
	return int32(this.config.WindowHeight)
}
func (this *Runtime) frame() {
	bigin := time.Now().UnixNano()
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.(type) {
		case *sdl.QuitEvent:
			this.Alive = false
		}
	}
	//render
	this.context.Clear(nil, d3d9.CLEAR_TARGET, 255, 1, 0)
	this.context.BeginScene()
	this.context.EndScene()
	this.context.Present(nil, nil, nil, nil)

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
