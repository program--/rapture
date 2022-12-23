package canvas

import (
	"reflect"
)

func Density(a any, b any, _ any) any {
	ai := reflect.ValueOf(a)
	bi := reflect.ValueOf(b)
	return ai.Float() + bi.Float()
}
