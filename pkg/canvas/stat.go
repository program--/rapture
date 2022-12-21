package canvas

import "reflect"

func Density(a any, b any) any {
	ai := reflect.ValueOf(a)
	bi := reflect.ValueOf(b)
	return ai.Int() + bi.Int()
}
