package scanner

type ScanInfo struct {
	TableName  string
	StructName string
	Columns    []string
	Values     []interface{}
}
