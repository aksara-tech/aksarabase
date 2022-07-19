package main

import (
	"fmt"
	aksarabase_v2 "gitlab.com/aksaratech/aksarabase-go/v3"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/constanta"
	"gitlab.com/aksaratech/aksarabase-go/v3/example/basic_usage"
	"gitlab.com/aksaratech/aksarabase-go/v3/example/complex_usage"
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/query_executor/debug_executor"
)

func main() {

	adb := aksarabase_v2.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{
		Engine: domain.Engine{
			OutputScanner:      nil,
			InputScanner:       nil,
			PointerScanner:     nil,
			InsertQueryBuilder: nil,
			UpdateQueryBuilder: nil,
			SelectBuilder:      nil,
			SqlExecutor:        debug_executor.NewExecutor(constanta.MYSQL, constanta.DSN_TEST),
		},
	})
	repo := complex_usage.NewRepository(*adb)
	repoB := basic_usage.NewRepositoryBasic(*adb)
	//============= FIRST ===============

	user, err := repo.GetUserByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	users, err := repoB.Find(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	usersWithCompany, err := repoB.FindWith(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(usersWithCompany)

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
