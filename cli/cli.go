package cli

import (
	"errors"
	"fmt"
	"rapture/pkg/config"
	"rapture/pkg/geometry"
	"rapture/pkg/grid"
	"rapture/pkg/util"
	"time"

	"github.com/briandowns/spinner"
)

func Run(cfg config.RaptureConfig) {
	spinner := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
	spinner.FinalMSG = fmt.Sprintf("Generated output image at %s\n", cfg.Output)
	spinner.Color("green", "bold")
	spinner.Start()

	spinner.Prefix = "rapture: 1/4 "
	spinner.Suffix = fmt.Sprintf(" parsing features from %s", cfg.Path)
	features, err := geometry.NewFeatureCollection(cfg.Path, nil)
	if err != nil {
		panic(err)
	}

	nfeatures := uint(len(features.Features))
	coalescer := util.Accumulator[float64]{}
	options := &grid.GridOptions{
		Palette:         parsePaletteFromString(cfg.Palette),
		BackgroundColor: parseColorFromString(cfg.BackgroundColor),
		Bins:            &cfg.Bins,
	}

	g := grid.
		NewGridFromBound[float64](features.Extent, cfg.Width, cfg.Height, nfeatures, spinner).
		WithHasher(util.MortonHasher{}).
		WithOptions(options).
		WithCoalescer(coalescer).
		WithFeatures(features, cfg.Prop)

	g = g.Condense()

	g = g.Summarise()

	g.Render(cfg.Output, cfg.Padding, cfg.Radius)
	spinner.Stop()
}

var errInvalidFormat = errors.New("invalid hex code format")
var errInvalidPalette = errors.New("invalid palette")
