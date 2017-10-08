package resource

import (
	"bytes"
	"encoding/gob"
)

type MaterialResource struct {
	ID          string
	TextureList [8]*TextureResource
}

func (this *MaterialResource) Parser(value []byte) error {
	var err error
	reader := bytes.NewBuffer(value)
	decoder := gob.NewDecoder(reader)
	err = decoder.Decode(this)
	return err
}
