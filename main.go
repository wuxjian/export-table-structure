package main

import (
	"fmt"
)

func main() {

	tables := QueryAllTables()

	queryColumnToken := make(chan struct{}, 1)

	readyTables := make(chan *TableInfo)

	finish := make(chan struct{})

	go func() {
		for _, t := range tables {
			func(t TableInfo){
				queryColumnToken <- struct{}{}
				QueryTableColumnInfo(&t)
				readyTables <- &t
				<-queryColumnToken
			}(t)

		}
		close(readyTables)
	}()

	Save(readyTables, finish)

	<- finish
	fmt.Println("all finish ...")
}
