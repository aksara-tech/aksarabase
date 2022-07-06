package debug_executor

import (
	"aksarabase-v2/utils/sql_driver"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type mysqlExecutor struct {
	db *sql.DB
}

func NewExecutor(driver string, dsn string) *mysqlExecutor {
	db := sql_driver.NewSql(driver, dsn)
	return &mysqlExecutor{
		db: db,
	}
}

func (e mysqlExecutor) Exec(query string) (interface{}, error) {
	s1 := time.Now()
	res, err := e.db.Exec(query)
	fmt.Printf("\n ==Query: %v ", query)
	fmt.Printf("\n ==Execute time: %v", time.Since(s1))
	if err != nil {
		fmt.Printf("\n ==Error: %v ", err)
		return nil, err
	}

	return res, nil
}

func (e mysqlExecutor) ExecWithCtx(ctx context.Context, query string) (interface{}, error) {
	s1 := time.Now()
	res, err := e.db.ExecContext(ctx, query)
	fmt.Printf("\n ==Query: %v ", query)
	fmt.Printf("\n ==Execute time: %v", time.Since(s1))
	if err != nil {
		fmt.Printf("\n ==Error: %v ", err)
		return nil, err
	}

	return res, nil
}

func (e mysqlExecutor) Row(query string) *sql.Row {
	s1 := time.Now()
	row := e.db.QueryRow(query)
	fmt.Printf("\n ==Query: %v ", query)
	fmt.Printf("\n ==Execute time: %v", time.Since(s1))
	if err := row.Err(); err != nil {
		fmt.Printf("\n ==Error: %v ", err)
	}
	return row
}

func (e mysqlExecutor) RowCtx(ctx context.Context, query string) *sql.Row {
	s1 := time.Now()
	row := e.db.QueryRowContext(ctx, query)
	fmt.Printf("\n ==Query: %v ", query)
	fmt.Printf("\n ==Execute time: %v", time.Since(s1))
	if err := row.Err(); err != nil {
		fmt.Printf("\n ==Error: %v ", err)
	}
	return row
}

func (e mysqlExecutor) Rows(query string) (*sql.Rows, error) {
	s1 := time.Now()
	row, err := e.db.Query(query)
	fmt.Printf("\n ==Query: %v ", query)
	fmt.Printf("\n ==Execute time: %v", time.Since(s1))
	if err != nil {
		fmt.Printf("\n ==Error: %v ", err)
		return nil, err
	}

	return row, nil

}

func (e mysqlExecutor) RowsCtx(ctx context.Context, query string) (*sql.Rows, error) {
	s1 := time.Now()
	row, err := e.db.QueryContext(ctx, query)
	fmt.Printf("\n ==Query: %v ", query)
	fmt.Printf("\n ==Execute time: %v", time.Since(s1))
	if err != nil {
		fmt.Printf("\n ==Error: %v ", err)
		return nil, err
	}

	return row, nil
}
