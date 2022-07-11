package info

import "gitlab.com/aksaratech/aksarabase-go/v3/domain/query"

//QueryInfo is a partial query builder that has modifiable structure before it's executed
type QueryInfo struct {
	Select  []string
	From    string
	Where   []string
	Join    []query.JoinRelation
	Limit   string
	OrderBy string
}
