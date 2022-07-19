package domain

import (
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_builder"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_executor"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/scanner"
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
	//BuildSelect
	SelectBuilder query_builder.SelectBuilder
	//SqlExecutor execute query or hooking executor
	SqlExecutor query_executor.SqlExecutor
}
