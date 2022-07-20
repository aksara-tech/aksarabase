package custom_engine

import "github.com/aksara-tech/aksarabase/domain/info"

type inputScanner struct {
}

func NewInputScanner() *inputScanner {
	return &inputScanner{}
}

func (i inputScanner) ScanStruct(dest interface{}) (info.ScanInfo, info.QueryInfo, error) {
	//TODO implement me
	panic("implemented")
}
