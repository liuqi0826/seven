package static

import (
	"time"
)

const (
	FPS30 = time.Duration(33333)
	FPS60 = time.Duration(16666)
)

const (
	GLES   = "gles"
	VULKAN = "vulkan"
	D3D9   = "d3d9"
	D3D12  = "d3d12"
)

const (
	ALWAYS        = "always"
	EQUAL         = "equal"
	GREATER       = "greater"
	GREATER_EQUAL = "greaterEqual"
	LESS          = "less"
	LESS_EQUAL    = "lessEqual"
	NEVER         = "never"
	NOT_EQUAL     = "notEqual"
)

const (
	BYTES_4 = "bytes4"
	FLOAT_1 = "float1"
	FLOAT_2 = "float2"
	FLOAT_3 = "float3"
	FLOAT_4 = "float4"
)

const (
	FORWARD  = "forward"
	DEFERRED = "deferred"
)
