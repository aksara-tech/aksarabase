package mysql

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/info"
	"gitlab.com/aksaratech/aksarabase-go/v3/domain/query"
	"gitlab.com/aksaratech/aksarabase-go/v3/utils/regex"
	"testing"
)

func Test_queryBuilderMysql_BuildSelectQuery(t *testing.T) {

	type args struct {
		info info.ScanInfo
		q    info.QueryInfo
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "BASIC SELECT",
			args: args{
				info: info.ScanInfo{
					TableName:  "companies",
					StructName: "Company",
					Columns: []string{
						"ID", "Name",
					},
					ColumnJson: []string{
						"id", "name",
					},
					ColumnWithAliases: []string{
						"Company.id", "Company.name",
					},
					Values: []interface{}{
						"1", "urban",
					},
				},
				q: info.QueryInfo{
					Select:  nil,
					From:    "companies Company",
					Where:   nil,
					Join:    nil,
					Limit:   "",
					OrderBy: "",
				},
			},
			want: "SELECT * FROM companies Company ",
		},
		{
			name: "Basic Where",
			args: args{
				info: info.ScanInfo{
					TableName:  "companies",
					StructName: "Company",
					Columns: []string{
						"ID", "Name",
					},
					ColumnJson: []string{
						"id", "name",
					},
					ColumnWithAliases: []string{
						"Company.id", "Company.name",
					},
					Values: []interface{}{
						"1", "urban",
					},
				},
				q: info.QueryInfo{
					Select: nil,
					From:   "companies Company",
					Where:  []string{"Company.id=2"},
					Join:   nil,
					Limit:  "",
				},
			},
			want: "SELECT * FROM companies Company WHERE Company.id=2 ",
		},
		{
			name: "Multi Where",
			args: args{
				info: info.ScanInfo{
					TableName:  "companies",
					StructName: "Company",
					Columns: []string{
						"ID", "Name",
					},
					ColumnJson: []string{
						"id", "name",
					},
					ColumnWithAliases: []string{
						"Company.id", "Company.name",
					},
					Values: []interface{}{
						"1", "urban",
					},
				},
				q: info.QueryInfo{
					Select: nil,
					From:   "companies Company",
					Where:  []string{"Company.id=2", "AND", "Company.name='urban'"},
					Join:   nil,
					Limit:  "",
				},
			},
			want: "SELECT * FROM companies Company WHERE Company.id=2 AND Company.name='urban' ",
		},
		{
			name: "Multi Where",
			args: args{
				info: info.ScanInfo{
					TableName:  "companies",
					StructName: "Company",
					Columns: []string{
						"ID", "Name",
					},
					ColumnJson: []string{
						"id", "name",
					},
					ColumnWithAliases: []string{
						"Company.id", "Company.name",
					},
					Values: []interface{}{
						"1", "urban",
					},
				},
				q: info.QueryInfo{
					Select: nil,
					From:   "companies Company",
					Where:  []string{"Company.id=2", "AND", "Company.name='urban'"},
					Join:   nil,
					Limit:  "",
				},
			},
			want: "SELECT * FROM companies Company WHERE Company.id=2 AND Company.name='urban' ",
		},
		{
			name: "BASIC JOIN",
			args: args{
				info: info.ScanInfo{
					TableName:  "users",
					StructName: "User",
					Columns: []string{
						"ID", "Name", "ID", "Phone", "company_id",
					},
					ColumnJson: []string{
						"id", "name", "id", "phone", "company_id",
					},
					ColumnWithAliases: []string{
						"Company.id", "Company.name", "User.id", "User.phone", "User.company_id",
					},
					Values: []interface{}{
						"1", "urban", "1", "08912",
					},
				},

				q: info.QueryInfo{
					Select: nil,
					From:   "users User",
					Where:  []string{"User.id=1"},
					Join: []query.JoinRelation{
						{
							Join:      "LEFT JOIN",
							TableName: "companies Company",
							ON:        "User.company_id=Company.id",
						},
					},
					Limit: "",
				},
			},
			want: "SELECT * FROM users User LEFT JOIN companies Company ON User.company_id=Company.id WHERE User.id=1 ",
		},
		{
			name: "BASIC LIMIT",
			args: args{
				info: info.ScanInfo{
					TableName:  "users",
					StructName: "User",
					Columns: []string{
						"ID", "Name", "ID", "Phone", "company_id",
					},
					ColumnJson: []string{
						"id", "name", "id", "phone", "company_id",
					},
					ColumnWithAliases: []string{
						"Company.id", "Company.name", "User.id", "User.phone", "User.company_id",
					},
					Values: []interface{}{
						"1", "urban", "1", "08912",
					},
				},

				q: info.QueryInfo{
					Select: nil,
					From:   "users User",
					Where:  []string{"User.id=1"},
					Join: []query.JoinRelation{
						{
							Join:      "LEFT JOIN",
							TableName: "companies Company",
							ON:        "User.company_id=Company.id",
						},
					},
					Limit: "3",
				},
			},
			want: "SELECT * FROM users User LEFT JOIN companies Company ON User.company_id=Company.id WHERE User.id=1 LIMIT 3",
		},
		{
			name: "ORDER BY",
			args: args{
				info: info.ScanInfo{
					TableName:  "users",
					StructName: "User",
					Columns: []string{
						"ID", "Name", "ID", "Phone", "company_id",
					},
					ColumnJson: []string{
						"id", "name", "id", "phone", "company_id",
					},
					ColumnWithAliases: []string{
						"Company.id", "Company.name", "User.id", "User.phone", "User.company_id",
					},
					Values: []interface{}{
						"1", "urban", "1", "08912",
					},
				},

				q: info.QueryInfo{
					Select: nil,
					From:   "users User",
					Where:  []string{"User.id=1"},
					Join: []query.JoinRelation{
						{
							Join:      "LEFT JOIN",
							TableName: "companies Company",
							ON:        "User.company_id=Company.id",
						},
					},
					Limit:   "3",
					OrderBy: "id DESC",
				},
			},
			want: "SELECT * FROM users User LEFT JOIN companies Company ON User.company_id=Company.id WHERE User.id=1 ORDER BY id DESC LIMIT 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nq := NewSelectBuilder()
			query := nq.BuildSelectQuery(tt.args.info, tt.args.q)
			assert.Equal(t, tt.want, regex.RemoveMultiString(query))
		})
	}
}
