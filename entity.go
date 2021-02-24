package main

type TableInfo struct {
	TableName    string
	TableComment string
	Columns      []ColumnsInfo
}

// 列信息
type ColumnsInfo struct {
	ColumnName string
	ColumnType string
	Required   string
	Comment    string
}
