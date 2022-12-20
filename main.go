package main

import (
	"flag"
	"log"
	"os"
	"rapture/cli"
	"runtime/pprof"
)

var path = flag.String("path", "", "input file")
var width = flag.Int("width", 800, "width of output")
var height = flag.Int("height", 800, "height of output")
var prof = flag.String("prof", "", "write cpu profile to file")
var output = flag.String("output", "", "output file path (.png)")

func main() {
	flag.Parse()

	if *prof != "" {
		f, err := os.Create(*prof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	cli.Run(*path, *width, *height, *output)
}
