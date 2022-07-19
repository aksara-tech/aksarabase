package complex_usage

import (
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
	"gitlab.com/aksaratech/aksarabase-go/v3/utils/s"
)

//GetUserWithCompany get user and his company
func (r repository) GetUserWithCompany(id int64) (u model.User, err error) {
	u = model.User{
		Company: &model.Company{},
	}

	i, err := r.db.Scanner.ScanStruct(&u)
	if err != nil {
		return model.User{}, err
	}

	query := r.db.QBuilder.BuildSelect(i)
	row := r.db.Exec.Row(query)
	err = r.db.Scanner.ToStruct(&u, row)
	return
}

//GetUserByID get single user
func (r repository) GetUserByID(id int64) (u model.User, err error) {
	u = model.User{
		Company: &model.Company{},
	}
	i, err := r.db.Scanner.ScanStruct(&u)
	if err != nil {
		return model.User{}, err
	}

	i.Query.WhereQuery = append(i.Query.WhereQuery, s.F("User.id=%v", id))

	query := r.db.QBuilder.BuildSelect(i)
	row := r.db.Exec.Row(query)

	err = r.db.Scanner.ToStruct(&u, row)
	return
}
