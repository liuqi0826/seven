package events

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
	MESSAGE     = "message"

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

	KEY_DOWN  = "keyDown"
	KEY_UP    = "keyUp"
	KEY_PRESS = "keyPress"

	ERROR    = "error"
	IO_ERROR = "ioError"
)

type Event struct {
	Type   string
	Target interface{}
	Data   interface{}
}

type MouseEventData struct {
	X         float32
	Y         float32
	MovementX float32
	MovementY float32
	AltKey    bool
	CtrlKey   bool
	ShiftKey  bool
	Delta     int32
}

type KeyboardEventData struct {
	CharCode uint32
	KeyCode  uint32
	Key      string
	AltKey   bool
	CtrlKey  bool
	ShiftKey bool
}
