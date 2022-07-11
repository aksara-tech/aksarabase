package reflect_scanner

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
	"testing"
)

type Person struct {
	model.BaseModel
	Name   string
	Number int
	IDCard *int64
}

func Test_pointerScanner_GetListPointer(t *testing.T) {
	person := new(Person)
	res := new([]interface{})
	p := NewPointerScanner()
	p.GetListPointer(person, res)
	assert.Equal(t, len(*res), 7)
}
