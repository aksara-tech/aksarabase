package basic_usage

import (
	"fmt"
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
)

func (c repositoryUsingCompiler) Find(id int64) (res []model.Company, err error) {
	company := &model.Company{}
	i, q, err := c.db.Scanner.ScanStruct(company)
	if err != nil {
		return nil, err
	}

	q.Where = append(q.Where, fmt.Sprintf("id=%v", id))

	compiler := c.db.ORM.OpenCompiler(i, q)
	err = compiler.Get(&res)
	if err != nil {
		return nil, err
	}

	return
}

func (c repositoryUsingCompiler) FindWith(id int64) (res []model.User, err error) {
	users := &model.User{
		Company: &model.Company{},
	}
	i, q, err := c.db.Scanner.ScanStruct(users)
	if err != nil {
		return nil, err
	}

	q.Where = append(q.Where, fmt.Sprintf("User.id=%v", id))
	compiler := c.db.ORM.OpenCompiler(i, q)
	err = compiler.Get(&res)
	if err != nil {
		return nil, err
	}

	return
}
