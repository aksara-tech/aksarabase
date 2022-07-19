package basic_usage

import "gitlab.com/aksaratech/aksarabase-go/v3/example/model"

func (c repositoryUsingCompiler) FirstWith(id int64) (res model.User, err error) {
	user := model.User{
		Company: &model.Company{},
	}
	i, err := c.db.Scanner.ScanStruct(&user)
	if err != nil {
		return model.User{}, err
	}

	i.Query.Where("User.id", id)

	err = c.db.Get(&user, i)
	res = user
	return
}
