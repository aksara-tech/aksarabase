package complex_usage

import (
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
	"gitlab.com/aksaratech/aksarabase-go/v3/utils/s"
)

func (r repository) FetchByID(id int64) (res []model.Company, err error) {
	companyFormer := model.Company{}
	i, err := r.db.Scanner.ScanStruct(&companyFormer)
	if err != nil {
		return nil, err
	}

	i.Query.WhereQuery = append(i.Query.WhereQuery, s.F("Company.id=%v", id))
	query := r.db.QBuilder.BuildSelect(i)
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
