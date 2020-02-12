package program

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type (
	// Service provides methods to compile and execute programs.
	Service interface {
		// Compile compiles a source code file to a binary file.
		Compile(ctx context.Context, sourceRelativePath, binaryRelativePath string) error

		// GetExecutionCommand returns an *exec.Cmd to execute the given program.
		GetExecutionCommand(ctx context.Context, sourceRelativePath, binaryRelativePath string) *exec.Cmd

		// GetSourceFileExtension returns the extension for source code files names.
		GetSourceFileExtension() string

		// GetBinaryFileExtension returns the extension for binary executable file names.
		GetBinaryFileExtension() string
	}

	programServiceRuntime interface {
		Run(cmd *exec.Cmd) error
	}

	programServiceDefaultRuntime struct {
	}

	// cpp11ProgramService implements compilation and execution for C++ 11.
	cpp11ProgramService struct {
		runtime programServiceRuntime
	}
)

// NewService creates a Service according to the given key.
// If nil is passed, the Service will be created with the default programServiceRuntime.
func NewService(programServiceKey string, runtime programServiceRuntime) (Service, error) {
	if runtime == nil {
		runtime = &programServiceDefaultRuntime{}
	}
	svc, ok := getServices(runtime)[programServiceKey]
	if !ok {
		return nil, fmt.Errorf(
			"invalid program service, want one in {%s}, got '%s'",
			strings.Join(GetServices(), ", "),
			programServiceKey,
		)
	}
	return svc, nil
}

// GetServices returns a string list of the available program services.
func GetServices() []string {
	programServices := getServices(nil)
	strings := make([]string, 0, len(programServices))
	for key := range programServices {
		strings = append(strings, "'"+key+"'")
	}
	return strings
}

// getServices returns the available program services.
func getServices(runtime programServiceRuntime) map[string]Service {
	return map[string]Service{
		"c++11": &cpp11ProgramService{runtime: runtime},
	}
}

func (*programServiceDefaultRuntime) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}

func (p *cpp11ProgramService) Compile(ctx context.Context, sourceRelativePath, binaryRelativePath string) error {
	cmd := exec.CommandContext(ctx, "g++", "-std=c++11", sourceRelativePath, "-o", binaryRelativePath)
	if err := p.runtime.Run(cmd); err != nil {
		return fmt.Errorf("error compiling for c++11: %w", err)
	}
	return nil
}

func (*cpp11ProgramService) GetExecutionCommand(ctx context.Context, sourceRelativePath, binaryRelativePath string) *exec.Cmd {
	return exec.CommandContext(ctx, binaryRelativePath)
}

func (*cpp11ProgramService) GetSourceFileExtension() string {
	return ".cpp"
}

func (*cpp11ProgramService) GetBinaryFileExtension() string {
	return ""
}
