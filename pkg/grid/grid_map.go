package grid

import (
	"rapture/pkg/geometry"
	"sync"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

func (grd *Grid[T]) WithFeatures(f *geometry.FeatureCollection, property string) *Grid[T] {
	wg := &sync.WaitGroup{}
	mapToGrid := func(feature *geometry.Feature) {
		defer wg.Done()
		value := getProp[T](feature.Properties, property)

		switch (*feature.Geometry).GeoJSONType() {
		case geojson.TypePoint:
			pt := (*feature.Geometry).(orb.Point)
			grd.AddPoint(&pt, value)
		case geojson.TypeLineString:
			panic("not implemented")
		default:
			panic("failed to map features")
		}
	}

	// process
	for _, v := range f.Features {
		wg.Add(1)
		go mapToGrid(v)
	}
	wg.Wait()

	return grd
}

func getProp[T cell_t](p *geojson.Properties, key string) (property T) {
	switch any(property).(type) {
	case int:
		property = T(p.MustInt(key, 0))
	case int32:
		property = T(p.MustInt(key, 0))
	case int64:
		property = T(p.MustInt(key, 0))
	case uint:
		property = T(p.MustInt(key, 0))
	case uint32:
		property = T(p.MustInt(key, 0))
	case uint64:
		property = T(p.MustInt(key, 0))
	case float32:
		property = T(p.MustFloat64(key, 0.0))
	case float64:
		property = T(p.MustFloat64(key, 0.0))
	}
	return
}
