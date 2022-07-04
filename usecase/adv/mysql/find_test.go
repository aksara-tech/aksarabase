package mysql

import (
	"aksarabase-v2/domain/constanta"
	"aksarabase-v2/example/model"
	"aksarabase-v2/usecase/scanner"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mysqlDB_Find(t *testing.T) {
	t.Run("Basic Test", func(t *testing.T) {
		d := new([]model.Company)
		mysql := NewMysqlDB(constanta.MYSQL, constanta.DSN_TEST, scanner.NewOutputScanner(scanner.NewPointerScanner()), scanner.NewInputScanner())
		err := mysql.Find(context.Background(), d, func() (query string, former interface{}) {
			return "select * from companies", &model.Company{}
		})
		assert.Nil(t, err)
	})
}
