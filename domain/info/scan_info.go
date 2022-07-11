package info

type ScanInfo struct {
	TableName         string
	StructName        string
	StructAddress     interface{}
	Columns           []string
	ColumnJson        []string
	ColumnWithAliases []string
	Values            []interface{}
}

//TODO: YOU CAN USE THIS STRUCT FOR CACHING
