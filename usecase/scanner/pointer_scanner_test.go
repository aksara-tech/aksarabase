package scanner

import (
	"fmt"
	"testing"
	"time"
)

type BaseModel struct {
	ID        int64 `json:"id"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type Person struct {
	BaseModel
	Name   string
	Number int
	IDCard *int64
}

func Test_pointerScanner_GetListPointer(t *testing.T) {
	person := new(Person)
	res := new([]interface{})
	p := NewPointerScanner()
	p.GetListPointer(person, res)
	fmt.Println(res)
}
