package take

import (
	"os"
	"path/filepath"

	"github.com/gsdocker/gslogger"
)

// Project project settings
type Project struct {
	gslogger.Log        // Mixin log
	take         *Take  // the build system's context object
	trisomyPath  string // the trisomy install path
	path         string
}

// NewProject create new project
func (take *Take) NewProject(path string) *Project {
	return &Project{
		Log:  gslogger.Get("take"),
		take: take,
		path: path,
	}
}

// Prepare prepare build project's builder
func (project *Project) Prepare() error {
	buildPath := filepath.Join(project.path, "build")
	project.I("create project build path : %s", buildPath)

	err := os.Mkdir(buildPath, 0755)
	if err != nil {

		if !os.IsExist(err) {
			project.E("create project build path err :%s", err)
			return err
		}

	}

	project.I("create project build path -- success")

	return nil
}
