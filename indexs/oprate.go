package indexs

import (
	"context"
	"fmt"
	"protect_es/global"
	"protect_es/models"
)

func CreateIndex() {
	if ExistIndex("user_index") {
		DeleteIndex("user_index")
	}
	var user models.UserModel
	creatIndex, err := global.ESClient.CreateIndex("user_index").BodyString(user.Mapping()).Do(context.Background())
	if err != nil {
		fmt.Println("create index err=", err)
	}
	fmt.Println(creatIndex)
	fmt.Println("索引创建成功")
}

func ExistIndex(index string) bool {
	exists, _ := global.ESClient.IndexExists(index).Do(context.Background())
	return exists
}

func DeleteIndex(index string) {
	_, err := global.ESClient.DeleteIndex(index).Do(context.Background())
	if err != nil {
		fmt.Println("delete index err=", err)
		return
	}
	fmt.Println(index, "索引删除成功")
}
