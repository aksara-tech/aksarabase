package demo_repo

import (
	"gitlab.com/wirawirw/aksarabase-go/v3/example/model"
	"gitlab.com/wirawirw/aksarabase-go/v3/utils/s"
)

func (r repository) FetchByID(id int64) (res []model.Company, err error) {
	companyFormer := model.Company{}
	i, q, err := r.db.Scanner.ScanStruct(&companyFormer)
	if err != nil {
		return nil, err
	}

	q.Where = append(q.Where, s.F("Company.id=%v", id))
	query := r.db.QBuilder.BuildSelectQuery(i, q)
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
