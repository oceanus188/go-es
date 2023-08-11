package main

import (
	"fmt"
	"protect_es/core"
	"protect_es/docs"
	"protect_es/global"
)

func main() {
	core.EsConnect()
	fmt.Println(global.ESClient)
	//indexs.CreateIndex()
	//docs.DocCreate()
	//docs.DocDeleteById()
	//docs.DocDeleteBatch()
	//docs.DocCreateBatch()
	//docs.DocFind()
	//docs.DocFindSearch()
	docs.DocUpdate()
}
