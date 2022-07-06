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
	res, err := e.db.Exec(query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (e mysqlExecutor) ExecWithCtx(ctx context.Context, query string) (interface{}, error) {
	return e.db.ExecContext(ctx, query)
}

func (e mysqlExecutor) Row(query string) *sql.Row {
	s1 := time.Now()
	row := e.db.QueryRow(query)
	fmt.Printf("==Execute time: %v", time.Since(s1))
	return row
}

func (e mysqlExecutor) RowCtx(ctx context.Context, query string) *sql.Row {
	return e.db.QueryRowContext(ctx, query)
}

func (e mysqlExecutor) Rows(query string) (*sql.Rows, error) {
	return e.db.Query(query)
}

func (e mysqlExecutor) RowsCtx(ctx context.Context, query string) (*sql.Rows, error) {
	return e.db.QueryContext(ctx, query)
}
