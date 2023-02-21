package reflection

import "reflect"

func walk(x interface{}, fn func(str string)) {
	// fn("I'am learning the Go!")
	// reflect 反射包
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())

}
