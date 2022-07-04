package aksarabase_v2

import (
	"aksarabase-v2/domain"
	"aksarabase-v2/usecase/query_builder"
	"aksarabase-v2/usecase/query_builder/mysql"
	"aksarabase-v2/usecase/query_executor"
	"aksarabase-v2/usecase/scanner"
)

//ADB Aksara Database Definition
type ADB struct {
	ADV
}

//ADV used for advance usage such auto generate query, scan struct and exec raw query
type ADV struct {
	//Exec execute native sql query
	Exec query_executor.SqlExecutor
	//Scanner scan input or output db scan to struct
	Scanner scanner.Scanner
	//QBuilder used for generate basic query
	QBuilder query_builder.QueryBuilder
}

//Open used to integrate all module and start the orm, scanner ,and other feature.
func Open(driver string, dsn string, config domain.Config) *ADB {
	//declare scanner
	sc := config.Engine.Scanner
	if config.Engine.Scanner == nil {
		psc := scanner.NewPointerScanner()
		osc := scanner.NewOutputScanner(psc)
		isc := scanner.NewInputScanner()

		sc = scanner.NewReflectScanner(osc, isc, psc)
	}

	//declare sql executor
	executor := query_executor.NewExecutor(driver, dsn)

	//declare query builder
	qb := mysql.NewQueryBuilder()

	return &ADB{
		ADV: ADV{
			Exec:     executor,
			Scanner:  sc,
			QBuilder: qb,
		},
	}
}
