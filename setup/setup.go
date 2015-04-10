package main

import (
	"flag"
	"path/filepath"

	"github.com/gsdocker/gslogger"
)

func main() {
	flag.Parse()

	defer gslogger.Join()

	log := gslogger.Get("setup")
	path, _ := filepath.Abs("./")
	log.I("%s", path)
}
