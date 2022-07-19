package basic_usage

import (
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
)

func (c repositoryUsingCompiler) CreateUser(user *model.User) error {
	return c.db.Insert(user)
}
