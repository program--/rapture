package canvas

import (
	"reflect"
)

func Density(a any, b any, max any) any {
	ai := reflect.ValueOf(a)
	bi := reflect.ValueOf(b)
	mx := reflect.ValueOf(max)
	return (ai.Float() + bi.Float()) / mx.Float()
}
