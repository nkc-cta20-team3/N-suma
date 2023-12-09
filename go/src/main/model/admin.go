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

// ReadUserListResponse は、ユーザー一覧を取得する際に使用する構造体
type ReadUserListResponse struct {
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	ClassAbbr string `json:"class_abbr"`
}

// SearchUserListRequest は、ユーザーを検索する際に使用する構造体
type SearchUserListRequest struct {
	PostID      int    `json:"post_id"`
	UserNumber  string    `json:"user_number"`
}

// SearchUserListResponse は、ユーザーを検索した結果を受け取る際に使用する構造体
type SearchUserListResponse struct {
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	ClassAbbr string `json:"class_abbr"`
}

// ReadUserRequest は、ユーザー情報を取得する際に使用する構造体
type ReadUserRequest struct {
	TargetUserID int `json:"user_id"`
}

// ReadUserResponse は、ユーザー情報を取得する際にデータを読み取るのに使用する構造体
type ReadUserResponse struct {
	UserUUID  	string 	`json:"user_uuid"`
	MailAddress string 	`json:"mail_address"`
	UserName   	string 	`json:"user_name"`
	UserNumber 	int    	`json:"user_number"`
	ClassName  	string 	`json:"class_name"`
	PostName    int    	`json:"post_name"`
}

// UpdateUserで使用する構造体
type UpdateUserRequest struct {
	UserID       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	UserNumber   int    `json:"user_number"`
	ClassID      int    `json:"class_id"`
	PostID       int    `json:"post_id"`
}

// DeleteUserで使用する構造体
type DeleteUserRequest struct {
	TargetUserID int `json:"user_id"`
}
