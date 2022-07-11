package complex_usage

import (
	"github.com/stretchr/testify/assert"
	aksarabase_v2 "gitlab.com/aksaratech/aksarabase-go/v3"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/constanta"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_executor/debug_executor"
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
