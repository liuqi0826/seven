package opengl

import (
	"fmt"

	"github.com/liuqi0826/seven/api/khronos/gl/gl"
	"github.com/liuqi0826/seven/engine/display/resource"
)

type Program3D struct {
	Index          uint32
	VertexShader   uint32
	FragmentShader uint32

	usedCount uint32
	uploaded  bool
}

func (this *Program3D) Upload(vertexProgram string, fragmentProgram string) error {
	var err error
	this.VertexShader, err = this.compileShader(vertexProgram, gl.GL_VERTEX_SHADER)
	if err != nil {
		fmt.Println(err)
		return err
	}
	this.FragmentShader, err = this.compileShader(fragmentProgram, gl.GL_FRAGMENT_SHADER)
	if err != nil {
		fmt.Println(err)
		return err
	}

	this.Index = gl.CreateProgram()
	//fmt.Println("GL CreateProgram", this.Index)
	//fmt.Println(this.Index, this.VertexShader, this.FragmentShader)

	gl.AttachShader(this.Index, this.VertexShader)
	gl.AttachShader(this.Index, this.FragmentShader)
	gl.LinkProgram(this.Index)

	var status int32
	gl.GetProgramiv(this.Index, gl.GL_LINK_STATUS, &status)
	if status == gl.GL_FALSE {
		var logLength int32
		gl.GetProgramiv(this.Index, gl.GL_INFO_LOG_LENGTH, &logLength)
		log := gl.GetProgramInfoLog(this.Index, logLength, nil)
		return fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(this.VertexShader)
	gl.DeleteShader(this.FragmentShader)

	if err == nil {
		this.uploaded = true
	}

	return err
}
func (this *Program3D) Dispose() {
	gl.DeleteShader(this.Index)
}
func (this *Program3D) AddCount() {
	this.usedCount++
}
func (this *Program3D) SubCount() {
	this.usedCount--
}
func (this *Program3D) GetCount() uint32 {
	return this.usedCount
}
func (this *Program3D) IsReady() bool {
	return this.uploaded
}
func (this *Program3D) compileShader(source string, shaderType int32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	//fmt.Println("compileShader:", shader, gl.GL_VERTEX_SHADER, shaderType)
	gl.ShaderSource(shader, 1, source, nil)
	gl.CompileShader(shader)
	var status int32
	gl.GetShaderiv(shader, gl.GL_COMPILE_STATUS, &status)
	if status == gl.GL_FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.GL_INFO_LOG_LENGTH, &logLength)
		log := gl.GetShaderInfoLog(shader, logLength, nil)
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
	#version 450
	layout (std140) uniform MatrixBlock {
		mat4 projection;
		mat4 camera;
		mat4 transform;
	};
	in vec3 position;
	in vec2 texcoord;
	in vec3 normal;
	out vec2 vtc;
	out vec3 vn;
	void main() {
		vtc = texcoord;
		vn = normal;
		gl_Position = projection * camera * transform * vec4(position, 1);
	}`
	ShaderResource["default"].Fragment = `
	#version 450
	layout (binding = 0) uniform sampler2D tex;
	in vec2 vtc;
	in vec3 vn;
	vec4 color;
	void main() {
		color = texture(tex, vtc);
		gl_FragColor = color;
	}`

	ShaderResource["stylize"] = new(resource.ShaderResource)
	ShaderResource["stylize"].ID = "stylize"
	ShaderResource["stylize"].Vertex = `
	#version 450
	layout (std140) uniform MatrixBlock {
		mat4 projection;
		mat4 camera;
		mat4 transform;
	};
	in vec3 position;
	in vec2 texcoord;
	in vec3 normal;
	out vec2 vtc;
	out vec3 vn;
	void main() {
		vtc = texcoord;
		vn = normal;
		gl_Position = projection * camera * transform * vec4(position, 1);
	}`
	ShaderResource["stylize"].Fragment = `
	#version 450
	layout (binding = 0) uniform sampler2D tex;
	in vec2 vtc;
	in vec3 vn;
	vec4 color;
	void main() {
		color = texture(tex, vtc);
		gl_FragColor = color;
	}`

	// ShaderResource["skeleton_matrix"] = new(resource.ShaderResource)
	// ShaderResource["skeleton_matrix"].ID = "skeleton_matrix"
	// ShaderResource["skeleton_matrix"].Vertex = `
	// #version 450
	// layout (std140) uniform MatrixBlock {
	// 	mat4 projection;
	// 	mat4 camera;
	// 	mat4 transform;
	// };
	// layout (std140) uniform JointBlock {
	// 	mat4 joint[100];
	// };
	// in vec3 position;
	// in vec2 texcoord;
	// in vec3 normal;
	// out vec2 vtc;
	// out vec3 vn;
	// void main() {
	// 	vtc = texcoord;
	// 	vn = normal;
	// 	gl_Position = projection * camera * transform * vec4(position, 1.0);
	// }`
	// ShaderResource["skeleton_matrix"].Fragment = `
	// #version 450
	// layout (binding = 0) uniform sampler2D tex;
	// in vec2 vtc;
	// in vec3 vn;
	// vec4 color;
	// void main() {
	// 	color = texture(tex, vtc);
	// 	//color.rgb = color.rgb + vn.rgb;
	// 	gl_FragColor = color;
	// }`
}
