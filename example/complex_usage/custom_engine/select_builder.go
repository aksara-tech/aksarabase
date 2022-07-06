package custom_engine

import "aksarabase-v2/domain/info"

type selectBuilder struct {
}

func (s selectBuilder) BuildSelectQuery(info info.ScanInfo, qInfo info.QueryInfo) string {
	//TODO implement me
	panic("implemented select query builder")
}

func NewSelectBuilder() *selectBuilder {
	return &selectBuilder{}
}
