package basic_usage

import (
	"github.com/aksara-tech/aksarabase"
	"github.com/aksara-tech/aksarabase/domain"
	"github.com/aksara-tech/aksarabase/domain/constanta"
	"github.com/aksara-tech/aksarabase/usecase/query_executor/debug_executor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repositoryUsingCompiler_First(t *testing.T) {
	db := aksarabase.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{
		Engine: domain.Engine{
			SqlExecutor: debug_executor.NewExecutor(constanta.MYSQL, constanta.DSN_TEST),
		},
	})
	c := repositoryUsingCompiler{
		db: *db,
	}
	gotRes, err := c.FirstWith(1)
	assert.NotNil(t, gotRes)
	assert.Nil(t, err)
}
