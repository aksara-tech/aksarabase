package complex_usage

import (
	"github.com/aksara-tech/aksarabase"
	"github.com/aksara-tech/aksarabase/domain"
	"github.com/aksara-tech/aksarabase/domain/constanta"
	"github.com/aksara-tech/aksarabase/usecase/query_executor/debug_executor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repository_PaginateUser(t *testing.T) {
	t.Run("Paginate test", func(t *testing.T) {
		adb := aksarabase.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{
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
