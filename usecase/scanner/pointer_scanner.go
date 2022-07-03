package scanner

import (
	"reflect"
)

type pointerScanner struct{}

func NewPointerScanner() *pointerScanner {
	return &pointerScanner{}
}

//GetListPointer Read field src struct
//  Append pointer field src struct to columns
func (s pointerScanner) GetListPointer(src interface{}, columns *[]interface{}) {
	//init to get value of dest and type detail
	val := reflect.Indirect(reflect.ValueOf(src))
	typ := reflect.TypeOf(src)

	//check if pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typF := typ.Field(i)
		x := typF.Type
		if x.Kind() == reflect.Ptr {
			x = x.Elem()
		}

		if x.Kind() == reflect.Struct {
			switch x.String() {
			case "time.Time":
				*columns = append(*columns, field.Addr().Interface())
				break
			default: //if normal struct
				if field.Kind() == reflect.Ptr {
					if field.IsNil() { //dest contain nil struct
						continue
					}
					d := field.Interface()
					s.GetListPointer(d, columns)
				} else {
					d := field.Addr().Interface()
					s.GetListPointer(d, columns)
				}
			}
		} else if x.Kind() != reflect.Slice { //if field is primitive type like string/int

			//store address field
			*columns = append(*columns, field.Addr().Interface())
		}
	}
}
