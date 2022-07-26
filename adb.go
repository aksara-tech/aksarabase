package aksarabase

import (
	"github.com/aksara-tech/aksarabase/domain"
	"github.com/aksara-tech/aksarabase/usecase/compiler"
	"github.com/aksara-tech/aksarabase/usecase/query_builder"
	"github.com/aksara-tech/aksarabase/usecase/query_builder/mysql"
	"github.com/aksara-tech/aksarabase/usecase/query_executor"
	"github.com/aksara-tech/aksarabase/usecase/query_executor/default_executor"
	"github.com/aksara-tech/aksarabase/usecase/scanner"
	"github.com/aksara-tech/aksarabase/usecase/scanner/reflect_scanner"
)

//ADB Aksara Database Definition
type ADB struct {
	//Exec execute native sql query
	Exec query_executor.SqlExecutor
	//Scanner scan input or output db scan to struct
	Scanner scanner.Scanner
	//QBuilder used for generate basic query
	QBuilder query_builder.QueryBuilder
	//ORM is easy and simple command to get or execute data from database
	compiler.Compiler
}

//Open used to integrate all module and start the compiler, scanner ,and other feature.
func Open(driver string, dsn string, config domain.Config) *ADB {
	//declare scanner
	scanner := setupScanner(config.Engine)

	//declare sql executor
	executor := setupExecutor(driver, dsn, config.Engine)

	//declare query builder
	queryBuilder := setupQueryBuilder(config.Engine)

	return &ADB{
		Exec:     executor,
		Scanner:  scanner,
		QBuilder: queryBuilder,
		Compiler: compiler.NewCompiler(scanner, queryBuilder, executor),
	}
}

func setupExecutor(driver string, dsn string, engine domain.Engine) (exc query_executor.SqlExecutor) {
	var e query_executor.SqlExecutor
	if engine.SqlExecutor != nil {
		e = engine.SqlExecutor
	} else {
		e = default_executor.NewExecutor(driver, dsn)
	}

	return e
}

func setupQueryBuilder(engine domain.Engine) (qb query_builder.QueryBuilder) {
	var insertBuilder query_builder.InsertQueryBuilder
	if engine.InsertQueryBuilder != nil {
		insertBuilder = engine.InsertQueryBuilder
	} else {
		insertBuilder = mysql.NewInsertBuilder()
	}

	var updateBuilder query_builder.UpdateQueryBuilder
	if engine.UpdateQueryBuilder != nil {
		updateBuilder = engine.UpdateQueryBuilder
	} else {
		updateBuilder = mysql.NewUpdateBuilder()
	}

	var SelectBuilder query_builder.SelectBuilder
	if engine.SelectBuilder != nil {
		SelectBuilder = engine.SelectBuilder
	} else {
		SelectBuilder = mysql.NewSelectBuilder()
	}
	qb = mysql.NewQueryBuilderMysql(insertBuilder, SelectBuilder, updateBuilder)
	return
}

func setupScanner(engine domain.Engine) (sc scanner.Scanner) {
	var psc scanner.PointerScanner
	if engine.PointerScanner != nil {
		psc = engine.PointerScanner
	} else {
		psc = reflect_scanner.NewPointerScanner()
	}

	var osc scanner.OutputScanner
	if engine.PointerScanner != nil {
		osc = engine.OutputScanner
	} else {
		osc = reflect_scanner.NewOutputScanner(psc)
	}

	var isc scanner.InputScanner
	if engine.InputScanner != nil {
		isc = engine.InputScanner
	} else {
		isc = reflect_scanner.NewInputScanner()
	}
	sc = reflect_scanner.NewReflectScanner(osc, isc, psc)
	return
}
