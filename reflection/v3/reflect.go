package reflection

import "reflect"

func walk(x interface{}, fn func(str string)) {
	// fn("I'am learning the Go!")
	// reflect 反射包查看 x 的值
	val := reflect.ValueOf(x)
	// reflect 中有个方法 NumField，返回值中的字段数
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() { // Kind 方法表示字段的类型
		case reflect.String:
			fn(field.String())

		case reflect.Struct: // 遇到嵌套的 struct，则再次调用 walk 函数
			walk(field.Interface(), fn)
		}
	}
}
