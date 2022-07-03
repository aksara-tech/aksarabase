package mysql

import (
	"aksarabase-v2/usecase/scanner"
	"aksarabase-v2/utils/sql_driver"
	"database/sql"
)

type mysqlDB struct {
	db            *sql.DB
	outputScanner scanner.OutputScanner
}

func NewMysqlDB(driver string, dsn string) *mysqlDB {
	db := sql_driver.NewSql(driver, dsn)
	return &mysqlDB{
		db:            db,
		outputScanner: scanner.NewOutputScanner(),
	}
}
