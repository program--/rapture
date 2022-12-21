package canvas

import (
	"encoding/json"
	"errors"
	"image/color"
	"rapture/pkg/grid"
	"strings"
)

var errInvalidFormat = errors.New("invalid hex code format")

type ColorRamp struct {
	axis *grid.Axis
}

func colorFromName(s string) (c color.RGBA, err error) {
	switch strings.ToLower(s) {
	case "lightsalmon":
		return color.RGBA{R: 255, G: 160, B: 122, A: 255}, nil
	case "salmon":
		return color.RGBA{R: 250, G: 128, B: 114, A: 255}, nil
	case "darksalmon":
		return color.RGBA{R: 233, G: 150, B: 122, A: 255}, nil
	case "lightcoral":
		return color.RGBA{R: 240, G: 128, B: 128, A: 255}, nil
	case "indianred":
		return color.RGBA{R: 205, G: 92, B: 92, A: 255}, nil
	case "crimson":
		return color.RGBA{R: 220, G: 20, B: 60, A: 255}, nil
	case "firebrick":
		return color.RGBA{R: 178, G: 34, B: 34, A: 255}, nil
	case "red":
		return color.RGBA{R: 255, G: 0, B: 0, A: 255}, nil
	case "darkred":
		return color.RGBA{R: 139, G: 0, B: 0, A: 255}, nil
	case "coral":
		return color.RGBA{R: 255, G: 127, B: 80, A: 255}, nil
	case "tomato":
		return color.RGBA{R: 255, G: 99, B: 71, A: 255}, nil
	case "orangered":
		return color.RGBA{R: 255, G: 69, B: 0, A: 255}, nil
	case "gold":
		return color.RGBA{R: 255, G: 215, B: 0, A: 255}, nil
	case "orange":
		return color.RGBA{R: 255, G: 165, B: 0, A: 255}, nil
	case "darkorange":
		return color.RGBA{R: 255, G: 140, B: 0, A: 255}, nil
	case "lightyellow":
		return color.RGBA{R: 255, G: 255, B: 224, A: 255}, nil
	case "lemonchiffon":
		return color.RGBA{R: 255, G: 250, B: 205, A: 255}, nil
	case "lightgoldenrodyellow":
		return color.RGBA{R: 250, G: 250, B: 210, A: 255}, nil
	case "papayawhip":
		return color.RGBA{R: 255, G: 239, B: 213, A: 255}, nil
	case "moccasin":
		return color.RGBA{R: 255, G: 228, B: 181, A: 255}, nil
	case "peachpuff":
		return color.RGBA{R: 255, G: 218, B: 185, A: 255}, nil
	case "palegoldenrod":
		return color.RGBA{R: 238, G: 232, B: 170, A: 255}, nil
	case "khaki":
		return color.RGBA{R: 240, G: 230, B: 140, A: 255}, nil
	case "darkkhaki":
		return color.RGBA{R: 189, G: 183, B: 107, A: 255}, nil
	case "yellow":
		return color.RGBA{R: 255, G: 255, B: 0, A: 255}, nil
	case "lawngreen":
		return color.RGBA{R: 124, G: 252, B: 0, A: 255}, nil
	case "chartreuse":
		return color.RGBA{R: 127, G: 255, B: 0, A: 255}, nil
	case "limegreen":
		return color.RGBA{R: 50, G: 205, B: 50, A: 255}, nil
	case "lime":
		return color.RGBA{R: 0, G: 255, B: 0, A: 255}, nil
	case "forestgreen":
		return color.RGBA{R: 34, G: 139, B: 34, A: 255}, nil
	case "green":
		return color.RGBA{R: 0, G: 128, B: 0, A: 255}, nil
	case "darkgreen":
		return color.RGBA{R: 0, G: 100, B: 0, A: 255}, nil
	case "greenyellow":
		return color.RGBA{R: 173, G: 255, B: 47, A: 255}, nil
	case "yellowgreen":
		return color.RGBA{R: 154, G: 205, B: 50, A: 255}, nil
	case "springgreen":
		return color.RGBA{R: 0, G: 255, B: 127, A: 255}, nil
	case "mediumspringgreen":
		return color.RGBA{R: 0, G: 250, B: 154, A: 255}, nil
	case "lightgreen":
		return color.RGBA{R: 144, G: 238, B: 144, A: 255}, nil
	case "palegreen":
		return color.RGBA{R: 152, G: 251, B: 152, A: 255}, nil
	case "darkseagreen":
		return color.RGBA{R: 143, G: 188, B: 143, A: 255}, nil
	case "mediumseagreen":
		return color.RGBA{R: 60, G: 179, B: 113, A: 255}, nil
	case "seagreen":
		return color.RGBA{R: 46, G: 139, B: 87, A: 255}, nil
	case "olive":
		return color.RGBA{R: 128, G: 128, B: 0, A: 255}, nil
	case "darkolivegreen":
		return color.RGBA{R: 85, G: 107, B: 47, A: 255}, nil
	case "olivedrab":
		return color.RGBA{R: 107, G: 142, B: 35, A: 255}, nil
	case "lightcyan":
		return color.RGBA{R: 224, G: 255, B: 255, A: 255}, nil
	case "cyan":
		return color.RGBA{R: 0, G: 255, B: 255, A: 255}, nil
	case "aqua":
		return color.RGBA{R: 0, G: 255, B: 255, A: 255}, nil
	case "aquamarine":
		return color.RGBA{R: 127, G: 255, B: 212, A: 255}, nil
	case "mediumaquamarine":
		return color.RGBA{R: 102, G: 205, B: 170, A: 255}, nil
	case "paleturquoise":
		return color.RGBA{R: 175, G: 238, B: 238, A: 255}, nil
	case "turquoise":
		return color.RGBA{R: 64, G: 224, B: 208, A: 255}, nil
	case "mediumturquoise":
		return color.RGBA{R: 72, G: 209, B: 204, A: 255}, nil
	case "darkturquoise":
		return color.RGBA{R: 0, G: 206, B: 209, A: 255}, nil
	case "lightseagreen":
		return color.RGBA{R: 32, G: 178, B: 170, A: 255}, nil
	case "cadetblue":
		return color.RGBA{R: 95, G: 158, B: 160, A: 255}, nil
	case "darkcyan":
		return color.RGBA{R: 0, G: 139, B: 139, A: 255}, nil
	case "teal":
		return color.RGBA{R: 0, G: 128, B: 128, A: 255}, nil
	case "powderblue":
		return color.RGBA{R: 176, G: 224, B: 230, A: 255}, nil
	case "lightblue":
		return color.RGBA{R: 173, G: 216, B: 230, A: 255}, nil
	case "lightskyblue":
		return color.RGBA{R: 135, G: 206, B: 250, A: 255}, nil
	case "skyblue":
		return color.RGBA{R: 135, G: 206, B: 235, A: 255}, nil
	case "deepskyblue":
		return color.RGBA{R: 0, G: 191, B: 255, A: 255}, nil
	case "lightsteelblue":
		return color.RGBA{R: 176, G: 196, B: 222, A: 255}, nil
	case "dodgerblue":
		return color.RGBA{R: 30, G: 144, B: 255, A: 255}, nil
	case "cornflowerblue":
		return color.RGBA{R: 100, G: 149, B: 237, A: 255}, nil
	case "steelblue":
		return color.RGBA{R: 70, G: 130, B: 180, A: 255}, nil
	case "royalblue":
		return color.RGBA{R: 65, G: 105, B: 225, A: 255}, nil
	case "blue":
		return color.RGBA{R: 0, G: 0, B: 255, A: 255}, nil
	case "mediumblue":
		return color.RGBA{R: 0, G: 0, B: 205, A: 255}, nil
	case "darkblue":
		return color.RGBA{R: 0, G: 0, B: 139, A: 255}, nil
	case "navy":
		return color.RGBA{R: 0, G: 0, B: 128, A: 255}, nil
	case "midnightblue":
		return color.RGBA{R: 25, G: 25, B: 112, A: 255}, nil
	case "mediumslateblue":
		return color.RGBA{R: 123, G: 104, B: 238, A: 255}, nil
	case "slateblue":
		return color.RGBA{R: 106, G: 90, B: 205, A: 255}, nil
	case "darkslateblue":
		return color.RGBA{R: 72, G: 61, B: 139, A: 255}, nil
	case "lavender":
		return color.RGBA{R: 230, G: 230, B: 250, A: 255}, nil
	case "thistle":
		return color.RGBA{R: 216, G: 191, B: 216, A: 255}, nil
	case "plum":
		return color.RGBA{R: 221, G: 160, B: 221, A: 255}, nil
	case "violet":
		return color.RGBA{R: 238, G: 130, B: 238, A: 255}, nil
	case "orchid":
		return color.RGBA{R: 218, G: 112, B: 214, A: 255}, nil
	case "fuchsia":
		return color.RGBA{R: 255, G: 0, B: 255, A: 255}, nil
	case "magenta":
		return color.RGBA{R: 255, G: 0, B: 255, A: 255}, nil
	case "mediumorchid":
		return color.RGBA{R: 186, G: 85, B: 211, A: 255}, nil
	case "mediumpurple":
		return color.RGBA{R: 147, G: 112, B: 219, A: 255}, nil
	case "blueviolet":
		return color.RGBA{R: 138, G: 43, B: 226, A: 255}, nil
	case "darkviolet":
		return color.RGBA{R: 148, G: 0, B: 211, A: 255}, nil
	case "darkorchid":
		return color.RGBA{R: 153, G: 50, B: 204, A: 255}, nil
	case "darkmagenta":
		return color.RGBA{R: 139, G: 0, B: 139, A: 255}, nil
	case "purple":
		return color.RGBA{R: 128, G: 0, B: 128, A: 255}, nil
	case "indigo":
		return color.RGBA{R: 75, G: 0, B: 130, A: 255}, nil
	case "pink":
		return color.RGBA{R: 255, G: 192, B: 203, A: 255}, nil
	case "lightpink":
		return color.RGBA{R: 255, G: 182, B: 193, A: 255}, nil
	case "hotpink":
		return color.RGBA{R: 255, G: 105, B: 180, A: 255}, nil
	case "deeppink":
		return color.RGBA{R: 255, G: 20, B: 147, A: 255}, nil
	case "palevioletred":
		return color.RGBA{R: 219, G: 112, B: 147, A: 255}, nil
	case "mediumvioletred":
		return color.RGBA{R: 199, G: 21, B: 133, A: 255}, nil
	case "white":
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}, nil
	case "snow":
		return color.RGBA{R: 255, G: 250, B: 250, A: 255}, nil
	case "honeydew":
		return color.RGBA{R: 240, G: 255, B: 240, A: 255}, nil
	case "mintcream":
		return color.RGBA{R: 245, G: 255, B: 250, A: 255}, nil
	case "azure":
		return color.RGBA{R: 240, G: 255, B: 255, A: 255}, nil
	case "aliceblue":
		return color.RGBA{R: 240, G: 248, B: 255, A: 255}, nil
	case "ghostwhite":
		return color.RGBA{R: 248, G: 248, B: 255, A: 255}, nil
	case "whitesmoke":
		return color.RGBA{R: 245, G: 245, B: 245, A: 255}, nil
	case "seashell":
		return color.RGBA{R: 255, G: 245, B: 238, A: 255}, nil
	case "beige":
		return color.RGBA{R: 245, G: 245, B: 220, A: 255}, nil
	case "oldlace":
		return color.RGBA{R: 253, G: 245, B: 230, A: 255}, nil
	case "floralwhite":
		return color.RGBA{R: 255, G: 250, B: 240, A: 255}, nil
	case "ivory":
		return color.RGBA{R: 255, G: 255, B: 240, A: 255}, nil
	case "antiquewhite":
		return color.RGBA{R: 250, G: 235, B: 215, A: 255}, nil
	case "linen":
		return color.RGBA{R: 250, G: 240, B: 230, A: 255}, nil
	case "lavenderblush":
		return color.RGBA{R: 255, G: 240, B: 245, A: 255}, nil
	case "mistyrose":
		return color.RGBA{R: 255, G: 228, B: 225, A: 255}, nil
	case "gainsboro":
		return color.RGBA{R: 220, G: 220, B: 220, A: 255}, nil
	case "lightgray":
		return color.RGBA{R: 211, G: 211, B: 211, A: 255}, nil
	case "silver":
		return color.RGBA{R: 192, G: 192, B: 192, A: 255}, nil
	case "darkgray":
		return color.RGBA{R: 169, G: 169, B: 169, A: 255}, nil
	case "gray":
		return color.RGBA{R: 128, G: 128, B: 128, A: 255}, nil
	case "dimgray":
		return color.RGBA{R: 105, G: 105, B: 105, A: 255}, nil
	case "lightslategray":
		return color.RGBA{R: 119, G: 136, B: 153, A: 255}, nil
	case "slategray":
		return color.RGBA{R: 112, G: 128, B: 144, A: 255}, nil
	case "darkslategray":
		return color.RGBA{R: 47, G: 79, B: 79, A: 255}, nil
	case "black":
		return color.RGBA{R: 0, G: 0, B: 0, A: 255}, nil
	case "cornsilk":
		return color.RGBA{R: 255, G: 248, B: 220, A: 255}, nil
	case "blanchedalmond":
		return color.RGBA{R: 255, G: 235, B: 205, A: 255}, nil
	case "bisque":
		return color.RGBA{R: 255, G: 228, B: 196, A: 255}, nil
	case "navajowhite":
		return color.RGBA{R: 255, G: 222, B: 173, A: 255}, nil
	case "wheat":
		return color.RGBA{R: 245, G: 222, B: 179, A: 255}, nil
	case "burlywood":
		return color.RGBA{R: 222, G: 184, B: 135, A: 255}, nil
	case "tan":
		return color.RGBA{R: 210, G: 180, B: 140, A: 255}, nil
	case "rosybrown":
		return color.RGBA{R: 188, G: 143, B: 143, A: 255}, nil
	case "sandybrown":
		return color.RGBA{R: 244, G: 164, B: 96, A: 255}, nil
	case "goldenrod":
		return color.RGBA{R: 218, G: 165, B: 32, A: 255}, nil
	case "peru":
		return color.RGBA{R: 205, G: 133, B: 63, A: 255}, nil
	case "chocolate":
		return color.RGBA{R: 210, G: 105, B: 30, A: 255}, nil
	case "saddlebrown":
		return color.RGBA{R: 139, G: 69, B: 19, A: 255}, nil
	case "sienna":
		return color.RGBA{R: 160, G: 82, B: 45, A: 255}, nil
	case "brown":
		return color.RGBA{R: 165, G: 42, B: 42, A: 255}, nil
	case "maroon":
		return color.RGBA{R: 128, G: 0, B: 0, A: 255}, nil
	default:
		return color.RGBA{R: 0, G: 0, B: 0, A: 255}, errors.New("invalid color name")
	}
}

func colorFromJson(s string) (c color.RGBA, err error) {
	c.A = 0xff

	var rgb struct {
		R uint8 `json:"R"`
		G uint8 `json:"G"`
		B uint8 `json:"B"`
	}

	err = json.Unmarshal([]byte(s), &rgb)

	c.R = rgb.R
	c.G = rgb.G
	c.B = rgb.B

	return c, err
}

func colorFromHex(s string) (c color.RGBA, err error) {
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

	return
}
