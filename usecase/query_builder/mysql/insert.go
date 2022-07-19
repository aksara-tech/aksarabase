package mysql

import (
	"fmt"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/constanta"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/info"
	"strings"
	"time"
)

type insertBuilder struct{}

func NewInsertBuilder() *insertBuilder {
	return &insertBuilder{}
}

func (b insertBuilder) BuildInsertQuery(info info.Info) string {
	var values []string
	var columns []string
	i := 0
	for _, value := range info.Scan.Values {
		if fmt.Sprintf("%v", value) == "0" {
			goto next
		}

		if value != nil {
			if info.Scan.Types[i] == "Time" {
				value = time.Now().UTC().Format(constanta.TIME_LAYOUT)
			}

			if value == constanta.TIME_NIL {
				goto next
			}

			values = append(values, fmt.Sprintf("'%v'", value))
			columns = append(columns, fmt.Sprintf("%v", info.Scan.ColumnJson[i]))
		}
	next:
		{
			i++
		}
	}

	query := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", info.Scan.TableName, strings.Join(columns, ","), strings.Join(values, ","))
	return query
}
