package mysql

import (
	"aksarabase-v2/domain/constanta"
	"aksarabase-v2/example/model"
	"aksarabase-v2/usecase/scanner"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mysqlDB_Insert(t *testing.T) {
	type args struct {
		dest interface{}
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{name: "Basic", args: args{
			dest: &model.Company{
				Name:     "Wirajus",
				PicName:  "08913",
				PicPhone: nil,
			},
		}, err: nil},
	}

	db := NewMysqlDB(constanta.MYSQL, constanta.DSN_TEST, scanner.NewOutputScanner(scanner.NewPointerScanner()), scanner.NewInputScanner())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.Insert(context.Background(), tt.args.dest)
			assert.Equal(t, tt.err, err)
		})
	}
}
