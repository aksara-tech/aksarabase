package reflect_scanner

import (
	"github.com/aksara-tech/aksarabase/example/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ScannerTest struct {
	model.BaseModel
	Name string `json:"name"`
}

type ScannerTestChild struct {
	model.BaseModel
	ScanResult    string       `json:"scan_result"`
	ScannerTestID int64        `json:"scanner_test_id"`
	ScannerTest   *ScannerTest `json:"scanner_test" foreign:"scanner_test_id"`
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
		queryFrom  string
		joinLen    int
	}{
		{
			name:       "basic",
			args:       args{dest: new(ScannerTest)},
			tableName:  "scanner_tests",
			structName: "ScannerTest",
			len:        5,
			queryFrom:  "scanner_tests ScannerTest",
			joinLen:    0,
		},
		{
			name:       "nested",
			args:       args{dest: new(ScannerTestChild)},
			tableName:  "scanner_test_child",
			structName: "ScannerTestChild",
			len:        6,
			queryFrom:  "scanner_test_child ScannerTestChild",
			joinLen:    0,
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
			queryFrom:  "scanner_tests ScannerTest",
			joinLen:    0,
		},
		{
			name: "has child",
			args: args{dest: ScannerTestChild{
				ScannerTest: &ScannerTest{},
			}},
			tableName:  "scanner_test_child",
			structName: "ScannerTestChild",
			len:        11,
			queryFrom:  "scanner_test_child ScannerTestChild",
			joinLen:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := NewInputScanner()
			res, err := i.ScanStruct(tt.args.dest)
			assert.Nil(t, err)
			assert.NotNil(t, res.Query)
			assert.NotNil(t, res)
			assert.Equal(t, tt.tableName, res.Scan.TableName)
			assert.Equal(t, tt.structName, res.Scan.StructName)
			assert.Equal(t, tt.len, len(res.Scan.Columns))
			//TODO: all len must equal
			assert.Equal(t, len(res.Scan.ColumnJson), len(res.Scan.Columns))
			assert.Equal(t, len(res.Scan.ColumnWithAliases), len(res.Scan.ColumnWithAliases))
			assert.Equal(t, tt.len, len(res.Query.SelectQuery))
			//TODO: FROM Query for SelectQuery must has alias
			assert.Equal(t, tt.queryFrom, res.Query.FromQuery)
			assert.Equal(t, tt.joinLen, len(res.Query.JoinQuery))
		})
	}
}
