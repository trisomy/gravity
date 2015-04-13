package gravity

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gsdocker/gserrors"
	"github.com/gsdocker/gslogger"
	"github.com/trisomy/gravity/files"
)

// Compiler errors
var (
	ErrBuilder = errors.New("project builder error")
)

// ProjectBuilder The project builder object
type ProjectBuilder struct {
	gslogger.Log `xml:"-"` // Mixin logger
	ID           string    // The project's uuid
	Source       string    // The project source path
	Target       string    // The project output path
	script       string    // The script
	xmlfile      string    // the project builder serialize xml file
}

// NewProjectBuilder create project builder
func NewProjectBuilder(source string, target string) (builder *ProjectBuilder, err error) {

	source, err = filepath.Abs(source)

	if err != nil {
		return
	}

	target, err = filepath.Abs(target)

	if err != nil {
		return
	}

	builder = &ProjectBuilder{
		Log:     gslogger.Get("Gravity"),
		Source:  source,
		Target:  target,
		xmlfile: filepath.Join(target, "gravity.xml"),
		script:  filepath.Join(source, "project.gr"),
	}

	builder.D(
		`
 ================================================
 project settings
 source :%s
 target :%s
 script :%s
 ================================================
		`,
		builder.Source, builder.Target, builder.script)

	// check project script file
	err = builder.checkProjectScript()

	if err != nil {
		return
	}

	err = builder.createProjectTargetDir()

	return
}

func (builder *ProjectBuilder) createProjectTargetDir() error {
	if !files.Exists(builder.Target) {
		return os.MkdirAll(builder.Target, 0755)
	}

	return nil
}

func (builder *ProjectBuilder) checkProjectScript() error {

	if !files.Exists(builder.script) {
		return gserrors.Newf(ErrBuilder, "project script not exists : %s", builder.script)
	}

	return nil
}

// CreateBuilder create builder
func (builder *ProjectBuilder) CreateBuilder() error {

	return nil
}
