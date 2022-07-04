package query_builder

import "aksarabase-v2/usecase/scanner"

type InsertQueryBuilder interface {
	//BuildInsertQuery build insert query
	//  INSERT %v SET %v created_at=NOW()
	BuildInsertQuery(info scanner.ScanInfo, addQuery ...string) string
}

type SelectQueryBuilder interface {
	BuildSelectQuery(info scanner.ScanInfo, addQuery ...string) string
}

type UpdateQueryBuilder interface {
	BuildUpdateQuery(info scanner.ScanInfo, addQuery ...string) string
}

type QueryBuilder interface {
	InsertQueryBuilder
	UpdateQueryBuilder
	SelectQueryBuilder
}
