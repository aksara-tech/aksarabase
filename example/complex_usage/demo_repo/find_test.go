package demo_repo

import (
	"github.com/stretchr/testify/assert"
	aksarabase_v2 "gitlab.com/wirawirw/aksarabase-go/v3"
	"gitlab.com/wirawirw/aksarabase-go/v3/domain"
	"gitlab.com/wirawirw/aksarabase-go/v3/domain/constanta"
	"gitlab.com/wirawirw/aksarabase-go/v3/example/model"
	"testing"
)

func Test_repository_FetchByID(t *testing.T) {
	adb := aksarabase_v2.Open(constanta.MYSQL, constanta.DSN_TEST, domain.Config{})
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
