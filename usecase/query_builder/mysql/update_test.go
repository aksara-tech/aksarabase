package mysql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/constanta"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/info"
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/scanner/reflect_scanner"
	"gitlab.com/aksaratech/aksarabase-go/v3/utils/regex"
	"testing"
	"time"
)

func Test_queryBuilderMysql_BuildUpdateQuery(t *testing.T) {
	type args struct {
		info  info.ScanInfo
		qInfo info.QueryInfo
	}

	scanner := reflect_scanner.NewInputScanner()
	q := NewUpdateBuilder()

	company := model.Company{
		Name: "UUS",
	}
	i, _ := scanner.ScanStruct(&company)
	i.Query.WhereQuery = append(i.Query.WhereQuery, "id=1")

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "BASIC",
			args: args{
				info:  i.Scan,
				qInfo: i.Query,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time := time.Now().UTC().Format(constanta.TIME_LAYOUT)
			query := q.BuildUpdateQuery(info.Info{
				Query: tt.args.qInfo,
				Scan:  tt.args.info,
			})
			assert.Equal(t, regex.RemoveMultiString(fmt.Sprintf("Update companies SET updated_at='%v',name='UUS' WHERE id=1", time)), query)
		})
	}
}
