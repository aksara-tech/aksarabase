package scanner

import (
	"aksarabase-v2/utils/reflector"
	"aksarabase-v2/utils/stringtor"
	"fmt"
	"reflect"
	"strings"
)

type inputScanner struct{}

func (in inputScanner) ScanStruct(dest interface{}) (s ScanInfo, err error) {
	s.StructName = reflector.GetStructName(dest)
	s.TableName = reflector.GetTableName(dest)
	in.getColumn(dest, s.StructName, &s.Columns, &s.ColumnWithAliases, &s.ColumnJson, &s.Values)
	return
}

func (in inputScanner) getColumn(dest interface{}, alias string, columns *[]string, columnWithAliases *[]string, columnJson *[]string, values *[]interface{}) {
	val := reflect.ValueOf(dest)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := reflect.TypeOf(dest)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	strict := true
	numField := val.NumField()
	for i := 0; i < numField; i++ {
		field := val.Field(i)
		if val.Field(i).Kind() == reflect.Ptr {
			field = field.Elem()
		}

		ty := typ.Field(i).Type
		if ty.Kind() == reflect.Ptr {
			ty = typ.Field(i).Type.Elem()
		}

		if strict {
			tag := val.Type().Field(i).Tag
			if ty.Kind() == reflect.Struct && ty.String() != "time.Time" {
				if !field.IsValid() {
					continue
				}
				//set alias if it is foreign struct
				foreignTag := tag.Get("foreign")
				if foreignTag != "" {
					alias = typ.Field(i).Name
				}
				in.getColumn(field.Interface(), alias, columns, columnWithAliases, columnJson, values)

			} else {
				if ty.Kind() == reflect.Slice {
					continue
				}

				jsonTag := tag.Get("json")
				jsonTags := strings.Split(jsonTag, ",")
				jsonTag = jsonTags[0]
				fieldName := typ.Field(i).Name
				if strings.Contains(jsonTag, "-") || jsonTag == "" {
					jsonTag = stringtor.ToSnakeCase(fieldName)
				}

				colWithAlias := fmt.Sprintf("%v.%v", alias, jsonTag)

				*columnWithAliases = append(*columnWithAliases, colWithAlias)
				*columnJson = append(*columnJson, jsonTag)
				*columns = append(*columns, fieldName)
				if field.IsValid() {
					*values = append(*values, field.Interface())
				} else {
					*values = append(*values, nil)
				}
			}
		}
	}

	return
}

func NewInputScanner() InputScanner {
	return &inputScanner{}
}
