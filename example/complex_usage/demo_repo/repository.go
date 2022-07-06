package demo_repo

import aksarabase_v2 "aksarabase-v2"

type repository struct {
	db aksarabase_v2.ADB
}

func NewRepository(db aksarabase_v2.ADB) *repository {
	return &repository{db: db}
}
