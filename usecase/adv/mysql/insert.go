package mysql

import (
	"context"
	"fmt"
	"strings"
)

func (m mysqlDB) Insert(ctx context.Context, dest interface{}) error {
	//TODO: get list column tag and table name
	res, err := m.inputScanner.ScanStruct(dest)
	if err != nil {
		return err
	}

	var params []string
	i := 0
	for _, j := range res.ColumnJson {
		if res.Values[i] != nil {
			params = append(params, fmt.Sprintf("%v='%v'", j, res.Values[i]))
		}
		i++
	}

	//TODO: build query
	query := fmt.Sprintf("INSERT %v SET %v created_at=NOW()", res.TableName, strings.Join(params, ","))

	fmt.Println(query)
	//TODO: exec query
	m.db.Exec(query)
	panic("implement me")
}
