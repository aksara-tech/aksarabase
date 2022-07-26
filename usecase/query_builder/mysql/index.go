package mysql

import "github.com/aksara-tech/aksarabase/usecase/query_builder"

type queryBuilderMysql struct {
	query_builder.InsertQueryBuilder
	query_builder.SelectBuilder
	query_builder.UpdateQueryBuilder
}

func NewQueryBuilderMysql(insertQueryBuilder query_builder.InsertQueryBuilder, SelectBuilder query_builder.SelectBuilder, updateQueryBuilder query_builder.UpdateQueryBuilder) *queryBuilderMysql {
	return &queryBuilderMysql{InsertQueryBuilder: insertQueryBuilder, SelectBuilder: SelectBuilder, UpdateQueryBuilder: updateQueryBuilder}
}
