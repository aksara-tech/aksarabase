package main

import (
	aksarabase_v2 "aksarabase-v2"
	"aksarabase-v2/domain"
	"aksarabase-v2/domain/constanta"
	"aksarabase-v2/example/complex_mode/demo_repo"
	"aksarabase-v2/example/model"
	"fmt"
)

func main() {
	adb := aksarabase_v2.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{})
	repo := demo_repo.NewRepository(*adb)

	//============= FIRST ===============
	company, err := repo.GetByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(company)

	//============= FIND ===============
	companies, err := repo.FetchByID(2)
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
	if err != nil {
		panic(err)
	}
	fmt.Println(cstore)

}
