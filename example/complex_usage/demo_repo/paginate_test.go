package demo_repo

import (
	aksarabase_v2 "aksarabase-v2"
	"aksarabase-v2/domain"
	"aksarabase-v2/domain/constanta"
	"aksarabase-v2/usecase/query_executor/debug_executor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repository_PaginateUser(t *testing.T) {
	t.Run("Paginate test", func(t *testing.T) {
		adb := aksarabase_v2.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{
			Engine: domain.Engine{
				SqlExecutor: debug_executor.NewExecutor(constanta.MYSQL, constanta.DSN_TEST),
			},
		})
		repo := NewRepository(*adb)
		u, err := repo.PaginateUser(Paginator{Limit: 1, Page: 1})
		assert.Nil(t, err)
		assert.NotNil(t, u)
	})
}
