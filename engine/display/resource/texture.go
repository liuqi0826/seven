package resource

import ()

type TextureResource struct {
	ID      string
	Type    string
	Texture []byte
}

func (this *TextureResource) Parser(value []byte) error {
	var err error
	return err
}
