package orm

import (
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/constanta"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/info"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_builder"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_executor"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/scanner"
	"gitlab.com/aksaratech/aksarabase-go/v3/utils/reflector"
)

//QueryCompiler is a shortcut to build query, execute, and scan output.
type QueryCompiler interface {
	//Get parse query info n scan info, execute the query, then compile to struct destination or list of struct destination
	Get(dest interface{}) error
	//Update parse query info n scan info, execute the update query
	Update(dest interface{}, q info.QueryInfo, s info.ScanInfo) (interface{}, error)
	//Insert parse query info n scan info, execute the insert query
	Insert(dest interface{}, q info.QueryInfo, s info.ScanInfo) (interface{}, error)
}

type Compiler struct {
	QueryInfo     info.QueryInfo
	ScanInfo      info.ScanInfo
	scanner       scanner.Scanner
	queryBuilder  query_builder.QueryBuilder
	queryExecutor query_executor.SqlExecutor
}

func (c Compiler) Get(dest interface{}) error {
	query := c.queryBuilder.BuildSelectQuery(c.ScanInfo, c.QueryInfo)
	if reflector.GetKind(dest) == constanta.SLICE {
		rows, err := c.queryExecutor.Rows(query)
		if err != nil {
			return err
		}

		err = c.scanner.ToStructs(dest, rows, func() (string, interface{}) {
			return query, c.ScanInfo.StructAddress
		})
		if err != nil {
			return err
		}
	} else {
		row := c.queryExecutor.Row(query)
		err := c.scanner.ToStruct(dest, row)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c Compiler) Update(dest interface{}, q info.QueryInfo, s info.ScanInfo) (interface{}, error) {
	i, q, err := c.scanner.ScanStruct(dest)
	if err != nil {
		return nil, err
	}

	query := c.queryBuilder.BuildUpdateQuery(i, q)
	return c.queryExecutor.Exec(query)
}

func (c Compiler) Insert(dest interface{}, q info.QueryInfo, s info.ScanInfo) (interface{}, error) {
	i, q, err := c.scanner.ScanStruct(dest)
	if err != nil {
		return nil, err
	}

	query := c.queryBuilder.BuildInsertQuery(i, q)
	return c.queryExecutor.Exec(query)
}
