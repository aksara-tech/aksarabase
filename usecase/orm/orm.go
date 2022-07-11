package orm

import (
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/info"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_builder"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_executor"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/scanner"
)

type ORM interface {
	OpenCompiler(i info.ScanInfo, q info.QueryInfo) QueryCompiler
}

type orm struct {
	s scanner.Scanner
	e query_executor.SqlExecutor
	q query_builder.QueryBuilder
}

func NewOrm(s scanner.Scanner, e query_executor.SqlExecutor, q query_builder.QueryBuilder) *orm {
	return &orm{s: s, e: e, q: q}
}

func (o orm) OpenCompiler(i info.ScanInfo, q info.QueryInfo) QueryCompiler {
	return &Compiler{
		QueryInfo:     q,
		ScanInfo:      i,
		scanner:       o.s,
		queryBuilder:  o.q,
		queryExecutor: o.e,
	}
}
