package mysql

import "context"

func (m mysqlDB) First(ctx context.Context, dest interface{}, query string) error {
	//exec query
	row := m.db.QueryRow(query)
	//stick the result to destination pointer
	err := m.outputScanner.ToStruct(dest, row)
	if err != nil {
		return err
	}
	return nil
}
