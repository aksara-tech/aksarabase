package reflect_scanner

import (
	"aksarabase-v2/example/model"
	"github.com/stretchr/testify/assert"
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
