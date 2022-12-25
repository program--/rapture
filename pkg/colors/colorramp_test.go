package colors_test

import (
	"image/color"
	"rapture/pkg/colors"
	"testing"
)

func TestColorRamp(t *testing.T) {
	start, err := colors.ColorFromHex("#000D73")
	if err != nil {
		t.Fatal("Failed to create start color\n")
	}

	end, err := colors.ColorFromHex("#032823")
	if err != nil {
		t.Fatal("Failed to create end color\n")
	}

	cr := colors.NewColorRamp(start, end, 8)
	expected := []string{
		"#000D73",
		"#001069",
		"#00135F",
		"#011755",
		"#011A4B",
		"#011D41",
		"#022137",
		"#02242D",
		"#032823",
	}

	for i := 0; i < 9; i++ {
		expect, err := colors.ColorFromHex(expected[i])
		if err != nil {
			t.Fatalf("Failed to create %d expected color\n", i)
		}

		generated := cr.Generate(i)
		t.Logf("[ GENERATED: %3v | EXPECTED:%3v ]", generated, expect)
		if !colorsEqual(generated, expect) {
			t.Fail()
		}
	}
}

func colorsEqual(a color.Color, b color.Color) bool {
	x := color.NRGBAModel.Convert(a).(color.NRGBA)
	y := color.NRGBAModel.Convert(b).(color.NRGBA)
	return x.R == y.R && x.G == y.G && x.B == y.B && x.A == y.A
}
