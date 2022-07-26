package complex_usage

import (
	"github.com/aksara-tech/aksarabase/example/model"
	"github.com/aksara-tech/aksarabase/utils/s"
)

//GetCompanyByID get single company by id
func (r repository) GetCompanyByID(id int64) (c model.Company, err error) {
	i, err := r.db.Scanner.ScanStruct(&c)
	if err != nil {
		return model.Company{}, err
	}

	i.Query.WhereQuery = append(i.Query.WhereQuery, s.F("%v", id))
	query := r.db.QBuilder.BuildSelect(i)

	row := r.db.Exec.Row(query)
	err = r.db.Scanner.ToStruct(&c, row)
	return
}
