package mysql

import (
	"aksarabase-v2/domain/constanta"
	"aksarabase-v2/domain/info"
	"fmt"
	"strings"
	"time"
)

type updateBuilder struct{}

func NewUpdateBuilder() *updateBuilder {
	return &updateBuilder{}
}

func (b updateBuilder) BuildUpdateQuery(info info.ScanInfo, qInfo info.QueryInfo) string {
	var params []string
	i := 0

	for _, j := range info.ColumnJson {
		if fmt.Sprintf("%v", info.Values[i]) == "" {
			goto next
		}
		if fmt.Sprintf("%v", info.Values[i]) == "0" {
			goto next
		}
		if fmt.Sprintf("%v", info.Values[i]) == constanta.TIME_NIL {
			goto next
		}
		if j == "updated_at" {
			value := fmt.Sprintf("%v='%v'", j, time.Now().UTC().Format(constanta.TIME_LAYOUT))
			params = append(params, value)
			goto next
		}

		if info.Values[i] != nil {
			value := fmt.Sprintf("%v='%v'", j, info.Values[i])

			params = append(params, value)
		}
	next:
		i++
	}

	query := fmt.Sprintf("Update %v SET %v WHERE %v", info.TableName, strings.Join(params, ","), strings.Join(qInfo.Where, ","))
	return query
}
