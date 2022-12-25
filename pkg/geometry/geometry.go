package geometry

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"math"
	"os"
	"path/filepath"
	"rapture/pkg/grid"
	"sync"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

// File -> Features -> Grid -> Canvas

var ErrUnsupportedFileType = errors.New("unsupported file type")
var ErrMissingFileExtension = errors.New("missing file extension")
var ErrNotImplemented = errors.New("not implemented")

type Collection struct {
	Features    chan orb.Geometry
	Properties  chan geojson.Properties
	NumFeatures int
	Extent      orb.Bound
}

// func MapToGrid(g geojson.Object, p string, grd *grid.Grid) {
// 	var v float64
// 	reuse := new(map[string]interface{})
// 	g.ForEach(func(geom geojson.Object) bool {
// 		pt := geom.Center()
// 		js := geom.(*geojson.Feature).JSON()
// 		json.Unmarshal([]byte(js), reuse)
//
// 		if value, ok := (*reuse)["properties"].(map[string]interface{})[p]; ok {
// 			v = value.(float64)
// 		} else {
// 			v = math.Inf(-1)
// 		}
//
// 		grd.AddCell(pt.X, pt.Y, v)
// 		return true
// 	})
// }

func MapToGrid(c *Collection, prop string, grd *grid.Grid) (int, error) {
	wg := &sync.WaitGroup{}
	wg.Add(c.NumFeatures)
	for k := 0; k < c.NumFeatures; k++ {
		go func() {
			defer wg.Done()
			f := <-c.Features
			p := <-c.Properties
			v := p.MustFloat64(prop, 0)

			switch f.GeoJSONType() {
			case geojson.TypePoint:
				grd.AddCell(f.(orb.Point).Lon(), f.(orb.Point).Lat(), v)
			case geojson.TypeLineString:
				l := f.(orb.LineString)
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
	features := make(chan orb.Geometry, nfeatures)
	properties := make(chan geojson.Properties, nfeatures)
	for _, v := range g.Features {
		go func(geom orb.Geometry, props geojson.Properties) {
			features <- geom
			properties <- props
		}(v.Geometry, v.Properties)
	}

	return &Collection{
		Features:    features,
		Properties:  properties,
		NumFeatures: nfeatures,
		Extent:      g.BBox.Bound(),
	}, nil
}

func parseGeojsonSeq(path string) (*Collection, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	nfeatures, err := countLines(f)
	if err != nil {
		return nil, err
	}

	f.Seek(0, 0) // return to beginning of file

	scanner := bufio.NewScanner(f)
	features := make(chan orb.Geometry, nfeatures)
	properties := make(chan geojson.Properties, nfeatures)
	bounds := make(chan orb.Bound, nfeatures)
	cancel := make(chan struct{})
	errc := make(chan error, 1)

	for scanner.Scan() {
		go func(line []byte, c chan struct{}) {
			g, err := geojson.UnmarshalFeature(line)
			if err != nil {
				errc <- err
				close(c)
			}

			select {
			case <-c:
				return
			default:
				features <- g.Geometry
				properties <- g.Properties
				bounds <- g.BBox.Bound()
			}
		}(scanner.Bytes(), cancel)
	}

	if err := <-errc; err != nil {
		return nil, err
	}

	count, xmax, xmin, ymax, ymin := 0, math.Inf(-1), math.Inf(1), math.Inf(-1), math.Inf(1)
	for b := range bounds {
		if count == nfeatures {
			close(bounds)
			break
		}

		if b.Max.X() > xmax {
			xmax = b.Max.X()
		}

		if b.Min.X() < xmin {
			xmin = b.Min.X()
		}

		if b.Max.Y() > ymax {
			ymax = b.Max.Y()
		}

		if b.Min.Y() < ymin {
			ymin = b.Min.Y()
		}

		count++
	}

	return &Collection{
		Features:    features,
		Properties:  properties,
		NumFeatures: nfeatures,
		Extent: orb.Bound{
			Min: orb.Point{xmin, ymin},
			Max: orb.Point{xmax, ymax},
		},
	}, nil
}

func countLines(r io.Reader) (int, error) {
	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}
