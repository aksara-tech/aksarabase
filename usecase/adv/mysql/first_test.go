package mysql

import (
	"aksarabase-v2/domain/constanta"
	"aksarabase-v2/example/model"
	"aksarabase-v2/usecase/scanner"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mysqlDB_First(t *testing.T) {
	d := new(model.Company)
	mysql := NewMysqlDB(constanta.MYSQL, constanta.DSN_TEST, scanner.NewOutputScanner(scanner.NewPointerScanner()), scanner.NewInputScanner())
	err := mysql.First(context.Background(), d, "select * from companies where id=1")
	assert.Nil(t, err)
}
