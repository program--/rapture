package canvas

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

// Interpolate to step between from and to. from/to will be either the R, G, or B channels from start/end colors.
func (this *ColorRamp) interpolate(from uint8, to uint8, step int) uint8 {
	var (
		dis, pct float64
		cst      uint8
	)
	if from < to {
		dis = float64(to - from)
		pct = float64(step) / float64(this.steps)
		cst = from
	} else {
		dis = float64(from - to)
		pct = 1 - (float64(step) / float64(this.steps))
		cst = to
	}

	return uint8(dis*pct) + cst
}

// Generated color from color ramp at given step. Returns *color.NRGBA{} object.
func (this *ColorRamp) Generate(step int) color.Color {
	if step == 0 {
		return this.startColor
	}

	if step == this.steps {
		return this.endColor
	}

	if step < 0 || step > this.steps {
		return nil
	}

	s := this.startColor.(*color.NRGBA)
	e := this.endColor.(*color.NRGBA)

	return &color.NRGBA{
		R: this.interpolate(s.R, e.R, step),
		G: this.interpolate(s.G, e.G, step),
		B: this.interpolate(s.B, e.B, step),
		A: this.interpolate(s.A, e.A, step),
	}
}
