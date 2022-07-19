package scanner

import (
	"database/sql"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/callbacks"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/info"
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
	//ScanStruct generate struct info and modifiable query info.
	//  This also auto generate join query if struct is nested
	ScanStruct(dest interface{}) (info.Info, error)
}

type PointerScanner interface {
	//GetListPointer collect list pointer field from given dest
	GetListPointer(src interface{}, columns *[]interface{})
}
