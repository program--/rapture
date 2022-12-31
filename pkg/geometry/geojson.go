package geometry

import (
	"bytes"
	"math"
	"sync"
	"sync/atomic"

	jsoniter "github.com/json-iterator/go"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

var c = jsoniter.Config{
	EscapeHTML:              true,
	SortMapKeys:             false,
	MarshalFloatWith6Digits: true,
}.Froze()

func init() {
	geojson.CustomJSONMarshaler = c
	geojson.CustomJSONUnmarshaler = c
}

func parseGeoJSON(data []byte) (*FeatureCollection, error) {
	g, err := geojson.UnmarshalFeatureCollection(data)
	if err != nil {
		return nil, err
	}

	numFeatures := len(g.Features)
	features := make([]*Feature, numFeatures)
	allBounds := make([]*orb.Bound, numFeatures)

	wg := sync.WaitGroup{}
	wg.Add(numFeatures)
	setSlices := func(i int, f *geojson.Feature) {
		defer wg.Done()
		b := f.Geometry.Bound()
		features[i] = geojsonToFeature(f)
		allBounds[i] = &b
	}

	for i, v := range g.Features {
		go setSlices(i, v)
	}
	wg.Wait()

	extent := compareBounds(allBounds)
	return &FeatureCollection{features, &extent}, nil
}

func parseGeoJSONSeq(data []byte) (*FeatureCollection, error) {
	lines := bytes.Split(data, []byte("\n"))
	numLines := len(lines)
	features := make([]*Feature, numLines)
	allBounds := make([]*orb.Bound, numLines)
	numFeatures := uint64(0)

	wg := sync.WaitGroup{}
	setSliceFromByte := func(i int, d []byte) {
		defer wg.Done()
		g, err := geojson.UnmarshalFeature(d)
		if err == nil {
			b := g.Geometry.Bound()
			features[i] = geojsonToFeature(g)
			allBounds[i] = &b
			atomic.AddUint64(&numFeatures, 1)
		}
	}

	for i, line := range lines {
		if len(line) != 0 {
			// line is not empty
			wg.Add(1)
			go setSliceFromByte(i, line)
		}
	}
	wg.Wait()

	extent := compareBounds(allBounds)
	return &FeatureCollection{features[:numFeatures], &extent}, nil

}

func geojsonToFeature(gj *geojson.Feature) *Feature {
	return &Feature{
		Geometry:   &gj.Geometry,
		Properties: &gj.Properties,
	}
}

func compareBounds(bounds []*orb.Bound) orb.Bound {
	xmax, xmin := math.Inf(-1), math.Inf(1)
	ymax, ymin := math.Inf(-1), math.Inf(1)
	for _, b := range bounds {
		if b != nil {
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
		}
	}

	return orb.Bound{
		Max: orb.Point{xmax, ymax},
		Min: orb.Point{xmin, ymin},
	}
}
