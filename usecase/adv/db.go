package adv

import (
	"aksarabase-v2/domain/callbacks"
	"context"
)

//DB Advance query that's will useful for complexity case
type DB interface {
	DBFirst
	DBFind
	DBInsert
	DBUpdate
}

type DBFirst interface {
	//First exec query and scan result to single struct
	// err:=First(ctx,&company,fmt.Sprintf("select * from companies where id=%v",id))
	First(ctx context.Context, dest interface{}, query string) error
}

type DBFind interface {
	//Find exec query and scan result to multi struct base on given struct form
	//  companies:=new([]company)
	//  err:= Find(ctx,companies,func(){
	//    return 'select * from companies', company{}
	//  })
	Find(ctx context.Context, dest interface{}, form callbacks.StructForm) error
}

type DBInsert interface {
	//Insert exec query and insert to database
	//  err:= Insert(ctx,&company)
	Insert(ctx context.Context, dest interface{}) error
}

type DBUpdate interface {
	//Update exec query and update to database
	//  err:= Update(ctx,&company,fmt.Sprintf("id=%v"))
	Update(ctx context.Context, dest interface{}, whereQuery string) error
}
