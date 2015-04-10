package main

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gsdocker/gslogger"
	shutil "github.com/termie/go-shutil"
	"github.com/trisomy/take"
)

func main() {

	flag.Parse()

	var log = gslogger.Get("setup")

	defer func() {
		gslogger.Join()

		if e := recover(); e != nil {
			log.E("exception :%s", e)
		}
	}()

	if len(flag.Args()) != 1 {
		log.E("expect install target path")
		return
	}

	takePath, _ := filepath.Abs("../")
	targetPath, _ := filepath.Abs(flag.Arg(0))

	GOROOT := runtime.GOROOT()

	log.I("GOROOT : %s", GOROOT)
	log.I("src path : %s", takePath)
	log.I("create target path : %s", targetPath)

	err := os.MkdirAll(targetPath, 0755)

	if err != nil {
		log.E("create target path err :%s", err)
		return
	}

	log.I("create target path -- success")

	targetGOROOT := filepath.Join(targetPath, "GOROOT")

	log.I("copy embed golang :\n\tfrom: %s\n\tto: %s", GOROOT, targetGOROOT)

	os.RemoveAll(targetGOROOT)

	shutil.CopyTree(GOROOT, targetGOROOT, nil)

	if err != nil {
		log.E("copy embed golang :%s", err)
		return
	}

	log.I("copy embed golang -- success")

	context := take.NewTake(targetPath)

	project := context.NewProject(takePath)

	err = project.Prepare()

	if err != nil {
		log.E("bootstrap build err :%s", err)
		return
	}
}
