package mysql

import (
	"aksarabase-v2/usecase/scanner"
	"fmt"
	"strings"
)

func (b queryBuilderMysql) BuildInsertQuery(info scanner.ScanInfo, addQuery ...string) string {
	var params []string
	i := 0

	for _, j := range info.ColumnJson {
		if info.Values[i] != nil {
			params = append(params, fmt.Sprintf("%v='%v'", j, info.Values[i]))
		}
		i++
	}

	return fmt.Sprintf("INSERT %v SET %v %v", info.TableName, strings.Join(params, ","), strings.Join(addQuery, " "))
}
