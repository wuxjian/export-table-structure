package main

import (
	"github.com/tealeg/xlsx/v3"
	"io"
	"log"
	"os"
)

func Save(c <-chan *TableInfo, finish chan<- struct{}) {
	go func() {
		//wb := xlsx.NewFile()
		file, err := os.OpenFile("表结构.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 066)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		for t := range c {
			saveTxt(t, file)
		}
		//wb.Save("./表结构.xlsx")
		finish <- struct{}{}
	}()
}

func saveTxt(t *TableInfo, w io.Writer) {
	tableName := t.TableName
	if t.TableComment != "" {
		tableName += "(" + t.TableComment + ")"
	}

	io.WriteString(w, "表名:"+tableName+"\r\n")

	io.WriteString(w, "字段名称\t数据类型\t必填\t注释\r\n")
	for _, column := range t.Columns {
		io.WriteString(w, column.ColumnName+"\t"+column.ColumnType+"\t"+column.Required+"\t"+column.Comment+"\r\n")
	}
}

func save(t *TableInfo, wb *xlsx.File) {
	sheetName := t.TableName
	if len(sheetName) > 31 {
		sheetName = sheetName[:31]
	}

	sheet, err := wb.AddSheet(sheetName)

	if err != nil {
		log.Fatalf("add sheet[%s] fail:%v\n", t.TableName, err)
	}

	// 第一行
	row := sheet.AddRow()
	row.SetHeight(12.8)
	cell := row.AddCell()

	firstContent := t.TableName
	if t.TableComment != "" {
		firstContent += "(" + t.TableComment + ")"
	}
	cell.SetValue(firstContent)

	// 第二行
	row = sheet.AddRow()
	row.SetHeight(12.8)
	cell = row.AddCell()
	cell.SetValue("字段名称")
	cell = row.AddCell()
	cell.SetValue("数据类型")
	cell = row.AddCell()
	cell.SetValue("必填")
	cell = row.AddCell()
	cell.SetValue("注释")

	for _, column := range t.Columns {
		row = sheet.AddRow()
		row.SetHeight(12.8)
		cell = row.AddCell()
		cell.SetValue(column.ColumnName)
		cell = row.AddCell()
		cell.SetValue(column.ColumnType)
		cell = row.AddCell()
		cell.SetValue(column.Required)
		cell = row.AddCell()
		cell.SetValue(column.Comment)
	}
}
