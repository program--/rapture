package geometry

import (
	"encoding/json"
	"math"
	"rapture/pkg/grid"

	"github.com/tidwall/geojson"
)

func MapToGrid(g geojson.Object, p string, grd *grid.Grid) {
	var v float64
	reuse := new(map[string]interface{})
	g.ForEach(func(geom geojson.Object) bool {
		pt := geom.Center()
		js := geom.(*geojson.Feature).JSON()
		json.Unmarshal([]byte(js), reuse)

		if value, ok := (*reuse)["properties"].(map[string]interface{})[p]; ok {
			v = value.(float64)
		} else {
			v = math.Inf(-1)
		}

		grd.AddCell(pt.X, pt.Y, v)
		return true
	})
}
