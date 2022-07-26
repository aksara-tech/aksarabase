package complex_usage

import (
	"github.com/aksara-tech/aksarabase/example/model"
	"github.com/aksara-tech/aksarabase/utils/s"
)

func (r repository) Update(company *model.Company) error {
	i, err := r.db.Scanner.ScanStruct(company)
	if err != nil {
		return err
	}

	i.Query.WhereQuery = append(i.Query.WhereQuery, s.F("id=%v", company.ID))
	query := r.db.QBuilder.BuildUpdateQuery(i)
	_, err = r.db.Exec.Exec(query)
	return err
}
