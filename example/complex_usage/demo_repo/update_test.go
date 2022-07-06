package demo_repo

import (
	aksarabase_v2 "aksarabase-v2"
	"aksarabase-v2/domain"
	"aksarabase-v2/domain/constanta"
	"aksarabase-v2/example/model"
	"github.com/stretchr/testify/assert"
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
