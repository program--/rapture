package grid

import (
	"image/color"

	"github.com/mazznoer/colorgrad"
)

type GridOptions struct {
	Palette         *colorgrad.Gradient
	BackgroundColor color.Color
	Bins            *uint
}
