package domain

import (
	"aksarabase-v2/usecase/query_builder"
	"aksarabase-v2/usecase/query_executor"
	"aksarabase-v2/usecase/scanner"
)

type Engine struct {
	//OutputScanner Default use is reflect_scanner
	OutputScanner scanner.OutputScanner
	//InputScanner Default use is reflect_scanner
	InputScanner scanner.InputScanner
	//PointerScanner Default use is reflect_scanner
	PointerScanner scanner.PointerScanner
	//InsertQueryBuilder
	InsertQueryBuilder query_builder.InsertQueryBuilder
	//UpdateQueryBuilder
	UpdateQueryBuilder query_builder.UpdateQueryBuilder
	//BuildSelectQuery
	SelectQueryBuilder query_builder.SelectQueryBuilder
	//SqlExecutor execute query or hooking executor
	SqlExecutor query_executor.SqlExecutor
}
