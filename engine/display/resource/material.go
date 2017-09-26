package resource

import ()

type MaterialResource struct {
	ID          string
	TextureList []*TextureResource
}

func (this *MaterialResource) Parser(value []byte) error {
	var err error
	return err
}
