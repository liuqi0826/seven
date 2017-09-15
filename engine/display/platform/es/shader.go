package es

import (
	"fmt"
	"strings"

	gl "github.com/go-gl/gl/v3.1/gles2"
)

type Program3D struct {
	Index          uint32
	VertexShader   uint32
	fragmentShader uint32
}

func (this *Program3D) Upload(vertexProgram string, fragmentProgram string) error {
	var err error
	this.VertexShader, err = this.compileShader(vertexProgram, gl.VERTEX_SHADER)
	if err != nil {
		return err
	}
	this.fragmentShader, err = this.compileShader(fragmentProgram, gl.FRAGMENT_SHADER)
	if err != nil {
		return err
	}

	this.Index = gl.CreateProgram()

	gl.AttachShader(this.Index, this.VertexShader)
	gl.AttachShader(this.Index, this.fragmentShader)
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
	gl.DeleteShader(this.fragmentShader)

	return nil
}
func (this *Program3D) Dispose() {
	gl.DeleteShader(this.Index)
}
func (this *Program3D) compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

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

var ShaderDefaultVertex string = `
	#version 300 es
	
	in vec3 position;
	
	void main() {
	    gl_Position = vec4(position, 1);
	}
` + "\x00"

var ShaderDefaultFragment string = `
	#version 300 es

	void main() {
	    gl_FragColor = vec4(1.0, 1.0, 1.0, 1.0);
	}
` + "\x00"
