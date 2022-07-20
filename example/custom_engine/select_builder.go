package custom_engine

import "github.com/aksara-tech/aksarabase/domain/info"

type SelectBuilder struct {
}

func (s SelectBuilder) BuildSelect(info info.ScanInfo, qInfo info.QueryInfo) string {
	//TODO implement me
	panic("implemented SelectQuery query builder")
}

func NewSelectBuilder() *SelectBuilder {
	return &SelectBuilder{}
}
