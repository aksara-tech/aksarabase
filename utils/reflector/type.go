package reflector

import "reflect"

func GetKind(dest interface{}) string {
	typ := reflect.TypeOf(dest)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	x := typ.Kind().String()
	return x
}
