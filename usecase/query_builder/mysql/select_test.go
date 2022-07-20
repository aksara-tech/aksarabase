package mysql

import (
	"github.com/aksara-tech/aksarabase/domain/info"
	"github.com/aksara-tech/aksarabase/domain/query"
	"github.com/aksara-tech/aksarabase/utils/regex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_queryBuilderMysql_BuildSelect(t *testing.T) {

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
			name: "BASIC SelectQuery",
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
					SelectQuery:  nil,
					FromQuery:    "companies Company",
					WhereQuery:   nil,
					JoinQuery:    nil,
					LimitQuery:   "",
					OrderByQuery: "",
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
					SelectQuery: nil,
					FromQuery:   "companies Company",
					WhereQuery:  []string{"Company.id=2"},
					JoinQuery:   nil,
					LimitQuery:  "",
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
					SelectQuery: nil,
					FromQuery:   "companies Company",
					WhereQuery:  []string{"Company.id=2", "AND", "Company.name='urban'"},
					JoinQuery:   nil,
					LimitQuery:  "",
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
					SelectQuery: nil,
					FromQuery:   "companies Company",
					WhereQuery:  []string{"Company.id=2", "AND", "Company.name='urban'"},
					JoinQuery:   nil,
					LimitQuery:  "",
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
					SelectQuery: nil,
					FromQuery:   "users User",
					WhereQuery:  []string{"User.id=1"},
					JoinQuery: []query.JoinRelation{
						{
							Join:      "LEFT JOIN",
							TableName: "companies Company",
							ON:        "User.company_id=Company.id",
						},
					},
					LimitQuery: "",
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
					SelectQuery: nil,
					FromQuery:   "users User",
					WhereQuery:  []string{"User.id=1"},
					JoinQuery: []query.JoinRelation{
						{
							Join:      "LEFT JOIN",
							TableName: "companies Company",
							ON:        "User.company_id=Company.id",
						},
					},
					LimitQuery: "3",
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
					SelectQuery: nil,
					FromQuery:   "users User",
					WhereQuery:  []string{"User.id=1"},
					JoinQuery: []query.JoinRelation{
						{
							Join:      "LEFT JOIN",
							TableName: "companies Company",
							ON:        "User.company_id=Company.id",
						},
					},
					LimitQuery:   "3",
					OrderByQuery: "id DESC",
				},
			},
			want: "SELECT * FROM users User LEFT JOIN companies Company ON User.company_id=Company.id WHERE User.id=1 ORDER BY id DESC LIMIT 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nq := NewSelectBuilder()
			query := nq.BuildSelect(info.Info{
				Query: tt.args.q,
				Scan:  tt.args.info,
			})
			assert.Equal(t, tt.want, regex.RemoveMultiString(query))
		})
	}
}
