package mysql

import (
	"fmt"
	"github.com/aksara-tech/aksarabase/domain/constanta"
	"github.com/aksara-tech/aksarabase/domain/info"
	"github.com/aksara-tech/aksarabase/example/model"
	"github.com/aksara-tech/aksarabase/usecase/scanner/reflect_scanner"
	"github.com/aksara-tech/aksarabase/utils/regex"
	"github.com/stretchr/testify/assert"
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
	i, _ := scanner.ScanStruct(company)

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
				info:  i.Scan,
				qInfo: i.Query,
			},
			want: fmt.Sprintf("INSERT INTO companies (created_at,name,pic_name) VALUES ('%v','TEST COMPANY','TEST')", time.Now().UTC().Format(constanta.TIME_LAYOUT)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qb := NewInsertBuilder()
			s := qb.BuildInsertQuery(info.Info{
				Query: tt.args.qInfo,
				Scan:  tt.args.info,
			})
			assert.Equal(t, tt.want, regex.RemoveMultiString(s))
		})
	}
}
