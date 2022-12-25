package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"rapture/cli"
	"rapture/pkg/config"
	"runtime/pprof"
)

const usage string = `rapture - visualize large point datasets as rasters

Usage: rapture -i (input) -o (output) [options]

`

var cfg = config.RaptureConfig{}

func main() {
	cfg.Init()
	flag.Parse()

	if cfg.Prof != "" {
		f, err := os.Create(cfg.Prof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if cfg.Path == "" || cfg.Output == "" {
		fmt.Print(usage)
		os.Exit(1)
	} else {
		cli.Run(cfg)
	}
}
