package shader

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Load the shaders, returning the ID of the resulting program.  Any problems
// compiling or linking will result in an error.
func Load(shaders *[]Info) (uint32, error) {
	return load(shaders, false)
}

// LoadSeparable is the same as Load with the exception that before the link stage
// GL_PROGRAM_SEPARABLE is set to GL_TRUE.
func LoadSeparable(shaders *[]Info) (uint32, error) {
	return load(shaders, true)
}

// load the shaders
func load(shaders *[]Info, separable bool) (uint32, error) {
	program := gl.CreateProgram()

	for _, s := range *shaders {
		if err := s.Compile(program); err != nil {
			cleanup(shaders)
			gl.DeleteProgram(program)
			return 0, err
		}
	}

	if separable {
		gl.ProgramParameteri(program, gl.PROGRAM_SEPARABLE, gl.TRUE)
	}

	gl.LinkProgram(program)
	cleanup(shaders)
	var linked int32
	if gl.GetProgramiv(program, gl.LINK_STATUS, &linked); linked == gl.FALSE {
		gl.DeleteProgram(program)
		return 0, fmt.Errorf("failed to link program: %s", getErrorMsg(false, program))
	}

	return program, nil
}

// cleanup all shaders by calling Delete on any non-zero shader in the slice.
func cleanup(shaders *[]Info) {
	for _, s := range *shaders {
		if s.shader != 0 {
			s.Delete()
		}
	}
}