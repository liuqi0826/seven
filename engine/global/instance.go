package global

import (
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/resource"
	"github.com/liuqi0826/seven/events"
)

var Context3D core.IContext

var Dispatch *events.EventDispatcher

var ResourceManager *resource.ResourceManager
