package cli

import (
	"image/color"
	"strings"

	"github.com/mazznoer/colorgrad"
)

func parseColorFromString(s string) color.Color {
	c := new(color.NRGBA)
	c.A = 0xff

	if s == "" {
		return nil
	}

	if s[0] != '#' {
		panic(errInvalidFormat)
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		default:
			panic(errInvalidFormat)
		}
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		panic(errInvalidFormat)
	}

	return c
}

func parsePaletteFromString(s string) *colorgrad.Gradient {
	var p colorgrad.Gradient

	switch strings.ToLower(s) {
	case "brbg":
		p = colorgrad.BrBG()
	case "prgn":
		p = colorgrad.PRGn()
	case "piyg":
		p = colorgrad.PiYG()
	case "puor":
		p = colorgrad.PuOr()
	case "rdbu":
		p = colorgrad.RdBu()
	case "rdgy":
		p = colorgrad.RdGy()
	case "rdylbu":
		p = colorgrad.RdYlBu()
	case "rdylgn":
		p = colorgrad.RdYlGn()
	case "spectral":
		p = colorgrad.Spectral()
	case "blues":
		p = colorgrad.Blues()
	case "greens":
		p = colorgrad.Greens()
	case "greys":
		p = colorgrad.Greys()
	case "oranges":
		p = colorgrad.Oranges()
	case "purples":
		p = colorgrad.Purples()
	case "reds":
		p = colorgrad.Reds()
	case "turbo":
		p = colorgrad.Turbo()
	case "viridis":
		p = colorgrad.Viridis()
	case "inferno":
		p = colorgrad.Inferno()
	case "magma":
		p = colorgrad.Magma()
	case "plasma":
		p = colorgrad.Plasma()
	case "cividis":
		p = colorgrad.Cividis()
	case "warm":
		p = colorgrad.Warm()
	case "cool":
		p = colorgrad.Cool()
	case "cubehelix":
		fallthrough
	case "cubehelixdefault":
		p = colorgrad.CubehelixDefault()
	case "bugn":
		p = colorgrad.BuGn()
	case "bupu":
		p = colorgrad.BuPu()
	case "gnbu":
		p = colorgrad.GnBu()
	case "orrd":
		p = colorgrad.OrRd()
	case "pubugn":
		p = colorgrad.PuBuGn()
	case "pubu":
		p = colorgrad.PuBu()
	case "purd":
		p = colorgrad.PuRd()
	case "rdpu":
		p = colorgrad.RdPu()
	case "ylgnbu":
		p = colorgrad.YlGnBu()
	case "ylgn":
		p = colorgrad.YlGn()
	case "ylorbr":
		p = colorgrad.YlOrBr()
	case "ylorrd":
		p = colorgrad.YlOrRd()
	case "rainbow":
		p = colorgrad.Rainbow()
	case "sinebow":
		p = colorgrad.Sinebow()
	default:
		panic(errInvalidPalette)
	}

	return &p
}
