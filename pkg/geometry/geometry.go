package geometry

import (
	"math"
	"rapture/pkg/grid"

	"github.com/tidwall/geojson"
)

type GeojsonSummary struct {
	MaxVal float64
	MinVal float64
	AvgVal float64
}

func MapToGrid(g geojson.Object, grd *grid.Grid) *GeojsonSummary {
	summary := &GeojsonSummary{
		MaxVal: math.Inf(-1),
		MinVal: math.Inf(1),
		AvgVal: 0.0,
	}

	g.ForEach(func(geom geojson.Object) bool {
		p := geom.Center()
		v := 1.0

		if v > summary.MaxVal {
			summary.MaxVal = v
		}

		if v < summary.MinVal {
			summary.MinVal = v
		}

		summary.AvgVal += v
		grd.AddCell(p.X, p.Y, v)
		return true
	})

	summary.AvgVal = summary.AvgVal / float64(grd.Cells().Len())

	return summary
}
