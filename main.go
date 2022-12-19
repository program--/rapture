package main

import (
	"flag"
	"rapture/cli"
)

var path = flag.String("path", "", "")

func main() {
	flag.Parse()
	cli.Run(*path)
}
