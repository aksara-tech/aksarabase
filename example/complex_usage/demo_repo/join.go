package demo_repo

import (
	"gitlab.com/wirawirw/aksarabase-go/v3/example/model"
	"gitlab.com/wirawirw/aksarabase-go/v3/utils/s"
)

//GetUserWithCompany get user and his company
func (r repository) GetUserWithCompany(id int64) (u model.User, err error) {
	u = model.User{
		Company: &model.Company{},
	}

	i, q, err := r.db.Scanner.ScanStruct(&u)
	if err != nil {
		return model.User{}, err
	}

	query := r.db.QBuilder.BuildSelectQuery(i, q)
	row := r.db.Exec.Row(query)
	err = r.db.Scanner.ToStruct(&u, row)
	return
}

//GetUserByID get single user
func (r repository) GetUserByID(id int64) (u model.User, err error) {
	u = model.User{
		Company: &model.Company{},
	}
	i, q, err := r.db.Scanner.ScanStruct(&u)
	if err != nil {
		return model.User{}, err
	}

	q.Where = append(q.Where, s.F("User.id=%v", id))

	query := r.db.QBuilder.BuildSelectQuery(i, q)
	row := r.db.Exec.Row(query)

	err = r.db.Scanner.ToStruct(&u, row)
	return
}
