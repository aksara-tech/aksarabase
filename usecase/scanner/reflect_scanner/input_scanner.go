package reflect_scanner

import (
	"fmt"
	"github.com/aksara-tech/aksarabase/domain/info"
	"github.com/aksara-tech/aksarabase/domain/query"
	"github.com/aksara-tech/aksarabase/usecase/scanner"
	"github.com/aksara-tech/aksarabase/utils/reflector"
	"github.com/aksara-tech/aksarabase/utils/stringtor"
	"reflect"
	"strings"
	"time"
)

type inputScanner struct{}

func (in inputScanner) ScanStruct(dest interface{}) (s info.Info, err error) {
	s.Scan.StructName = reflector.GetStructName(dest)
	s.Scan.TableName = reflector.GetTableName(dest)
	s.Scan.StructAddress = dest
	in.getColumn(dest, s.Scan.StructName, &s)

	//TODO: append SelectQuery
	for _, si := range s.Scan.ColumnWithAliases {
		s.Query.SelectQuery = append(s.Query.SelectQuery, si)
	}
	s.Query.FromQuery = fmt.Sprintf("%v %v", s.Scan.TableName, s.Scan.StructName)

	return
}

func (in inputScanner) getColumn(dest interface{}, alias string, info *info.Info) {
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

					info.Query.JoinQuery = append(info.Query.JoinQuery, relation)
					fmt.Println(relation)
					in.getColumn(field.Interface(), alias, info)
				} else {
					in.getColumn(field.Interface(), alias, info)
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

				info.Scan.ColumnWithAliases = append(info.Scan.ColumnWithAliases, colWithAlias)
				info.Scan.ColumnJson = append(info.Scan.ColumnJson, jsonTag)
				info.Scan.Columns = append(info.Scan.Columns, fieldName)
				info.Scan.Types = append(info.Scan.Types, ty.Name())
				if field.IsValid() {
					if ty.String() == "time.Time" {
						time, ok := field.Interface().(time.Time)
						if !ok {
							continue
						}

						formattedTime := time.UTC().Format("2006-01-02 03:04:05")
						info.Scan.Values = append(info.Scan.Values, formattedTime)
					} else {
						info.Scan.Values = append(info.Scan.Values, field.Interface())
					}
				} else {
					info.Scan.Values = append(info.Scan.Values, nil)
				}
			}
		}
	}

	return
}

func NewInputScanner() scanner.InputScanner {
	return &inputScanner{}
}
