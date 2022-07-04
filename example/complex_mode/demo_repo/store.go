package demo_repo

import (
	"aksarabase-v2/example/model"
	"fmt"
)

func (r repository) Store(company *model.Company) error {
	fmt.Println(company)
	i, err := r.db.Scanner.ScanStruct(company)
	if err != nil {
		return err
	}

	query := r.db.QBuilder.BuildInsertQuery(i)
	res, err := r.db.Exec.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}
