package mysql

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/info"
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
	"gitlab.com/aksaratech/aksarabase-go/v3/usecase/scanner/reflect_scanner"
	"testing"
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
	info, qinfo, _ := scanner.ScanStruct(&company)
	qinfo.Where = append(qinfo.Where, "id=1")

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "BASIC",
			args: args{
				info:  info,
				qInfo: qinfo,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := q.BuildUpdateQuery(tt.args.info, tt.args.qInfo)
			assert.Equal(t, "", query)
		})
	}
}
