package mysql

import "gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_builder"

type queryBuilderMysql struct {
	query_builder.InsertQueryBuilder
	query_builder.SelectQueryBuilder
	query_builder.UpdateQueryBuilder
}

func NewQueryBuilderMysql(insertQueryBuilder query_builder.InsertQueryBuilder, selectQueryBuilder query_builder.SelectQueryBuilder, updateQueryBuilder query_builder.UpdateQueryBuilder) *queryBuilderMysql {
	return &queryBuilderMysql{InsertQueryBuilder: insertQueryBuilder, SelectQueryBuilder: selectQueryBuilder, UpdateQueryBuilder: updateQueryBuilder}
}
