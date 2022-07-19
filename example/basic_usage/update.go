package basic_usage

import "gitlab.com/aksaratech/aksarabase-go/v3/example/model"

func (c repositoryUsingCompiler) UpdateUser(user *model.User) error {
	i, err := c.db.Scanner.ScanStruct(user)
	if err != nil {
		return err
	}

	return c.db.Update(user, i)
}
