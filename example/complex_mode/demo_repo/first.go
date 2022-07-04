package demo_repo

import (
	"aksarabase-v2/example/model"
	"fmt"
)

func (r repository) GetByID(id int64) (c model.Company, err error) {
	i, err := r.db.ADV.Scanner.ScanStruct(&c)
	if err != nil {
		return model.Company{}, err
	}

	selectQ := r.db.ADV.QBuilder.BuildSelectQuery(i, fmt.Sprintf("where id=%v", id))

	row := r.db.Exec.Row(selectQ)
	err = r.db.ADV.Scanner.ToStruct(&c, row)
	return
}
