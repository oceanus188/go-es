package models

type UserModel struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Age      int    `json:"age"`
	CreateAt string `json:"create_at"`
}

func (u *UserModel) Index() string {
	return "user_index"
}

func (u *UserModel) Mapping() string {
	mapping :=
		`{
		"mappings": {
			"properties": {
				"nick_name": {
					"type": "text" // 查询的时候是分词匹配
				},
				"user_name": {
					"type": "keyword" // 完整匹配
				},
				"user_id": {
					"type": "integer"
				},
				"created_at":{
					"type": "date",
					"null_value": "null",
					"format": "[yyyy-MM-dd HH:mm:ss]"
				}
			}
		}
	}`
	return mapping
}
