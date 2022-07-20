package info

import (
	"fmt"
	"github.com/aksara-tech/aksarabase/domain/query"
	"strings"
)

//QueryInfo is a partial query builder that has modifiable structure before it's executed
type QueryInfo struct {
	SelectQuery  []string
	FromQuery    string
	WhereQuery   []string
	JoinQuery    []query.JoinRelation
	LimitQuery   string
	OrderByQuery string
}

func (q *QueryInfo) Where(column string, val interface{}) (qr *QueryInfo) {
	and := ""
	if len(q.WhereQuery) > 0 {
		and = "AND"
	}
	if strings.Contains(column, "=") || strings.Contains(column, "LIKE") {
		q.WhereQuery = append(q.WhereQuery, fmt.Sprintf("%v %v '%v'", and, column, val))
	} else {
		q.WhereQuery = append(q.WhereQuery, fmt.Sprintf("%v %v='%v'", and, column, val))
	}

	return q
}
