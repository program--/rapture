package config_test

import (
	"flag"
	"os"
	"rapture/pkg/config"
	"testing"
)

func TestConfig(t *testing.T) {
	os.Args = []string{
		"",
		"-i=some/fake/path.geojson",
		"-output=my/new/img.png",
		"-stat=density",
		"-property=fakeprop",
		"-width=1000",
		"-h=1000",
	}

	tc := &config.RaptureConfig{}
	tc.Init()
	flag.Parse()

	if tc.Path != "some/fake/path.geojson" {
		t.Errorf("Expected \"some/fake/path.geojson\", received: %v", tc.Path)
	}

	if tc.Output != "my/new/img.png" {
		t.Errorf("Expected \"my/new/img.png\", received: %v", tc.Output)
	}

	if tc.Stat != "density" {
		t.Errorf("Expected \"density\", received: %v", tc.Stat)
	}

	if tc.Prop != "fakeprop" {
		t.Errorf("Expected \"fakeprop\", received: %v", tc.Prop)
	}

	if tc.Width != 1000 {
		t.Errorf("Expected 1000, received: %v", tc.Width)
	}

	if tc.Height != 1000 {
		t.Errorf("Expected 1000, received: %v", tc.Height)
	}

}
