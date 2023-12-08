package model

// ReadPrepareInformationResponse は、役職が未登録ユーザーのid,uuid,emailをリストで一覧取得する際に使用する構造体
type ReadPrepareInformationResponse struct {
	UserID   int    `json:"user_id"`
	UserUUID  string 	`json:"user_uuid"`
	MailAddress string 	`json:"mail_address"`
}	

// CreateUserRequest は、ユーザー作成時に必要な情報を受け取るために使用する構造体
type CreateUserRequest struct {
	UserID      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	UserNumber  int    `json:"user_number"`
	PostID      int    `json:"post_id"`
	ClassID     int    `json:"class_id"`
}

