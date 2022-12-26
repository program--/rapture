package geometry

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"rapture/pkg/grid"
	"sync"

	jsoniter "github.com/json-iterator/go"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

// File -> Features -> Grid -> Canvas

var ErrUnsupportedFileType = errors.New("unsupported file type")
var ErrMissingFileExtension = errors.New("missing file extension")
var ErrNotImplemented = errors.New("not implemented")

var c = jsoniter.Config{
	EscapeHTML:              true,
	SortMapKeys:             false,
	MarshalFloatWith6Digits: true,
}.Froze()

func init() {
	geojson.CustomJSONMarshaler = c
	geojson.CustomJSONUnmarshaler = c
}

type Feature struct {
	Geometry   orb.Geometry
	Properties geojson.Properties
}

type Collection struct {
	Features    chan Feature
	NumFeatures int
	Extent      orb.Bound
}

func MapToGrid(c *Collection, prop string, grd *grid.Grid) (int, error) {
	wg := &sync.WaitGroup{}
	wg.Add(c.NumFeatures)
	for k := 0; k < c.NumFeatures; k++ {
		go func() {
			defer wg.Done()
			f := <-c.Features
			v := f.Properties.MustFloat64(prop, 0)

			switch f.Geometry.GeoJSONType() {
			case geojson.TypePoint:
				grd.AddCell(f.Geometry.(orb.Point).Lon(), f.Geometry.(orb.Point).Lat(), v)
			case geojson.TypeLineString:
				l := f.Geometry.(orb.LineString)
				for i := 0; i < len(l)-1; i++ {
					a := l[i]
					b := l[i+1]
					grd.AddLine(a.Lon(), a.Lat(), b.Lon(), b.Lat(), v)
				}
			default:
				panic("maptogrid failure")
			}
		}()
	}

	wg.Wait()
	return c.NumFeatures, nil
}

func Parse(path string) (*Collection, error) {
	// determine path file type
	ext := filepath.Ext(path)

	switch ext {
	case ".fgb":
		return parseFlatgeobuf(path)
	case ".geojson":
		return parseGeojson(path)
	case ".geojsonl":
		fallthrough
	case ".geojsons":
		return parseGeojsonSeq(path)
	case "":
		return nil, ErrMissingFileExtension
	default:
		return nil, ErrUnsupportedFileType
	}
}

func parseFlatgeobuf(path string) (*Collection, error) {
	return nil, ErrNotImplemented
}

func parseGeojson(path string) (*Collection, error) {
	d, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	g, err := geojson.UnmarshalFeatureCollection(d)
	if err != nil {
		return nil, err
	}

	nfeatures := len(g.Features)
	features := make(chan Feature, nfeatures)
	for _, v := range g.Features {
		go func(geom orb.Geometry, props geojson.Properties) {
			features <- Feature{
				Geometry:   geom,
				Properties: props,
			}
		}(v.Geometry, v.Properties)
	}

	return &Collection{
		Features:    features,
		NumFeatures: nfeatures,
		Extent:      g.BBox.Bound(),
	}, nil
}

func parseGeojsonSeq(path string) (*Collection, error) {
	d, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := bytes.Split(d, []byte("\n"))
	features := make(chan Feature, len(lines))
	nfeatures := 0

	for _, line := range lines {
		if len(line) != 0 {
			go func(l []byte) {
				// UNSAFE
				g, _ := geojson.UnmarshalFeature(l)
				features <- Feature{
					Geometry:   g.Geometry,
					Properties: g.Properties,
				}
			}(line)
			nfeatures++
		}
	}

	return &Collection{
		Features:    features,
		NumFeatures: nfeatures,
	}, nil
}
