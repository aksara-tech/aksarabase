package demo_repo

import (
	"aksarabase-v2/example/model"
	"aksarabase-v2/utils/s"
)

//GetCompanyByID get single company by id
func (r repository) GetCompanyByID(id int64) (c model.Company, err error) {
	i, q, err := r.db.Scanner.ScanStruct(&c)
	if err != nil {
		return model.Company{}, err
	}

	q.Where = append(q.Where, s.F("%v", id))
	query := r.db.QBuilder.BuildSelectQuery(i, q)

	row := r.db.Exec.Row(query)
	err = r.db.Scanner.ToStruct(&c, row)
	return
}