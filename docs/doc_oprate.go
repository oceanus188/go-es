package docs

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"protect_es/global"
	"protect_es/models"
	"time"
)

func DocCreate() {
	user := models.UserModel{
		ID:       13,
		Title:    "es文档测试-44",
		UserName: "上上下下左左右右",
		NickName: "这是最漂亮上上下下操作左左右右",
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	res, err := global.ESClient.Index().Index(user.Index()).BodyJson(user).Do(context.Background())
	if err != nil {
		fmt.Println("create doc err=", err)
		return
	}
	fmt.Printf("%v\n", res)
}

func DocCreateBatch() {
	var user models.UserModel
	index := user.Index()

	list := []models.UserModel{
		{
			ID:       12,
			UserName: "fengfeng",
			NickName: "夜空中最亮的枫枫",
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			ID:       13,
			UserName: "lisa",
			NickName: "夜空中最亮的丽萨",
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	bulk := global.ESClient.Bulk().Index(index).Refresh("true")
	for _, model := range list {
		req := elastic.NewBulkCreateRequest().Doc(model)
		bulk.Add(req)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println("batch create doc err=", err)
		return
	}
	fmt.Println("batch create doc res=", res)
}

func DocDeleteById() {
	id := "8ksG44kBn1FjpgHEJtun"
	var user models.UserModel
	index := user.Index()

	res, err := global.ESClient.Delete().Index(index).Id(id).Refresh("true").Do(context.Background())
	if err != nil {
		fmt.Println("delete doc err=", err)
		return
	}
	fmt.Println(res)

}

func DocDeleteBatch() {
	var idlist []string = []string{"9kuj44kBn1FjpgHEiduu", "9Uui44kBn1FjpgHE79uY"}
	var user models.UserModel
	index := user.Index()

	bulk := global.ESClient.Bulk().Index(index).Refresh("true")
	for _, s := range idlist {
		req := elastic.NewBulkDeleteRequest().Id(s)
		bulk.Add(req)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println("batch delete doc err=", err)
		return
	}
	fmt.Println("batch delete doc res=", res)
}

//DocFind 列表查询
func DocFind() {
	size := 2
	page := 3
	limit := (page - 1) * size
	var user models.UserModel
	index := user.Index()
	req, err := global.ESClient.Search(index).Index().From(limit).Size(2).Do(context.Background())
	if err != nil {
		fmt.Println("find doc err=", err)
		return
	}
	count := req.Hits.TotalHits.Value
	fmt.Println("count=", count)
	for _, hit := range req.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}

//DocFindSearch 模糊匹配和精确匹配
func DocFindSearch() {
	size := 2
	page := 1
	limit := (page - 1) * size

	var user models.UserModel
	index := user.Index()

	//query := elastic.NewMatchQuery("nick_name", "夜空中最亮的丽萨") //模糊匹配，适用于text类型，只要text中有的字都可以查询出来
	query:=elastic.NewTermQuery("user_name","上上下下") //精确匹配，使用keyword字段，
	res, err := global.ESClient.Search(index).Index().Query(query).From(limit).Size(10).Do(context.Background())
	if err != nil {
		fmt.Println("match search err=", err)
		return
	}
	count := res.Hits.TotalHits.Value
	fmt.Println("match count=", count)

	for _, hit := range res.Hits.Hits {
		fmt.Println(string(hit.Source))
	}

}

//DocUpdate 文档更新
func DocUpdate(){
	var user models.UserModel
	index := user.Index()
	id:="-EvC44kBn1FjpgHEVNsw"
	res,err:=global.ESClient.Update().Index(index).Id(id).Doc(map[string]any{
		"user_name":"lihaiyang",
	}).Do(context.Background())
	if err!=nil{
		fmt.Println("update doc err=",err)
		return
	}
	fmt.Printf("%v\n",res)


}
