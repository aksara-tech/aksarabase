package mysql

import (
	"aksarabase-v2/usecase/scanner"
	"fmt"
	"strings"
)

func (b queryBuilderMysql) BuildSelectQuery(info scanner.ScanInfo, addQuery ...string) string {
	tableName := info.TableName
	param := strings.Join(info.ColumnWithAliases, ",")
	query := fmt.Sprintf("select %v from %v %v %v", param, tableName, info.StructName, strings.Join(addQuery, ","))
	return query
}
