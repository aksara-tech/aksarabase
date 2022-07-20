package complex_usage

import (
	"github.com/aksara-tech/aksarabase"
	"github.com/aksara-tech/aksarabase/domain"
	"github.com/aksara-tech/aksarabase/domain/constanta"
	"github.com/aksara-tech/aksarabase/example/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repository_FetchByID(t *testing.T) {
	adb := aksarabase.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{})
	repo := NewRepository(*adb)
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		wantRes []model.Company
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "BASIC",
			args: args{
				id: 1,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := repo.FetchByID(tt.args.id)
			assert.Nil(t, err)
			assert.NotNil(t, res)

		})
	}
}
