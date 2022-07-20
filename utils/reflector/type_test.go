package reflector

import (
	"github.com/aksara-tech/aksarabase/example/model"
	"github.com/stretchr/testify/assert"
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
