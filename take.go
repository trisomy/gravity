package take

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gsdocker/gslogger"
	trisomyos "github.com/trisomy/take/os"
)

// Take the trisomy build tool context object
type Take struct {
	gslogger.Log        // minx log
	path         string // The trisomy build system install path
}

// NewTake create new trisomy build system context object
func NewTake(path string) *Take {
	if path == "" {
		path = os.Getenv("TAKEPATH")
		if path == "" {
			panic("env varible not set : TAKEPATH")
		}
	}

	os.Setenv("GOROOT", filepath.Join(path, "GOROOT"))

	return &Take{
		Log:  gslogger.Get("take"),
		path: path,
	}
}

// Go run embed go command
func (take *Take) Go(GOPATH string, args ...string) error {

	cmdName := filepath.Join(take.path, "GOROOT/bin/go", trisomyos.ExeFileExtension)

	take.I("execute : %s %s", cmdName, args)

	cmd := exec.Command(cmdName, args...)

	cmd.Stdout = os.Stdout

	return cmd.Run()
}
