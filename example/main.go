package main

import (
	"aksarabase-v2"
	"aksarabase-v2/usecase/query_executor/mysql"
)

func main() {
	aksarabase_v2.Open(mysql.NewMysqlDB())
}
