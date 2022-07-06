package custom_engine

import "gitlab.com/wirawirw/aksarabase-go/v3/domain/info"

type selectBuilder struct {
}

func (s selectBuilder) BuildSelectQuery(info info.ScanInfo, qInfo info.QueryInfo) string {
	//TODO implement me
	panic("implemented select query builder")
}

func NewSelectBuilder() *selectBuilder {
	return &selectBuilder{}
}
