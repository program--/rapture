package geometry

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

var ErrUnsupportedFileType = errors.New("unsupported file type")
var ErrMissingFileExtension = errors.New("missing file extension")
var ErrNotImplemented = errors.New("not implemented")

type Feature struct {
	Geometry   *orb.Geometry
	Properties *geojson.Properties
}

type FeatureCollection struct {
	Features []*Feature
	Extent   *orb.Bound
}

func NewFeatureCollection(path string, filter *orb.Bound) (*FeatureCollection, error) {
	ext := filepath.Ext(path)
	d, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	switch ext {
	case ".fgb":
		return parseFlatgeobuf(d, filter)
	case ".geojson":
		return parseGeoJSON(d, filter)
	case ".geojsonl":
		fallthrough
	case ".geojsons":
		return parseGeoJSONSeq(d, filter)
	case "":
		return nil, ErrMissingFileExtension
	default:
		return nil, ErrUnsupportedFileType
	}
}
