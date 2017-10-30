package opengl

import (
	"bytes"
	"errors"
	"image"
	"image/draw"
	_ "image/png"

	"github.com/go-gl/gl/v4.5-core/gl"
)

const (
	BITMAP = "bitmap"
	ASTC   = "astc"
	ATC    = "atc"
	DXT1   = "dxt1"
	DXT3   = "dxt3"
	DXT5   = "dxt5"
	ETC1   = "etc1"
	ETC2   = "etc2"
	PVRTC  = "pvrtc"
)

type Texture struct {
	Index uint32
	Type  string
}

func (this *Texture) Upload(source []byte, sourceType string) error {
	var err error
	this.Type = sourceType

	//	switch this.Type {
	//	case BITMAP:
	//	case ASTC:
	//	case ATC:
	//	case DXT1:
	//	case DXT3:
	//	case DXT5:
	//	case ETC1:
	//	case ETC2:
	//	case PVRTC:
	//	}

	buff := bytes.NewBuffer(source)
	img, _, err := image.Decode(buff)
	if err != nil {
		return err
	}
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		err = errors.New("Unsupported stride.")
		return err
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	gl.GenTextures(1, &this.Index)
	gl.BindTexture(gl.TEXTURE_2D, this.Index)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(1024), int32(1024), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	return err
}
func (this *Texture) UploadCompressedTexture(source []byte) error {
	var err error
	return err
}
func (this *Texture) SetSlot(index int32) {
	var solt uint32
	switch index {
	case 0:
		solt = gl.TEXTURE0
	case 1:
		solt = gl.TEXTURE1
	case 2:
		solt = gl.TEXTURE2
	case 3:
		solt = gl.TEXTURE3
	case 4:
		solt = gl.TEXTURE4
	case 5:
		solt = gl.TEXTURE5
	case 6:
		solt = gl.TEXTURE6
	case 7:
		solt = gl.TEXTURE7
	}
	gl.ActiveTexture(solt)
	gl.BindTexture(gl.TEXTURE_2D, this.Index)
}
func (this *Texture) Dispose() {

}
