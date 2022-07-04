package scanner

import (
	"aksarabase-v2/example/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ScannerTest struct {
	model.BaseModel
	Name string `json:"name"`
}

type ScannerTestChild struct {
	model.BaseModel
	ScanResult    string      `json:"scan_result"`
	ScannerTestID int64       `json:"scanner_test_id"`
	ScannerTest   ScannerTest `json:"scanner_test" foreign:"scanner_test_id"`
}

func (s ScannerTest) TableName() string {
	return "scanner_tests"
}

func Test_inputScanner_ScanStruct(t *testing.T) {
	type args struct {
		dest interface{}
	}
	tests := []struct {
		name       string
		args       args
		tableName  string
		structName string
		len        int
	}{
		{
			name:       "basic",
			args:       args{dest: new(ScannerTest)},
			tableName:  "scanner_tests",
			structName: "ScannerTest",
			len:        5,
		},
		{
			name:       "nested",
			args:       args{dest: new(ScannerTestChild)},
			tableName:  "scanner_test_child",
			structName: "ScannerTestChild",
			len:        11,
		},
		{
			name: "has value",
			args: args{dest: ScannerTest{
				BaseModel: model.BaseModel{},
				Name:      "Wira",
			}},
			tableName:  "scanner_tests",
			structName: "ScannerTest",
			len:        5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := NewInputScanner()
			res, err := i.ScanStruct(tt.args.dest)
			assert.Nil(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, tt.tableName, res.TableName)
			assert.Equal(t, tt.structName, res.StructName)
			assert.Equal(t, tt.len, len(res.Columns))
		})
	}
}
