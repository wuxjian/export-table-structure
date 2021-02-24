package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func init() {
	InitDB()
}

func InitDB() {
	//数据库连接
	var err error
	DB, err = sql.Open("mysql", "root:123456@(10.30.14.187:3306)/webserver3")
	if err != nil {
		log.Fatalf("数据库连接失败: error=%s\n", err.Error())
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalf("数据库连接失败: error=%s\n", err.Error())
	}
}

func QueryAllTables() []TableInfo {
	rows, err := DB.Query("select table_name , table_comment from INFORMATION_SCHEMA.tables where table_schema ='webserver3'")
	if err != nil {
		log.Fatal("查询所有表失败", err)
	}
	var r []TableInfo
	var sum int
	for rows.Next() {
		var item TableInfo
		rows.Scan(&item.TableName, &item.TableComment)
		sum ++
		r = append(r, item)
	}
	fmt.Printf("共有%d张表\n", sum)
	return r
}

func QueryTableColumnInfo(t *TableInfo){
	rows, err := DB.Query("select COLUMN_NAME, COLUMN_TYPE, IF(IS_NULLABLE='NO','是','否') AS '必填', COLUMN_COMMENT 注释 from INFORMATION_SCHEMA.COLUMNS where table_schema ='webserver3' and table_name = ?", t.TableName)
	if err != nil {
		log.Fatalf("查询表[%s]中的列表失败:error=%s", t.TableName, err.Error())
	}
	var columns []ColumnsInfo

	for rows.Next() {
		var item ColumnsInfo
		rows.Scan(&item.ColumnName, &item.ColumnType, &item.Required, &item.Comment)
		columns = append(columns, item)
	}

	t.Columns = columns
}