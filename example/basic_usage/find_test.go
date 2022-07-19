package basic_usage

import (
	"github.com/stretchr/testify/assert"
	aksarabase_v2 "gitlab.com/aksaratech/aksarabase-go/v3"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/constanta"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_executor/debug_executor"
	"testing"
)

func Test_repositoryUsingCompiler_Find(t *testing.T) {
	db := aksarabase_v2.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{
		Engine: domain.Engine{
			SqlExecutor: debug_executor.NewExecutor(constanta.MYSQL, constanta.DSN_TEST),
		},
	})

	type fields struct {
		db aksarabase_v2.ADB
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
