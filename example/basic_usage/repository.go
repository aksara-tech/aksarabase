package basic_usage

import "github.com/aksara-tech/aksarabase"

type repositoryUsingCompiler struct {
	db aksarabase.ADB
}

func NewRepositoryBasic(db aksarabase.ADB) *repositoryUsingCompiler {
	return &repositoryUsingCompiler{db: db}
}
