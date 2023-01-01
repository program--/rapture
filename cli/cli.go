package cli

import (
	"fmt"
	"rapture/pkg/config"
	"rapture/pkg/geometry"
	"rapture/pkg/grid"
	"rapture/pkg/util"
)

func Run(cfg config.RaptureConfig) {
	// Parse file
	fmt.Println("Parsing")
	features, err := geometry.NewFeatureCollection(cfg.Path)
	if err != nil {
		panic(err)
	}

	nfeatures := uint(len(features.Features))

	fmt.Printf("Creating grid with %d features...\n", nfeatures)
	grid.
		NewGridFromBound[float64](features.Extent, cfg.Width, cfg.Height, nfeatures).
		WithCoalescer(util.Accumulator[float64]{}).
		WithFeatures(features, cfg.Prop).
		Summarise().
		Render(cfg.Output, cfg.Padding)
}
