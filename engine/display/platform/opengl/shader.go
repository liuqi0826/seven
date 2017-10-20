package opengl

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.5-core/gl"

	"github.com/liuqi0826/seven/engine/display/resource"
)

type Program3D struct {
	Index          uint32
	VertexShader   uint32
	FragmentShader uint32
	userCount      uint32
}

func (this *Program3D) Upload(vertexProgram string, fragmentProgram string) error {
	var err error
	this.VertexShader, err = this.compileShader(vertexProgram, gl.VERTEX_SHADER)
	if err != nil {
		fmt.Println(err)
		return err
	}
	this.FragmentShader, err = this.compileShader(fragmentProgram, gl.FRAGMENT_SHADER)
	if err != nil {
		fmt.Println(err)
		return err
	}

	this.Index = gl.CreateProgram()

	gl.AttachShader(this.Index, this.VertexShader)
	gl.AttachShader(this.Index, this.FragmentShader)
	gl.LinkProgram(this.Index)

	var status int32
	gl.GetProgramiv(this.Index, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(this.Index, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(this.Index, logLength, nil, gl.Str(log))

		return fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(this.VertexShader)
	gl.DeleteShader(this.FragmentShader)

	return nil
}
func (this *Program3D) Dispose() {
	gl.DeleteShader(this.Index)
}
func (this *Program3D) AddCount() {
	this.userCount++
}
func (this *Program3D) SubCount() {
	this.userCount--
}
func (this *Program3D) GetCount() uint32 {
	return this.userCount
}
func (this *Program3D) compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	fmt.Println(shader)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

var ShaderResource map[string]*resource.ShaderResource

func init() {
	ShaderResource = make(map[string]*resource.ShaderResource)

	ShaderResource["default"] = new(resource.ShaderResource)
	ShaderResource["default"].ID = "default"
	ShaderResource["default"].Vertex = `
	#version 330
	uniform mat4 projection;
	uniform mat4 transform;
	in vec3 position;
	in vec2 texcoord;
	in vec3 normal;
	out vec2 vtc;
	out vec3 vn;
	void main() {
		vtc = texcoord;
		vn = normal;
		gl_Position = projection * transform * vec4(position, 1.0);
	}
	` + "\x00"
	ShaderResource["default"].Fragment = `
	#version 330
	uniform sampler2D tex;
	in vec2 vtc;
	in vec3 vn;
	vec4 color;
	void main() {
		color = texture(tex, vtc);
		color.rgb = color.rgb + vn.rgb;
		gl_FragColor = color;
	}
	` + "\x00"

}
