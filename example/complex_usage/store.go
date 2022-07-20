package complex_usage

import (
	"github.com/aksara-tech/aksarabase/example/model"
	"time"
)

func (r repository) Store(company *model.Company) error {
	company.CreatedAt = time.Now()

	i, err := r.db.Scanner.ScanStruct(company)
	if err != nil {
		return err
	}

	query := r.db.QBuilder.BuildInsertQuery(i)
	_, err = r.db.Exec.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
