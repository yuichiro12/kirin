package kirin

import "reflect"

func Walk(v interface{}, f func(a interface{})) {
	walk(reflect.ValueOf(v), f)
}

func walk(v reflect.Value, f func(a interface{})) {
	switch v.Kind() {
	case reflect.Invalid:
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i), f)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			walk(v.Field(i), f)
		}
	case reflect.Map:
		for _, i := range v.MapKeys() {
			walk(v.MapIndex(i), f)
		}
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			f(nil)
		} else {
			walk(v.Elem(), f)
		}
	default:
		f(v)
	}
}

