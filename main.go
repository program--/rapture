package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"rapture/cli"
	"runtime/pprof"
)

const usage string = `rapture - visualize large point datasets as rasters

Usage: rapture -i (input) -o (output) [options]

`

var path string
var output string
var stat string
var width int
var height int
var bbox string
var prof string

func init() {
	flag.StringVar(&path, "input", "", "input file")
	flag.StringVar(&path, "i", "", "input file")

	flag.StringVar(&output, "output", "", "output file path (.png)")
	flag.StringVar(&output, "o", "", "output file path (.png)")

	flag.IntVar(&width, "width", 800, "width of output")
	flag.IntVar(&width, "w", 800, "width of output")

	flag.IntVar(&height, "height", 800, "height of output")
	flag.IntVar(&height, "h", 800, "height of output")

	flag.StringVar(&stat, "stat", "density", "function for displaying values, can be one of: density, mean")
	flag.StringVar(&stat, "s", "density", "function for displaying values, can be one of: density, mean")

	flag.StringVar(&bbox, "bbox", "", "string delimited bounding box in order: xmax,xmin,ymax,ymin")
	flag.StringVar(&bbox, "b", "", "string delimited bounding box in order: xmax,xmin,ymax,ymin")

	flag.StringVar(&prof, "prof", "", "write cpu profile to file")
}

func main() {
	flag.Parse()

	if prof != "" {
		f, err := os.Create(prof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if path == "" || output == "" {
		fmt.Print(usage)
		os.Exit(1)
	} else {
		cli.Run(path, width, height, output)
	}
}
