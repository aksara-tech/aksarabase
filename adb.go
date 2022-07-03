package aksarabase_v2

import "aksarabase-v2/usecase/query_executor"

//ADB Aksara Database Definition
type ADB struct {
	DB query_executor.DB
}

//Open used to integrate all module and start the orm
func Open(db query_executor.DB) *ADB {
	return &ADB{
		DB: db,
	}
}
