package render

import (
	"github.com/liuqi0826/seven/engine/display/core"
	"github.com/liuqi0826/seven/engine/display/platform/opengl"
	"github.com/liuqi0826/seven/engine/static"
)

func CreateRender(id string) core.IRenderer {
	switch static.API {
	case static.GL:
		switch id {
		case "default":
			renderer := new(opengl.DefaultRender)
			return renderer
		}
	case static.VULKAN:
	case static.D3D9:
	case static.D3D12:
	}
	return nil
}
