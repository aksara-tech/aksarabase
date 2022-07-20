package mysql

import (
	"fmt"
	"github.com/aksara-tech/aksarabase/domain/info"
	"strings"
)

type SelectBuilder struct{}

func NewSelectBuilder() *SelectBuilder {
	return &SelectBuilder{}
}

func (b SelectBuilder) BuildSelect(i info.Info) string {
	//TODO: build query SelectQuery
	s := "*"
	if len(i.Query.SelectQuery) > 0 {
		s = strings.Join(i.Query.SelectQuery, ",")
	}

	//TODO: build query where
	var w string
	if len(i.Query.WhereQuery) > 0 {
		w = fmt.Sprintf("WHERE %v", strings.Join(i.Query.WhereQuery, " "))
	}

	//TODO: build join
	var j string
	if len(i.Query.JoinQuery) > 0 {
		var joins []string
		for _, join := range i.Query.JoinQuery {
			joins = append(joins, fmt.Sprintf("%v %v ON %v ", join.Join, join.TableName, join.ON))
		}

		j = strings.Join(joins, " ")
	}

	//TODO: build limit
	var l string
	if i.Query.LimitQuery != "" {
		l = fmt.Sprintf("LIMIT %v", i.Query.LimitQuery)
	}
	//TODO: build order by
	var o string
	if i.Query.OrderByQuery != "" {
		o = fmt.Sprintf("ORDER BY %v", i.Query.OrderByQuery)
	}

	query := fmt.Sprintf("SELECT %v FROM %v %v %v %v %v", s, i.Query.FromQuery, j, w, o, l)

	return query
}
