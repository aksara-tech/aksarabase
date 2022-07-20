package basic_usage

import (
	"github.com/aksara-tech/aksarabase"
	"github.com/aksara-tech/aksarabase/domain"
	"github.com/aksara-tech/aksarabase/domain/constanta"
	"github.com/aksara-tech/aksarabase/usecase/query_executor/debug_executor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repositoryUsingCompiler_Find(t *testing.T) {
	db := aksarabase.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{
		Engine: domain.Engine{
			SqlExecutor: debug_executor.NewExecutor(constanta.MYSQL, constanta.DSN_TEST),
		},
	})

	type fields struct {
		db aksarabase.ADB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "basic",
			fields: fields{
				db: *db,
			},
			args: args{
				id: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := repositoryUsingCompiler{
				db: tt.fields.db,
			}
			gotRes, err := c.Find(tt.args.id)
			assert.NotNil(t, gotRes)
			assert.Nil(t, err)
		})
	}
}
