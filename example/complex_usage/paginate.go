package complex_usage

import (
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
	"gitlab.com/aksaratech/aksarabase-go/v3/utils/s"
)

type Paginator struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (r repository) PaginateUser(paginator Paginator) (res []model.Company, err error) {
	former := model.Company{}
	i, q, err := r.db.Scanner.ScanStruct(&former)
	if err != nil {
		return nil, err
	}
	q.Limit = s.F("%v,%v", paginator.Page, paginator.Limit)
	query := r.db.QBuilder.BuildSelectQuery(i, q)
	rows, err := r.db.Exec.Rows(query)
	if err != nil {
		return nil, err
	}

	err = r.db.Scanner.ToStructs(&res, rows, func() (string, interface{}) {
		return query, &former
	})
	if err != nil {
		return nil, err
	}

	return
}
