package demo_repo

import (
	"aksarabase-v2/example/model"
	"fmt"
)

func (r repository) FetchByID(id int64) (res []model.Company, err error) {
	companyFormer := model.Company{}
	i, err := r.db.Scanner.ScanStruct(&companyFormer)
	if err != nil {
		return nil, err
	}

	query := r.db.QBuilder.BuildSelectQuery(i, fmt.Sprintf("where id=%v", id))
	rows, err := r.db.Exec.Rows(query)
	if err != nil {
		return nil, err
	}

	err = r.db.Scanner.ToStructs(&res, rows, func() (query string, former interface{}) {
		return query, &companyFormer
	})
	if err != nil {
		return nil, err
	}

	return
}
