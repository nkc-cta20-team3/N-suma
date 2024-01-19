package model

// GetPostRequest は、役職名を取得する際に使用する構造体
type GetPostResponse struct {
	PostName string `json:"post_name"` //役職名
}
