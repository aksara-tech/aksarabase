package basic_usage

import aksarabase_v2 "gitlab.com/aksaratech/aksarabase-go/v3"

type repositoryUsingCompiler struct {
	db aksarabase_v2.ADB
}

func NewRepositoryBasic(db aksarabase_v2.ADB) *repositoryUsingCompiler {
	return &repositoryUsingCompiler{db: db}
}
