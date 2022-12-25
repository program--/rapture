package geometry_test

import (
	"rapture/pkg/geometry"
	"testing"
)

func TestParsing(t *testing.T) {
	geojson_path := "../../test/data/example.geojson"
	// geojsons_path := "../../test/data/example.geojsonl"
	// fgb_path := "test/data/example.fgb"

	testPath(geojson_path, t)
	// testPath(geojsons_path, t)
}

func testPath(path string, t *testing.T) {
	c, err := geometry.Parse(path)
	if err != nil {
		t.Error(err)
	}

	if c.NumFeatures != 100 {
		t.Error("NumFeatures is not 100")
	}
}
