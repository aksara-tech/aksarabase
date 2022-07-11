package complex_usage

import (
	"github.com/stretchr/testify/assert"
	aksarabase_v2 "gitlab.com/aksaratech/aksarabase-go/v3"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/constanta"
	"testing"
)

func Test_repository_GetUserByID(t *testing.T) {
	t.Run("Basic test", func(t *testing.T) {
		adb := aksarabase_v2.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{})
		repo := NewRepository(*adb)
		u, err := repo.GetUserByID(1)
		assert.Nil(t, err)
		assert.NotNil(t, u)
	})
}
