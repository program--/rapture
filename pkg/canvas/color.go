package canvas

import (
	"encoding/json"
	"errors"
	"image/color"
)

var errInvalidFormat = errors.New("invalid hex code format")

func colorFromJson(s string) (color.Color, error) {
	c := &color.NRGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}

	var rgb struct {
		R *uint8 `json:"R"`
		G *uint8 `json:"G"`
		B *uint8 `json:"B"`
		A *uint8 `json:"A"`
	}

	err := json.Unmarshal([]byte(s), &rgb)

	if rgb.R != nil {
		c.R = *rgb.R
	}

	if rgb.G != nil {
		c.G = *rgb.G
	}

	if rgb.B != nil {
		c.B = *rgb.B
	}

	if rgb.A != nil {
		c.A = *rgb.A
	}

	return c, err
}

func ColorFromHex(s string) (*color.NRGBA, error) {
	var err error

	c := new(color.NRGBA)
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
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
		err = errInvalidFormat
	}

	return c, err
}
