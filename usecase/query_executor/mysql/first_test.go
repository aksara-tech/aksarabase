package mysql

import (
	"aksarabase-v2/domain/constanta"
	"context"
	"testing"
)

type User struct {
}

func Test_mysqlDB_First(t *testing.T) {
	mysql := NewMysqlDB(constanta.MYSQL, constanta.DSN_TEST)
	mysql.First(context.Background(), "", "")
}
