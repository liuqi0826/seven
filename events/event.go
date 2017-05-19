package events

import ()

const (
	ADDED       = "added"
	CANCEL      = "cancel"
	CHANGE      = "change"
	CLEAR       = "clear"
	CLOSE       = "close"
	CLOSEING    = "closeing"
	COMPLETE    = "complete"
	CONNECT     = "connect"
	ENTER_FRAME = "enterFrame"
	EXIT_FRAME  = "exitFrame"
	FULLSCREEN  = "fullscreen"
	INIT        = "init"
	OPEN        = "open"
	PASTE       = "paste"
	REMOVE      = "remove"
	RENDER      = "render"
	RESIZE      = "resize"
	SCROLL      = "scroll"
	SELECT      = "select"

	CLICK             = "click"
	DOUBLE_CLICK      = "doubleClick"
	MIDDLE_CLICK      = "middleClick"
	MIDDLE_MOUSE_DOWN = "middleMouseDown"
	MIDDLE_MOUSE_UP   = "middleMouseUp"
	MOUSE_DOWN        = "mouseDown"
	MOUSE_UP          = "mouseUp"
	MOUSE_MOVE        = "mouseMove"
	MOUSE_WHEEL       = "mouseWheel"
	RIGHT_CLICK       = "rightClick"
	RIGHT_MOUSE_DOWN  = "rightMouseDown"
	RIGHT_MOUSE_UP    = "rightMouseUp"
	ROLL_OUT          = "rollOut"
	ROLL_OVER         = "rollOver"

	ERROR    = "error"
	IO_ERROR = "ioError"
)

type Event struct {
	Type string
	Data interface{}
}
