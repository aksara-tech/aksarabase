package compiler

import (
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/constanta"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/info"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_builder"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_executor"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/scanner"
	"gitlab.com/aksaratech/aksarabase-go/v3/utils/reflector"
	"reflect"
	"time"
)

//QueryCompiler is a shortcut to build query, execute, and scan output.
type Compiler interface {
	//Get parse query info n scan info, execute the query, then compile to struct destination or list of struct destination
	Get(dest interface{}, info info.Info) error
	//Update parse query info n scan info, execute the update query
	Update(dest interface{}, info info.Info) error
	//Insert parse query info n scan info, execute the insert query
	Insert(dest interface{}) error
	//Pagion
}

type compiler struct {
	scanner       scanner.Scanner
	queryBuilder  query_builder.QueryBuilder
	queryExecutor query_executor.SqlExecutor
}

func NewCompiler(scanner scanner.Scanner, queryBuilder query_builder.QueryBuilder, queryExecutor query_executor.SqlExecutor) *compiler {
	return &compiler{scanner: scanner, queryBuilder: queryBuilder, queryExecutor: queryExecutor}
}

func (c compiler) Get(dest interface{}, info info.Info) error {
	query := c.queryBuilder.BuildSelect(info)
	if reflector.GetKind(dest) == constanta.SLICE {
		rows, err := c.queryExecutor.Rows(query)
		if err != nil {
			return err
		}

		err = c.scanner.ToStructs(dest, rows, func() (string, interface{}) {
			return query, info.Scan.StructAddress
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

func (c compiler) Update(dest interface{}, info info.Info) error {
	i, err := c.scanner.ScanStruct(dest)
	if err != nil {
		return err
	}

	query := c.queryBuilder.BuildUpdateQuery(i)
	_, err = c.queryExecutor.Exec(query)
	return err
}

func (c compiler) Insert(dest interface{}) error {
	i, err := c.scanner.ScanStruct(dest)
	if err != nil {
		return err
	}

	query := c.queryBuilder.BuildInsertQuery(i)
	res, err := c.queryExecutor.Exec(query)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	//set dest
	val := reflect.ValueOf(dest)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	vID := val.FieldByName("ID")
	if vID.CanSet() {
		vID.SetInt(lastID)
	}
	vid := val.FieldByName("id")
	if vid.CanSet() {
		vid.SetInt(lastID)
	}
	vCreatedAt := val.FieldByName("CreatedAt")
	if vCreatedAt.CanSet() {
		now := time.Now()
		vCreatedAt.Set(reflect.ValueOf(now))
	}

	return err

}
