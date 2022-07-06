package scanner

import (
	"aksarabase-v2/domain/callbacks"
	"aksarabase-v2/domain/info"
	"database/sql"
)

type Scanner interface {
	OutputScanner
	InputScanner
	PointerScanner
}

type OutputScanner interface {
	//ToStruct fill given destination using single row result
	ToStruct(dest interface{}, row *sql.Row) error
	//ToStructs fill given destination using multi row result
	ToStructs(dest interface{}, rows *sql.Rows, m callbacks.StructForm) error
}

type InputScanner interface {
	//ScanStruct get struct info
	ScanStruct(dest interface{}) (info.ScanInfo, info.QueryInfo, error)
}

type PointerScanner interface {
	//GetListPointer collect list pointer field from given dest
	GetListPointer(src interface{}, columns *[]interface{})
}
