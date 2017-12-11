package resource

import (
	"encoding/json"
)

type AnimationClipResource struct {
	ID          string              `json:"id"`
	Duration    int32               `json:"duration"`
	UploadIndex []uint32            `json:"uploadIndex"`
	Channel     []*AnimationChannel `json:"channel"`
}
type AnimationChannel struct {
	Transform []float32 `json:"transform"`
	Time      []int32   `json:"time"`
}

func (this *AnimationClipResource) Parser(value []byte) error {
	var err error
	err = json.Unmarshal(value, this)
	return err
}
