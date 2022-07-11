package reflector

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/aksaratech/aksarabase-go/v3/example/model"
	"testing"
)

func TestGetType(t *testing.T) {

	type args struct {
		dest interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "SLICE PTR",
			args: args{
				dest: &[]model.Company{},
			},
			want: "slice",
		},
		{
			name: "SLICE",
			args: args{
				dest: []model.Company{},
			},
			want: "slice",
		},
		{
			name: "STRUCT PTR",
			args: args{
				dest: &model.Company{},
			},
			want: "struct",
		},
		{
			name: "STRUCT",
			args: args{
				dest: model.Company{},
			},
			want: "struct",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.want, GetKind(tt.args.dest))
		})
	}
}
