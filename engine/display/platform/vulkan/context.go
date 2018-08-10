package vulkan

import (
//	"fmt"

	"github.com/liuqi0826/seven/engine/display/platform"
	"github.com/liuqi0826/seven/events"

	"github.com/vulkan-go/glfw/v3.3/glfw"
//	"github.com/vulkan-go/vulkan"
)

type Context struct {
	events.EventDispatcher

//	window       *glfw.Window
//	windowHandle uintptr
//
//	applicationInfo *vulkan.ApplicationInfo
//
//	deviceInfo    DeviceInfo
//	swapchainInfo SwapchainInfo
	//renderInfo      VulkanRenderInfo
	//bufferInfo      VulkanBufferInfo
	//gfxPipelineInfo VulkanGfxPipelineInfo

	ready       bool
	enableDebug bool

	clearRed   float32
	clearGreen float32
	clearBlue  float32
	clearAlpha float32

	depthMask       bool
	passCompareMode string
}

func (this *Context) Setup(window *glfw.Window, enableDebug bool) error {
	var err error

	this.EventDispatcher.EventDispatcher(this)

//	this.window = window
//	this.enableDebug = enableDebug
//
//	err = vulkan.Init()
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//
//	this.applicationInfo = &vulkan.ApplicationInfo{
//		SType:              vulkan.StructureTypeApplicationInfo,
//		ApiVersion:         vulkan.MakeVersion(1, 0, 0),
//		ApplicationVersion: vulkan.MakeVersion(1, 0, 0),
//		PApplicationName:   "SevenEngine\x00",
//		PEngineName:        "SevenEngine\x00",
//	}
//
//	err = this.createVulkanDevice()
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//
//	this.swapchainInfo, err = this.deviceInfo.CreateSwapchain()
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//	//VulkanInit(&this.deviceInfo, &this.swapchainInfo, &this.renderInfo, &this.bufferInfo, &this.gfxPipelineInfo)
//
//	this.ready = true

	return err
}
func (this *Context) Clear(color bool, depth bool, stencil bool) {
}
func (this *Context) ConfigureBackBuffer() {
}
func (this *Context) CreateCubeTexture() {
}
func (this *Context) CreateProgram() platform.IProgram3D {
	return nil
}
func (this *Context) CreateTexture() {
}
func (this *Context) CreateIndexBuffer() platform.IIndexBuffer {
	return nil
}
func (this *Context) CreateVertexBuffer() platform.IVertexBuffer {
	return nil
}
func (this *Context) Dispose() {
	if this == nil {
		return
	}
}
func (this *Context) Present() {
//	this.window.SwapBuffers()
}
func (this *Context) DrawTriangles(indexBuffer platform.IIndexBuffer, firstIndex int32, numTriangles int32) {
}
func (this *Context) SetBlendFactors() {
}
func (this *Context) SetClearColor(red float32, green float32, blue float32, alpha float32) {
}
func (this *Context) SetColorMask() {
}
func (this *Context) SetCulling() {
}
func (this *Context) SetDepthTest(depthMask bool, passCompareMode string) {
}
func (this *Context) SetProgram(program platform.IProgram3D) {
}
func (this *Context) SetProgramConstantsFromByteArray() {
}
func (this *Context) SetProgramConstantsFromMatrix() {
}
func (this *Context) SetProgramConstantsFromVector() {
}
func (this *Context) SetRenderToBackBuffer() {
}
func (this *Context) SetRenderToTexture() {
}
func (this *Context) SetScissorRectangle() {
}
func (this *Context) SetStencilActions() {
}
func (this *Context) SetStencilReferenceValue() {
}
func (this *Context) SetTextureAt() {
}
func (this *Context) SetVertexBufferAt(value string, stride int32, bufferOffset int32, format string) {
}

