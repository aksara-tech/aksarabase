package basic_usage

import (
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	aksarabase_v2 "gitlab.com/aksaratech/aksarabase-go/v3"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/constanta"
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_executor/debug_executor"
	"testing"
)

func Test_repositoryUsingCompiler_Create(t *testing.T) {
	db := aksarabase_v2.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{
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
