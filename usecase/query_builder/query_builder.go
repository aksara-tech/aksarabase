package query_builder

import (
	"github.com/aksara-tech/aksarabase/domain/info"
)

type InsertQueryBuilder interface {
	//BuildInsertQuery build insert query using ScanInfo
	BuildInsertQuery(info info.Info) string
}

type SelectBuilder interface {
	/*
		Convert ScanInfo from struct Scanner and QueryInfo, which has
		  SelectQuery, From, Where, Join, Limit, OrderBy to string query
	*/
	BuildSelect(info info.Info) string
}

type UpdateQueryBuilder interface {
	//BuildUpdateQuery build update query using where in QueryInfo and ScanInfo
	BuildUpdateQuery(info info.Info) string
}

type QueryBuilder interface {
	InsertQueryBuilder
	UpdateQueryBuilder
	SelectBuilder
}
