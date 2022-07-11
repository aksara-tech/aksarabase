package reflect_scanner

import (
	"database/sql"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/callbacks"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/scanner"
	"reflect"
)

type outputScanner struct {
	pointerScanner scanner.PointerScanner
}

func NewOutputScanner(pointerScanner scanner.PointerScanner) *outputScanner {
	return &outputScanner{pointerScanner: pointerScanner}
}

func (o outputScanner) ToStructs(dest interface{}, rows *sql.Rows, m callbacks.StructForm) error {
	val := reflect.ValueOf(dest)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for rows.Next() {
		//generate new struct n refresh nested pointer from hooker function
		//create new memory of typ struct
		_, x := m()

		d := reflect.ValueOf(x)
		//build list address for scan
		columns := new([]interface{})
		o.pointerScanner.GetListPointer(d.Interface(), columns)

		//set list address with result data
		err := rows.Scan(*columns...)
		if err != nil {
			return err
		}

		//append dest with
		if val.CanSet() {
			val.Set(reflect.Append(val, d.Elem()))
		}
	}

	err := rows.Close()
	if err != nil {
		return err
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

func (o outputScanner) ToStruct(dest interface{}, row *sql.Row) error {
	var listAddr []interface{}
	o.pointerScanner.GetListPointer(dest, &listAddr)
	return row.Scan(listAddr...)
}
