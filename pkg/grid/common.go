package grid

import (
	"fmt"

	"github.com/paulmach/orb"
)

type uint_t interface{ uint | uint32 | uint64 }
type int_t interface{ int | int32 | int64 }
type float_t interface{ float32 | float64 }
type numeric_t interface{ uint_t | int_t | float_t }
type cell_t numeric_t

func checkGridIndex(column int, row int, p *orb.Point) error {
	if column == -1 && row == -1 {
		return fmt.Errorf("point's x and y coordinates (%f, %f) not within axis bounds", p.X(), p.Y())
	}

	if column == -1 {
		return fmt.Errorf("point's x-coordinate (%f) not within axis bounds", p.X())
	}

	if row == -1 {
		return fmt.Errorf("point's y-coordinate (%f) not within axis bounds", p.Y())
	}

	return nil
}
