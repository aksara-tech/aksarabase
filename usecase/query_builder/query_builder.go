package query_builder

import (
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/info"
)

type InsertQueryBuilder interface {
	//BuildInsertQuery build insert query using ScanInfo
	BuildInsertQuery(info info.ScanInfo, qInfo info.QueryInfo) string
}

type SelectQueryBuilder interface {
	/*
		Convert ScanInfo from struct Scanner and QueryInfo, which has
		  Select, From, Where, Join, Limit, OrderBy to string query
	*/
	BuildSelectQuery(info info.ScanInfo, qInfo info.QueryInfo) string
}

type UpdateQueryBuilder interface {
	//BuildUpdateQuery build update query using where in QueryInfo and ScanInfo
	BuildUpdateQuery(info info.ScanInfo, qInfo info.QueryInfo) string
}

type QueryBuilder interface {
	InsertQueryBuilder
	UpdateQueryBuilder
	SelectQueryBuilder
}
