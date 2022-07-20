package complex_usage

import (
	"github.com/aksara-tech/aksarabase"
	"github.com/aksara-tech/aksarabase/domain"
	"github.com/aksara-tech/aksarabase/domain/constanta"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repository_GetUserByID(t *testing.T) {
	t.Run("Basic test", func(t *testing.T) {
		adb := aksarabase.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{})
		repo := NewRepository(*adb)
		u, err := repo.GetUserByID(1)
		assert.Nil(t, err)
		assert.NotNil(t, u)
	})
}
