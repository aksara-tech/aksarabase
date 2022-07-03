package sql_driver

import (
	"database/sql"
	"fmt"
	"strings"
)

//NewSql open driver and database connection
func NewSql(driver string, dsn string) *sql.DB {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	fmt.Println("DRIVER IS RUNNING")
	fmt.Println("==DRIVER: " + strings.ToUpper(driver))

	return db
}
