package basic_usage

import (
	"github.com/aksara-tech/aksarabase"
	"github.com/aksara-tech/aksarabase/domain"
	"github.com/aksara-tech/aksarabase/domain/constanta"
	"github.com/aksara-tech/aksarabase/example/model"
	"github.com/aksara-tech/aksarabase/usecase/query_executor/debug_executor"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repositoryUsingCompiler_Create(t *testing.T) {
	db := aksarabase.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{
		Engine: domain.Engine{
			SqlExecutor: debug_executor.NewExecutor(constanta.MYSQL, constanta.DSN_TEST),
		},
	})
	c := repositoryUsingCompiler{
		db: *db,
	}
	user := model.User{
		BaseModel: model.BaseModel{},
		Name:      "Wirax22",
		Email:     "wirajus@gmail.com",
		CompanyID: 1,
		Company:   nil,
	}

	faker.FakeData(&user)

	err := c.CreateUser(&user)
	assert.Nil(t, err)
}
