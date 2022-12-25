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

func init() {
	flag.StringVar(&cfg.Path, "input", "", "input file")
	flag.StringVar(&cfg.Path, "i", "", "input file")

	flag.StringVar(&cfg.Output, "output", "", "output file path (.png)")
	flag.StringVar(&cfg.Output, "o", "", "output file path (.png)")

	flag.IntVar(&cfg.Width, "width", 800, "width of output")
	flag.IntVar(&cfg.Width, "w", 800, "width of output")

	flag.IntVar(&cfg.Height, "height", 800, "height of output")
	flag.IntVar(&cfg.Height, "h", 800, "height of output")

	flag.StringVar(&cfg.Prop, "property", "", "property to use from input")
	flag.StringVar(&cfg.Prop, "p", "", "property to use from input")

	flag.StringVar(&cfg.Stat, "stat", "density", "function for displaying values, can be one of: density, mean")
	flag.StringVar(&cfg.Stat, "s", "density", "function for displaying values, can be one of: density, mean")

	flag.StringVar(&cfg.Bbox, "bbox", "", "string delimited bounding box in order: xmax,xmin,ymax,ymin")
	flag.StringVar(&cfg.Bbox, "b", "", "string delimited bounding box in order: xmax,xmin,ymax,ymin")

	flag.StringVar(&cfg.Prof, "prof", "", "write cpu profile to file")
}

func main() {
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
