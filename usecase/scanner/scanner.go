package scanner

import "database/sql"

type Scanner interface {
	OutputScanner
	InputScanner
}

type OutputScanner interface {
	//ToStruct fill given destination using single row result
	ToStruct(dest interface{}, row *sql.Row) error
	//ToStructs fill given destination using multi row result
	ToStructs(dest interface{}, row *sql.Rows) error
}

type InputScanner interface {
	//GetStruct get struct info
	GetStruct(dest interface{}) (ScanInfo, error)
}

type PointerScanner interface {
	//GetListPointer collect list pointer from dest
	GetListPointer(src interface{}, columns *[]interface{})
}
