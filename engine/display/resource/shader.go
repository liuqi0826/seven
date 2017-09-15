package resource

import ()

type ShaderResource struct {
	ID       string `json:"id"`
	Vertex   string `json:"vertex"`
	Fragment string `json:"fragment"`
}
