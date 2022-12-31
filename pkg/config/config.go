package config

import (
	"flag"
	"reflect"
	"strconv"
)

type RaptureConfig struct {
	Path    string `flag:"input" usage:"input file" alias:"i"`
	Output  string `flag:"output" usage:"output file path (.png)" alias:"o"`
	Stat    string `flag:"stat" usage:"function for displaying values, can be of one: density, mean" default:"density" alias:"s"`
	Prop    string `flag:"property" usage:"property to use from input" alias:"p"`
	Width   uint   `flag:"width" usage:"width of output" default:"800" alias:"w"`
	Height  uint   `flag:"height" usage:"height of output" default:"800" alias:"h"`
	Padding uint   `flag:"padding" usage:"padding around main image" alias:"pad"`
	Bbox    string `flag:"bbox" usage:"string delimited bounding box in order: xmax, xmin, ymax, ymin" alias:"b"`
	Prof    string `flag:"prof" usage:"write cpu profile to file"`
}

// Sets up flag calls to struct fields based on struct tags/type
func (c *RaptureConfig) Init() {
	typeref := reflect.TypeOf(c).Elem()
	valref := reflect.ValueOf(c).Elem()
	nfields := valref.NumField()

	for i := 0; i < nfields; i++ {
		field := valref.Field(i)
		tags := typeref.Field(i).Tag
		ptr := field.Addr().UnsafePointer()

		flag_name := tags.Get("flag")
		flag_alias := tags.Get("alias")
		flag_default := tags.Get("default")
		flag_usage := tags.Get("usage")

		switch field.Kind() {
		case reflect.String:
			flag.StringVar((*string)(ptr), flag_name, flag_default, flag_usage)
			if flag_alias != "" {
				flag.StringVar((*string)(ptr), flag_alias, flag_default, flag_usage)
			}
			continue
		case reflect.Uint:
			flag_default_uint64, _ := strconv.ParseUint(flag_default, 10, 64)
			flag_default_uint := uint(flag_default_uint64)
			flag.UintVar((*uint)(ptr), flag_name, flag_default_uint, flag_usage)
			if flag_alias != "" {
				flag.UintVar((*uint)(ptr), flag_alias, flag_default_uint, flag_usage)
			}
			continue
		}
	}
}
