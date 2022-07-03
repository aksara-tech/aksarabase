package scanner

import "database/sql"

type outputScanner struct {
	pointerScanner PointerScanner
}

func NewOutputScanner() *outputScanner {
	return &outputScanner{}
}

func (o outputScanner) ToStructs(dest interface{}, row *sql.Rows) error {

	panic("implement me")
}

func (o outputScanner) ToStruct(dest interface{}, row *sql.Row) error {
	var listAddr []interface{}
	o.pointerScanner.GetListPointer(dest, &listAddr)
	return row.Scan(listAddr...)
}
