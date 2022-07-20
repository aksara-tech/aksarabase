package complex_usage

import (
	"github.com/aksara-tech/aksarabase/example/model"
	"github.com/aksara-tech/aksarabase/utils/s"
)

type Paginator struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (r repository) PaginateUser(paginator Paginator) (res []model.Company, err error) {
	former := model.Company{}
	i, err := r.db.Scanner.ScanStruct(&former)
	if err != nil {
		return nil, err
	}
	i.Query.LimitQuery = s.F("%v,%v", paginator.Page, paginator.Limit)
	query := r.db.QBuilder.BuildSelect(i)
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
