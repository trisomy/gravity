package main

import (
	"flag"
	"path/filepath"

	"github.com/gsdocker/gslogger"
)

var ip = flag.String("backend", "cmake", "the trismoy build tool backend engines : cmake")

func main() {

	defer gslogger.Join()

	log := gslogger.Get("take")

	log.I("start the trismoy build tool v0.0.1")

	flag.Parse()

	//var projectPath string

	projectPath, err := filepath.Abs("./hello")

	if err != nil {
		log.E("get trismoy execute path err :%s", err)
		return
	}

	projectPath = filepath.Dir(projectPath)

	switch len(flag.Args()) {
	case 0:
	case 1:
		{
			projectPath, err = filepath.Abs(flag.Arg(0))

			if err != nil {
				log.E("get filepath %s abs :%s", flag.Arg(0), err)
				return
			}
		}
	default:
		{
			log.E("invalid options")
			flag.PrintDefaults()
			return
		}
	}

	log.I("project path : %s", projectPath)
}
