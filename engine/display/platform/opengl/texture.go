package opengl

import (
	"bytes"
	"errors"
	"image"
	"image/draw"
	_ "image/png"

	"github.com/liuqi0826/seven/api/khronos/gl/gl"
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
	gl.BindTexture(gl.GL_TEXTURE_2D, this.Index)
	gl.TexParameteri(gl.GL_TEXTURE_2D, gl.GL_TEXTURE_MIN_FILTER, gl.GL_LINEAR)
	gl.TexParameteri(gl.GL_TEXTURE_2D, gl.GL_TEXTURE_MAG_FILTER, gl.GL_LINEAR)
	gl.TexParameteri(gl.GL_TEXTURE_2D, gl.GL_TEXTURE_WRAP_S, gl.GL_CLAMP_TO_EDGE)
	gl.TexParameteri(gl.GL_TEXTURE_2D, gl.GL_TEXTURE_WRAP_T, gl.GL_CLAMP_TO_EDGE)
	gl.TexImage2D(gl.GL_TEXTURE_2D, 0, gl.GL_RGBA, int32(1024), int32(1024), 0, gl.GL_RGBA, gl.GL_UNSIGNED_BYTE, rgba.Pix)
	return err
}
func (this *Texture) UploadCompressedTexture(source []byte) error {
	var err error
	return err
}
func (this *Texture) SetSlot(index int32) {
	var solt int32
	switch index {
	case 0:
		solt = gl.GL_TEXTURE0
	case 1:
		solt = gl.GL_TEXTURE1
	case 2:
		solt = gl.GL_TEXTURE2
	case 3:
		solt = gl.GL_TEXTURE3
	case 4:
		solt = gl.GL_TEXTURE4
	case 5:
		solt = gl.GL_TEXTURE5
	case 6:
		solt = gl.GL_TEXTURE6
	case 7:
		solt = gl.GL_TEXTURE7
	}
	gl.ActiveTexture(solt)
	gl.BindTexture(gl.GL_TEXTURE_2D, this.Index)
}
func (this *Texture) Dispose() {

}
