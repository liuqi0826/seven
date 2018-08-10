package engine

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/vulkan-go/glfw/v3.3/glfw"

	"github.com/liuqi0826/seven/engine/display"
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform/opengl"
	"github.com/liuqi0826/seven/engine/display/platform/vulkan"
	"github.com/liuqi0826/seven/engine/global"
	"github.com/liuqi0826/seven/engine/input"
	"github.com/liuqi0826/seven/engine/resource"
	"github.com/liuqi0826/seven/engine/static"
	"github.com/liuqi0826/seven/engine/utils"
	"github.com/liuqi0826/seven/events"
)

var Runtime *Engine

func init() {
	runtime.LockOSThread()

	Runtime = new(Engine)
	Runtime.Engine()
}

//++++++++++++++++++++ Runtime ++++++++++++++++++++

type Engine struct {
	events.EventDispatcher
	sync.Mutex

	Alive bool
	Ready bool

	Window          *glfw.Window
	Stage           *Stage
	KeyboardManager *input.KeyboardManager
	MouseManager    *input.MouseManager

	config        *utils.Config
	instanceIndex uint32

	actionList    []func()
	actionChannel chan func()
}

func (this *Engine) Engine() {
	this.instanceIndex = 0
	this.actionList = make([]func(), 0)
	this.actionChannel = make(chan func(), 256)
}
func (this *Engine) Setup(config *utils.Config) error {
	var err error
	this.config = config

	static.Debug = this.config.Debug
	static.API = this.config.API
	static.WindowTitle = this.config.WindowTitle
	static.WindowWidth = this.config.WindowWidth
	static.WindowHeight = this.config.WindowHeight
	static.WindowX = this.config.WindowX
	static.WindowY = this.config.WindowY

	err = glfw.Init()
	if err != nil {
		return err
	}

	//glfw.WindowHint(glfw.Resizable, glfw.False)
	//glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	//glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	this.Window, err = glfw.CreateWindow(this.config.WindowWidth, this.config.WindowHeight, this.config.WindowTitle, nil, nil)
	if err != nil {
		return err
	}

	this.Window.SetSizeCallback(this.resizeCallback)

	this.Window.MakeContextCurrent()

	switch this.config.API {
	case static.GL:
		global.Context3D = new(opengl.Context)
	case static.VULKAN:
		global.Context3D = new(vulkan.Context)
	case static.D3D9:
	case static.D3D12:
	}

	this.KeyboardManager = new(input.KeyboardManager)
	this.KeyboardManager.Setup(this.Window)
	this.MouseManager = new(input.MouseManager)
	this.MouseManager.Setup(this.Window)

	err = global.Context3D.Setup(this.Window, this.config.Debug)
	utils.Check("context", err)

	global.ResourceManager = new(resource.ResourceManager)
	global.ResourceManager.Setup(global.Context3D)
	global.ResourceManager.AddEventListener(static.RESOURCE_EVENT, this.onResourceEvent)

	this.Ready = true

	event := new(events.Event)
	event.Type = events.INIT
	this.DispatchEvent(event)
	return err
}
func (this *Engine) GetNextInstanceID() string {
	this.instanceIndex++
	return fmt.Sprintf("%d", this.instanceIndex)
}
func (this *Engine) Start(function func()) {
	if this.Ready {
		this.Alive = true
		if this.config.FrameInterval == 0 {
			this.config.FrameInterval = static.FPS60
		}

		this.Stage = new(Stage)
		this.Stage.setup(this)

		go this.actionListion()

		if function != nil {
			go function()
		}

		for this.Alive {
			this.frame()
		}
	}
}
func (this *Engine) Stop() {
	this.Alive = false
}
func (this *Engine) AddActionOSThread(function func()) {
	this.actionChannel <- function
}
func (this *Engine) Destroy() {
	glfw.Terminate()
}

func (this *Engine) onResourceEvent(event *events.Event) {
	this.AddActionOSThread(global.ResourceManager.Upload)
}
func (this *Engine) actionListion() {
	for act := range this.actionChannel {
		this.Lock()
		this.actionList = append(this.actionList, act)
		this.Unlock()
	}
}
func (this *Engine) action() {
	this.Lock()
	defer this.Unlock()
	for _, fun := range this.actionList {
		if fun != nil {
			fun()
		}
	}
	this.actionList = make([]func(), 0)
}
func (this *Engine) frame() {
	if this.Window.ShouldClose() {
		this.Alive = false
	}

	begin := time.Now().UnixNano()

	this.action()
	this.Stage.frame()

	end := time.Now().UnixNano()
	itv := time.Duration(end - begin)
	if itv < this.config.FrameInterval {
		time.Sleep(time.Nanosecond * time.Duration(this.config.FrameInterval-itv))
	}

	glfw.PollEvents()
}

func (this *Engine) frameBufferSizeCallback() {
}
func (this *Engine) scrollCallback() {
}
func (this *Engine) makeContextCurrentCallback() {
}
func (this *Engine) resizeCallback(window *glfw.Window, width int, height int) {
	fmt.Println(width, height)
}

//++++++++++++++++++++ Stage ++++++++++++++++++++

type Stage struct {
	events.EventDispatcher

	engine *Engine
	view   display.Viewport
}

func (this *Stage) AddChild(displayObject core.IDisplayObject) {
	this.view.Scene.AddChild(displayObject)
}
func (this *Stage) RemoveChild(displayObject core.IDisplayObject) core.IDisplayObject {
	return this.view.Scene.RemoveChild(displayObject)
}
func (this *Stage) StageWidth() int32 {
	return int32(this.engine.config.WindowWidth)
}
func (this *Stage) StageHeight() int32 {
	return int32(this.engine.config.WindowHeight)
}
func (this *Stage) setup(engine *Engine) {
	this.EventDispatcher.EventDispatcher(this)
	this.engine = engine
	this.view = display.Viewport{}
	this.view.Viewport(uint32(this.engine.config.WindowWidth), uint32(this.engine.config.WindowHeight), static.FORWARD)

	event := new(events.Event)
	event.Type = events.INIT
	this.DispatchEvent(event)
}
func (this *Stage) frame() {
	event := new(events.Event)
	event.Type = events.ENTER_FRAME
	this.DispatchEvent(event)

	this.view.Frame()
}
