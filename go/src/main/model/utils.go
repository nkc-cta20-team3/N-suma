package model

// ReadClassResponse は、クラス一覧取得時に使用する構造体
type ReadClassResponse struct {
	ClassID   int    `json:"class_id"`
	ClassAbbr string `json:"class_abbr"`
	ClassName string `json:"class_name"`
}

// ReadPostResponse は、役職一覧取得時に使用する構造体
type ReadPostPresponse struct {
	PostID   int    `json:"post_id"`
	PostName string `json:"post_name"`
}
