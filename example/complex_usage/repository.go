package complex_usage

import "github.com/aksara-tech/aksarabase"

type repository struct {
	db aksarabase.ADB
}

func NewRepository(db aksarabase.ADB) *repository {
	return &repository{db: db}
}
