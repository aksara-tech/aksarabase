package reflector

import (
	"aksarabase-v2/domain/interfaces"
	"aksarabase-v2/utils/stringtor"
	"fmt"
	"reflect"
)

func GetStructName(v interface{}) string {
	t := reflect.ValueOf(v)
	modelType := reflect.Indirect(t).Type()
	tableName := modelType.Name()
	return tableName
}

func GetTableName(v interface{}) string {
	t := reflect.ValueOf(v)
	modelType := reflect.Indirect(t).Type()
	tableName := modelType.Name()
	modelValue := reflect.New(modelType)
	if table, ok := modelValue.Interface().(interfaces.Table); ok {
		tableName = table.TableName()
	} else {
		fmt.Println(fmt.Sprintf("== table name not implement: %v ==", tableName))
		tableName = stringtor.ToSnakeCase(modelType.Name())
	}
	return tableName
}
