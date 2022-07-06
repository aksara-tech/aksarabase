package mysql

import (
	"aksarabase-v2/domain/info"
	"fmt"
	"strings"
)

type selectBuilder struct{}

func NewSelectBuilder() *selectBuilder {
	return &selectBuilder{}
}

func (b selectBuilder) BuildSelectQuery(i info.ScanInfo, q info.QueryInfo) string {
	//TODO: build query select
	s := "*"
	if len(q.Select) > 0 {
		s = strings.Join(q.Select, ",")
	}

	//TODO: build query where
	var w string
	if len(q.Where) > 0 {
		w = fmt.Sprintf("WHERE %v", strings.Join(q.Where, " "))
	}

	//TODO: build join
	var j string
	if len(q.Join) > 0 {
		var joins []string
		for _, join := range q.Join {
			joins = append(joins, fmt.Sprintf("%v %v ON %v ", join.Join, join.TableName, join.ON))
		}

		j = strings.Join(joins, " ")
	}

	//TODO: build limit
	var l string
	if q.Limit > 0 {
		l = fmt.Sprintf("LIMIT %v", q.Limit)
	}
	//TODO: build order by
	var o string
	if q.OrderBy != "" {
		o = fmt.Sprintf("ORDER BY %v", q.OrderBy)
	}

	query := fmt.Sprintf("SELECT %v FROM %v %v %v %v %v", s, q.From, j, w, o, l)

	return query
}
