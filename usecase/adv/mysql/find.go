package mysql

import (
	"aksarabase-v2/domain/callbacks"
	"context"
)

func (m mysqlDB) Find(ctx context.Context, dest interface{}, form callbacks.StructForm) error {
	query, _ := form()

	rows, err := m.db.Query(query)
	if err != nil {
		return err
	}

	err = m.outputScanner.ToStructs(dest, rows, form)
	if err != nil {
		return err
	}

	return nil
}
