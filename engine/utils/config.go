package utils

import (
	"time"
)

type Config struct {
	Debug         bool
	API           string
	WindowTitle   string
	WindowX       int
	WindowY       int
	WindowWidth   int
	WindowHeight  int
	FrameInterval time.Duration
}
