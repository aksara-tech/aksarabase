package reflect_scanner

import (
	"fmt"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/info"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/query"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/scanner"
	"gitlab.com/aksaratech/aksarabase-go/v3/utils/reflector"
	"gitlab.com/aksaratech/aksarabase-go/v3/utils/stringtor"
	"reflect"
	"strings"
	"time"
)

type inputScanner struct{}

func (in inputScanner) ScanStruct(dest interface{}) (s info.ScanInfo, q info.QueryInfo, err error) {
	s.StructName = reflector.GetStructName(dest)
	s.TableName = reflector.GetTableName(dest)
	s.StructAddress = dest
	in.getColumn(dest, s.StructName, &s.Columns, &s.ColumnWithAliases, &s.ColumnJson, &s.Values, &q.Join)

	//TODO: append select
	for _, si := range s.ColumnWithAliases {
		q.Select = append(q.Select, si)
	}
	q.From = fmt.Sprintf("%v %v", s.TableName, s.StructName)

	return
}

func (in inputScanner) getColumn(dest interface{}, alias string, columns *[]string, columnWithAliases *[]string, columnJson *[]string, values *[]interface{}, relations *[]query.JoinRelation) {
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
					refTag := "id"
					if t := tag.Get("ref"); t != "" {
						refTag = t
					}
					relation := query.JoinRelation{
						Join:      "LEFT JOIN",
						TableName: fmt.Sprintf("%v %v", reflector.GetTableName(field.Interface()), typ.Field(i).Name),
						ON:        fmt.Sprintf("%v.%v=%v.%v", alias, foreignTag, typ.Field(i).Name, refTag),
					}
					alias = typ.Field(i).Name
					*relations = append(*relations, relation)
					fmt.Println(relation)
					in.getColumn(field.Interface(), alias, columns, columnWithAliases, columnJson, values, relations)
				} else {
					in.getColumn(field.Interface(), alias, columns, columnWithAliases, columnJson, values, relations)
				}
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
					if ty.String() == "time.Time" {
						time, ok := field.Interface().(time.Time)
						if !ok {
							continue
						}

						formattedTime := time.UTC().Format("2006-01-02 03:04:05")
						*values = append(*values, formattedTime)
					} else {
						*values = append(*values, field.Interface())
					}
				} else {
					*values = append(*values, nil)
				}
			}
		}
	}

	return
}

func NewInputScanner() scanner.InputScanner {
	return &inputScanner{}
}
