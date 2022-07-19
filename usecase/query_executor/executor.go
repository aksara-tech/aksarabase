package query_executor

import (
	"context"
	"database/sql"
)

type SqlExecutor interface {
	Exec(query string) (sql.Result, error)
	ExecWithCtx(ctx context.Context, query string) (interface{}, error)
	Row(query string) *sql.Row
	RowCtx(ctx context.Context, query string) *sql.Row
	Rows(query string) (*sql.Rows, error)
	RowsCtx(ctx context.Context, query string) (*sql.Rows, error)
	Close() error
}
