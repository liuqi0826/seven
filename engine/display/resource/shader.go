package resource

import (
	"encoding/json"
)

type ShaderResource struct {
	ID       string `json:"id"`
	Vertex   string `json:"vertex"`
	Fragment string `json:"fragment"`
}

func (this *ShaderResource) Parser(value []byte) error {
	var err error
	err = json.Unmarshal(value, this)
	return err
}
