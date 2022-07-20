package main

import (
	"fmt"
	"github.com/aksara-tech/aksarabase"
	"github.com/aksara-tech/aksarabase/domain"
	"github.com/aksara-tech/aksarabase/domain/constanta"
	"github.com/aksara-tech/aksarabase/example/basic_usage"
	"github.com/aksara-tech/aksarabase/example/complex_usage"
	"github.com/aksara-tech/aksarabase/example/model"
	"github.com/aksara-tech/aksarabase/usecase/query_executor/debug_executor"
)

func main() {

	adb := aksarabase.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{
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
