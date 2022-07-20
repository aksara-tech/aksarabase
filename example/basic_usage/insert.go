package basic_usage

import (
	"github.com/aksara-tech/aksarabase/example/model"
)

func (c repositoryUsingCompiler) CreateUser(user *model.User) error {
	return c.db.Insert(user)
}
