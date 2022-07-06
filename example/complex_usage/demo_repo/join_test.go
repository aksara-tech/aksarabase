package demo_repo

import (
	aksarabase_v2 "aksarabase-v2"
	"aksarabase-v2/domain"
	"aksarabase-v2/domain/constanta"
	"github.com/stretchr/testify/assert"
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