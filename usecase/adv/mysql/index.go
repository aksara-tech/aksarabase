package mysql

import (
	"aksarabase-v2/usecase/scanner"
	"aksarabase-v2/utils/sql_driver"
	"database/sql"
)

type mysqlDB struct {
	db            *sql.DB
	outputScanner scanner.OutputScanner
	inputScanner  scanner.InputScanner
}

func NewMysqlDB(driver string, dsn string, o scanner.OutputScanner, i scanner.InputScanner) *mysqlDB {
	db := sql_driver.NewSql(driver, dsn)
	return &mysqlDB{
		db:            db,
		outputScanner: o,
		inputScanner:  i,
	}
}
