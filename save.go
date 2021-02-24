package main

import (
	"github.com/tealeg/xlsx/v3"
	"log"
)

func Save(c <-chan *TableInfo, finish chan<- struct{})  {
	go func() {
		wb := xlsx.NewFile()

		for t := range c{
			save(t, wb)
		}
		wb.Save("./表结构.xlsx")
		finish <- struct{}{}
	}()
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