//
func (this *Context) createVulkanDevice() error {
	var err error

//	this.deviceInfo = DeviceInfo{}
//
//	// Phase 1: vulkan.CreateInstance with vulkan.InstanceCreateInfo
//
//	existingExtensions := getInstanceExtensions()
//	fmt.Println("[INFO] Instance extensions:", existingExtensions)
//
//	instanceExtensions := vulkan.GetRequiredInstanceExtensions()
//	if this.enableDebug {
//		instanceExtensions = append(instanceExtensions, "VK_EXT_debug_report\x00")
//	}
//
//	// ANDROID:
//	// these layers must be included in APK,
//	// see Android.mk and ValidationLayers.mk
//	instanceLayers := []string{
//		// "VK_LAYER_GOOGLE_threading\x00",
//		// "VK_LAYER_LUNARG_parameter_validation\x00",
//		// "VK_LAYER_LUNARG_object_tracker\x00",
//		// "VK_LAYER_LUNARG_core_validation\x00",
//		// "VK_LAYER_LUNARG_api_dump\x00",
//		// "VK_LAYER_LUNARG_image\x00",
//		// "VK_LAYER_LUNARG_swapchain\x00",
//		// "VK_LAYER_GOOGLE_unique_objects\x00",
//	}
//
//	instanceCreateInfo := vulkan.InstanceCreateInfo{
//		SType:                   vulkan.StructureTypeInstanceCreateInfo,
//		PApplicationInfo:        this.applicationInfo,
//		EnabledExtensionCount:   uint32(len(instanceExtensions)),
//		PpEnabledExtensionNames: instanceExtensions,
//		EnabledLayerCount:       uint32(len(instanceLayers)),
//		PpEnabledLayerNames:     instanceLayers,
//	}
//	var v DeviceInfo
//	err = vulkan.Error(vulkan.CreateInstance(&instanceCreateInfo, nil, &v.Instance))
//	if err != nil {
//		err = fmt.Errorf("vulkan.CreateInstance failed with %s", err)
//		return err
//	} else {
//		vulkan.InitInstance(v.Instance)
//	}
//
//	// Phase 2: vulkan.CreateAndroidSurface with vulkan.AndroidSurfaceCreateInfo
//
//	err = vulkan.Error(vulkan.CreateWindowSurface(v.Instance, this.window.GLFWWindow(), nil, &v.Surface))
//	if err != nil {
//		vulkan.DestroyInstance(v.Instance, nil)
//		err = fmt.Errorf("vkCreateWindowSurface failed with %s", err)
//		return err
//	}
//	if v.PhysicalDeviceList, err = getPhysicalDevices(v.Instance); err != nil {
//		v.PhysicalDeviceList = nil
//		vulkan.DestroySurface(v.Instance, v.Surface, nil)
//		vulkan.DestroyInstance(v.Instance, nil)
//		return err
//	}
//
//	existingExtensions = getDeviceExtensions(v.PhysicalDeviceList[0])
//	fmt.Println("[INFO] Device extensions:", existingExtensions)
//
//	// Phase 3: vk.CreateDevice with vk.DeviceCreateInfo (a logical device)
//
//	// ANDROID:
//	// these layers must be included in APK,
//	// see Android.mk and ValidationLayers.mk
//	deviceLayers := []string{
//		// "VK_LAYER_GOOGLE_threading\x00",
//		// "VK_LAYER_LUNARG_parameter_validation\x00",
//		// "VK_LAYER_LUNARG_object_tracker\x00",
//		// "VK_LAYER_LUNARG_core_validation\x00",
//		// "VK_LAYER_LUNARG_api_dump\x00",
//		// "VK_LAYER_LUNARG_image\x00",
//		// "VK_LAYER_LUNARG_swapchain\x00",
//		// "VK_LAYER_GOOGLE_unique_objects\x00",
//	}
//
//	queueCreateInfos := []vulkan.DeviceQueueCreateInfo{{
//		SType:            vulkan.StructureTypeDeviceQueueCreateInfo,
//		QueueCount:       1,
//		PQueuePriorities: []float32{1.0},
//	}}
//	deviceExtensions := []string{
//		"VK_KHR_swapchain\x00",
//	}
//	deviceCreateInfo := vulkan.DeviceCreateInfo{
//		SType:                   vulkan.StructureTypeDeviceCreateInfo,
//		QueueCreateInfoCount:    uint32(len(queueCreateInfos)),
//		PQueueCreateInfos:       queueCreateInfos,
//		EnabledExtensionCount:   uint32(len(deviceExtensions)),
//		PpEnabledExtensionNames: deviceExtensions,
//		EnabledLayerCount:       uint32(len(deviceLayers)),
//		PpEnabledLayerNames:     deviceLayers,
//	}
//	var device vulkan.Device // we choose the first GPU available for this device
//	err = vulkan.Error(vulkan.CreateDevice(v.PhysicalDeviceList[0], &deviceCreateInfo, nil, &device))
//	if err != nil {
//		v.PhysicalDeviceList = nil
//		vulkan.DestroySurface(v.Instance, v.Surface, nil)
//		vulkan.DestroyInstance(v.Instance, nil)
//		err = fmt.Errorf("vulkan.CreateDevice failed with %s", err)
//		return err
//	} else {
//		v.Device = device
//		var queue vulkan.Queue
//		vulkan.GetDeviceQueue(device, 0, 0, &queue)
//		v.Queue = queue
//	}
//
//	if this.enableDebug {
//		// Phase 4: vk.CreateDebugReportCallback
//
//		dbgCreateInfo := vulkan.DebugReportCallbackCreateInfo{
//			SType:       vulkan.StructureTypeDebugReportCallbackCreateInfo,
//			Flags:       vulkan.DebugReportFlags(vulkan.DebugReportErrorBit | vulkan.DebugReportWarningBit),
//			PfnCallback: dbgCallbackFunc,
//		}
//		var dbg vulkan.DebugReportCallback
//		err = vulkan.Error(vulkan.CreateDebugReportCallback(v.Instance, &dbgCreateInfo, nil, &dbg))
//		if err != nil {
//			err = fmt.Errorf("vulkan.CreateDebugReportCallback failed with %s", err)
//			fmt.Println("[WARN]", err)
//			return nil
//		}
//		v.DebugReportCallback = dbg
//	}
//
//	this.deviceInfo = v

	return err
}
