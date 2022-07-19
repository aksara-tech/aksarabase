package mysql

import (
	"fmt"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/constanta"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/info"
	"strings"
	"time"
)

type updateBuilder struct{}

func NewUpdateBuilder() *updateBuilder {
	return &updateBuilder{}
}

func (b updateBuilder) BuildUpdateQuery(info info.Info) string {
	var params []string
	i := 0

	for _, j := range info.Scan.ColumnJson {
		if fmt.Sprintf("%v", info.Scan.Values[i]) == "" {
			goto next
		}
		if fmt.Sprintf("%v", info.Scan.Values[i]) == "0" {
			goto next
		}
		if fmt.Sprintf("%v", info.Scan.Values[i]) == constanta.TIME_NIL {
			goto next
		}
		if j == "updated_at" {
			value := fmt.Sprintf("%v='%v'", j, time.Now().UTC().Format(constanta.TIME_LAYOUT))
			params = append(params, value)
			goto next
		}

		if info.Scan.Values[i] != nil {
			value := fmt.Sprintf("%v='%v'", j, info.Scan.Values[i])

			params = append(params, value)
		}
	next:
		i++
	}

	query := fmt.Sprintf("Update %v SET %v WHERE %v", info.Scan.TableName, strings.Join(params, ","), strings.Join(info.Query.WhereQuery, ","))
	return query
}
