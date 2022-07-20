package complex_usage

import (
	"github.com/aksara-tech/aksarabase"
	"github.com/aksara-tech/aksarabase/domain"
	"github.com/aksara-tech/aksarabase/domain/constanta"
	"github.com/aksara-tech/aksarabase/example/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repository_Update(t *testing.T) {
	t.Run("Basic test", func(t *testing.T) {
		adb := aksarabase.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{})
		repo := NewRepository(*adb)
		company := model.Company{
			Name: "cc23",
		}
		err := repo.Update(&company)
		assert.Nil(t, err)
	})
}
