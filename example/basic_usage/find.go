package basic_usage

import (
	"fmt"
	"github.com/aksara-tech/aksarabase/example/model"
)

func (c repositoryUsingCompiler) Find(id int64) (res []model.Company, err error) {
	company := &model.Company{}
	i, err := c.db.Scanner.ScanStruct(company)
	if err != nil {
		return nil, err
	}

	i.Query.Where("id", id)

	err = c.db.Get(&res, i)
	if err != nil {
		return nil, err
	}

	return
}

func (c repositoryUsingCompiler) FindWith(id int64) (res []model.User, err error) {
	users := &model.User{
		Company: &model.Company{},
	}
	i, err := c.db.Scanner.ScanStruct(users)
	if err != nil {
		return nil, err
	}

	i.Query.WhereQuery = append(i.Query.WhereQuery, fmt.Sprintf("User.id=%v", id))
	err = c.db.Get(&res, i)
	if err != nil {
		return nil, err
	}

	return
}
