package complex_usage

import aksarabase_v2 "gitlab.com/aksaratech/aksarabase-go/v3"

type repository struct {
	db aksarabase_v2.ADB
}

func NewRepository(db aksarabase_v2.ADB) *repository {
	return &repository{db: db}
}
