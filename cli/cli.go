package cli

import (
	"fmt"
	"rapture/pkg/config"
	"rapture/pkg/geometry"
	"rapture/pkg/grid"
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
		NewGridFromBound[float64](features.Extent, cfg.Width+cfg.Padding, cfg.Height+cfg.Padding, nfeatures).
		WithCoalescer(grid.Accumulator[float64]{}).
		WithFeatures(features, cfg.Prop).
		Summarise().
		Render(cfg.Output)
}
