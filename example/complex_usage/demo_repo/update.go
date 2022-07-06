package demo_repo

import (
	"gitlab.com/wirawirw/aksarabase-go/v3/example/model"
	"gitlab.com/wirawirw/aksarabase-go/v3/utils/s"
)

func (r repository) Update(company *model.Company) error {
	i, q, err := r.db.Scanner.ScanStruct(company)
	if err != nil {
		return err
	}

	q.Where = append(q.Where, s.F("id=%v", company.ID))
	query := r.db.QBuilder.BuildUpdateQuery(i, q)
	_, err = r.db.Exec.Exec(query)
	return err
}
