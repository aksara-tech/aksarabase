package custom_engine

import "gitlab.com/aksaratech/aksarabase-go/v3/domain/info"

type SelectBuilder struct {
}

func (s SelectBuilder) BuildSelect(info info.ScanInfo, qInfo info.QueryInfo) string {
	//TODO implement me
	panic("implemented SelectQuery query builder")
}

func NewSelectBuilder() *SelectBuilder {
	return &SelectBuilder{}
}
