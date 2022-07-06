package main

import (
	aksarabase_v2 "aksarabase-v2"
	"aksarabase-v2/domain"
	"aksarabase-v2/domain/constanta"
	"aksarabase-v2/example/complex_usage/demo_repo"
	"aksarabase-v2/example/model"
	"aksarabase-v2/usecase/query_executor/debug_executor"
	"fmt"
)

func main() {
	adb := aksarabase_v2.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{
		Engine: domain.Engine{
			OutputScanner:      nil,
			InputScanner:       nil,
			PointerScanner:     nil,
			InsertQueryBuilder: nil,
			UpdateQueryBuilder: nil,
			SelectQueryBuilder: nil,
			SqlExecutor:        debug_executor.NewExecutor(constanta.MYSQL, constanta.DSN_TEST),
		},
	})
	repo := demo_repo.NewRepository(*adb)

	//============= FIRST ===============

	user, err := repo.GetUserByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	//============= FIND ===============
	companies, err := repo.FetchByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(companies)

	//============= STORE ===============
	phone := "089129430124"
	cstore := &model.Company{
		Name:     "CS STORE",
		PicName:  "wirnat cs",
		PicPhone: &phone,
	}

	err = repo.Store(cstore)

	fmt.Println(cstore)

}
