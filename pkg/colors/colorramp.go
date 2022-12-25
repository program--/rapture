package colors

import (
	"image/color"
)

type ColorRamp struct {
	startColor color.Color
	endColor   color.Color
	steps      int
}

func NewColorRamp(from color.Color, to color.Color, steps int) *ColorRamp {
	f := color.NRGBAModel.Convert(from).(color.NRGBA)
	t := color.NRGBAModel.Convert(to).(color.NRGBA)
	return &ColorRamp{
		startColor: &f,
		endColor:   &t,
		steps:      steps,
	}
}

func (cr *ColorRamp) Steps() int {
	return cr.steps
}

// Interpolate to step between from and to. from/to will be either the R, G, or B channels from start/end colors.
func (cr *ColorRamp) interpolate(from uint8, to uint8, step int) uint8 {
	var (
		dis, pct float64
		cst      uint8
	)
	if from < to {
		dis = float64(to - from)
		pct = float64(step) / float64(cr.steps)
		cst = from
	} else {
		dis = float64(from - to)
		pct = 1 - (float64(step) / float64(cr.steps))
		cst = to
	}

	return uint8(dis*pct) + cst
}

// Generated color from color ramp at given step. Returns *color.NRGBA{} object.
func (cr *ColorRamp) Generate(step int) color.Color {
	if step == 0 {
		return cr.startColor
	}

	if step == cr.steps {
		return cr.endColor
	}

	if step < 0 || step > cr.steps {
		return nil
	}

	s := cr.startColor.(*color.NRGBA)
	e := cr.endColor.(*color.NRGBA)

	return &color.NRGBA{
		R: cr.interpolate(s.R, e.R, step),
		G: cr.interpolate(s.G, e.G, step),
		B: cr.interpolate(s.B, e.B, step),
		A: cr.interpolate(s.A, e.A, step),
	}
}
