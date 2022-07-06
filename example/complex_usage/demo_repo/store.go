package demo_repo

import (
	"aksarabase-v2/example/model"
	"fmt"
	"time"
)

func (r repository) Store(company *model.Company) error {
	fmt.Println(company)
	company.CreatedAt = time.Now()

	i, q, err := r.db.Scanner.ScanStruct(company)
	if err != nil {
		return err
	}

	query := r.db.QBuilder.BuildInsertQuery(i, q)
	res, err := r.db.Exec.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}
