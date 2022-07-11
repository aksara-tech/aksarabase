package complex_usage

import (
	"github.com/stretchr/testify/assert"
	aksarabase_v2 "gitlab.com/aksaratech/aksarabase-go/v3"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/constanta"
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
	"testing"
)

func Test_repository_Update(t *testing.T) {
	t.Run("Basic test", func(t *testing.T) {
		adb := aksarabase_v2.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{})
		repo := NewRepository(*adb)
		company := model.Company{
			Name: "cc23",
		}
		err := repo.Update(&company)
		assert.Nil(t, err)
	})
}
