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

func Test_queryBuilderMysql_BuildInsertQuery(t *testing.T) {
	scanner := reflect_scanner.NewInputScanner()
	company := &model.Company{
		Name:     "TEST COMPANY",
		PicName:  "TEST",
		PicPhone: nil,
	}
	i, q, _ := scanner.ScanStruct(company)

	type args struct {
		info  info.ScanInfo
		qInfo info.QueryInfo
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "BASIC",
			args: args{
				info:  i,
				qInfo: q,
			},
			want: fmt.Sprintf("INSERT INTO companies (created_at,name,pic_name) VALUES ('%v','TEST COMPANY','TEST')", time.Now().UTC().Format(constanta.TIME_LAYOUT)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qb := NewInsertBuilder()
			s := qb.BuildInsertQuery(tt.args.info, tt.args.qInfo)
			assert.Equal(t, tt.want, regex.RemoveMultiString(s))
		})
	}
}
