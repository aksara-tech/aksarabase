package default_executor

import (
	"context"
	"database/sql"
	"gitlab.com/aksaratech/aksarabase-go/v3/utils/sql_driver"
)

type mysqlExecutor struct {
	db *sql.DB
}

func NewExecutor(driver string, dsn string) *mysqlExecutor {
	db := sql_driver.NewSql(driver, dsn)
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	return &mysqlExecutor{
		db: db,
	}
}

func (e mysqlExecutor) Exec(query string) (sql.Result, error) {
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
	return e.db.QueryRow(query)
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

func (e mysqlExecutor) Close() error {
	return e.db.Close()
}
