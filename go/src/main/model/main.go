package model

import (
	"time"
)

// ReadAuthListで使用する構造体
type ReadAuthListRequest struct {
	UserID int `json:"user_id"` //ユーザID
}

type ReadAuthListResponse struct {
	ClassName    string `json:"class_name"`
	UserName     string `json:"user_name"`
	DivisionName string `json:"division_name"`
	DocumentID   int    `json:"document_id"`
	Status       int    `json:"status"`
}

type TakeClassID struct {
	PostID int `json:"post_id"`
}

// ReadDocumentで使用する構造体
type ReadDocumentRequest struct {
	DocumentID int `json:"document_id"`
	UserID     int `json:"user_id"`
}

type ReadDocumentResponse struct {
	DocumentID     int    `json:"document_id"`     //書類ID
	RequestAt      string `json:"request_at"`      //申請日
	StartTime      string `json:"start_time"`      //公欠開始日時
	StartFlame     int    `json:"start_flame"`     //公欠開始時限
	EndTime        string `json:"end_time"`        //公欠終了日時
	EndFlame       int    `json:"end_flame"`       //公欠終了時限
	Location       string `json:"location"`        //場所
	StudentComment string `json:"student_comment"` //学生コメント
	TeacherComment string `json:"teacher_comment"` //教員コメント
	DivisionName   string `json:"division_name"`   //区分名
}

type Document struct {
	DocumentID     int       `json:"document_id"`     //書類ID
	RequestAt      time.Time `json:"request_at"`      //申請日
	StartTime      time.Time `json:"start_time"`      //公欠開始日時
	StartFlame     int       `json:"start_flame"`     //公欠開始時限
	EndTime        time.Time `json:"end_time"`        //公欠終了日時
	EndFlame       int       `json:"end_flame"`       //公欠終了時限
	Location       string    `json:"location"`        //場所
	StudentComment string    `json:"student_comment"` //学生コメント
	TeacherComment string    `json:"teacher_comment"` //教員コメント
	DivisionName   string    `json:"division_name"`   //区分名
}

// UpdateAuthで使用する構造体
type UpdateAuthRequest struct {
	DocumentID     int    `json:"document_id"` //ドキュメントID
	UserID         int    `json:"user_id"`
	TeacherComment string `json:"teacher_comment"` //教員コメント
}

type UpdateDocument struct {
	Status         int    `json:"status"`          // ステータス
	TeacherComment string `json:"teacher_comment"` // 教員コメント
	ReadFlag       *bool  `json:"read_flag"`       //既読フラグ
}

// RejectAuthで使用する構造体
type RejectAuthRequest struct {
	DocumentID     int    `json:"document_id"`
	UserID         int    `json:"user_id"`
	TeacherComment string `json:"teacher_comment"` // 教員コメント
}

// NextDocumentで使用する構造体
type NextDocumentRequest struct {
	DocumentID int  `json:"document_id"` //書類ID
	UserID     int  `json:"user_id"`     //ユーザID
	NextFlag   bool `json:"next_flag"`   //フラグ
}

/*
NextFlagがtrueのときは次の書類
NextFlagがfalseのときは前の書類
*/

type NextDocumentResponse struct {
	DocumentID     int    `json:"document_id"`     //書類ID
	RequestAt      string `json:"request_at"`      //申請日
	StartTime      string `json:"start_time"`      //公欠開始日時
	StartFlame     int    `json:"start_flame"`     //公欠開始時限
	EndTime        string `json:"end_time"`        //公欠終了日時
	EndFlame       int    `json:"end_flame"`       //公欠終了時限
	Location       string `json:"location"`        //場所
	StudentComment string `json:"student_comment"` //学生コメント
	TeacherComment string `json:"teacher_comment"` //教員コメント
	DivisionName   string `json:"division_name"`   //区分名
}

// ReadPositionで使用する構造体
type UserData struct {
	UserID int `json:"user_id"`
}

type UserPosition struct {
	PostID int `json:"post_id"`
}

// CreateDocumentで使用する構造体
type CreateDocumentRequest struct {
	UserID         int    `json:"user_id"`         //ユーザID
	RequestAt      string `json:"request_at"`      // 申請日
	StartTime      string `json:"start_time"`      // 欠席開始日
	EndTime        string `json:"end_time"`        // 欠席終了日
	Location       string `json:"location"`        // 場所
	StudentComment string `json:"student_comment"` // 学生コメント
	DivisionID     int    `json:"division_id"`     //区分ID
}

type CreateDocument struct {
	UserID         int    // ユーザID
	RequestAt      string // 申請日
	StartTime      string // 欠席開始日
	StartFlame     int    // 開始時限
	EndTime        string // 欠席終了日
	EndFlame       int    // 終了時限
	Location       string // 場所
	Status         int    // 認可状態
	StudentComment string // 学生コメント
	ReadFlag       bool   // 既読フラグ
	DivisionID     int    // 区分ID
}

// type CreateDocument struct {
// 	UserID         int       // ユーザID
// 	RequestAt      time.Time // 申請日
// 	StartTime      time.Time // 欠席開始日
// 	StartFlame     int       // 開始時限
// 	EndTime        time.Time // 欠席終了日
// 	EndFlame       int       // 終了時限
// 	Location       string    // 場所
// 	Status         int       // 認可状態
// 	StudentComment string    // 学生コメント
// 	ReadFlag       bool      // 既読フラグ
// 	DivisionID     int       // 区分ID
// }

// ResubmitDocumentで使用する構造体
type ResubmitDocumentRequest struct {
	DocumentID     int    `json:"document_id"`
	UserID         int    `json:"user_id"`
	RequestAt      string `json:"request_at"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	Location       string `json:"location"`
	StudentComment string `json:"student_comment"`
	DivisionID     int    `json:"division_id"`
}

type ResubmitDocument struct {
	DocumentID     int    `json:"document_id"`     //書類ID
	RequestAt      string `json:"request_at"`      // 申請日
	StartTime      string `json:"start_time"`      // 欠席開始日
	EndTime        string `json:"end_time"`        // 欠席終了日
	Location       string `json:"location"`        // 場所
	Status         int    `json:"status"`          //ステータス
	ReadFlag       bool   `json:"read_flag"`       //既読フラグ
	StudentComment string `json:"student_comment"` // 学生コメント
	DivisionID     int    `json:"division_id"`     //区分ID
}

type CheckAlarmRequest struct {
	UserID int `json:"user_id"` //ユーザID
}

type TakePostID struct {
	PostID int `json:"post_id"` //役職ID
}


// UpdateUserで使用する構造体
type UpdateUserRequest struct {
	UserID       int    `json:"user_id"`
	UpdateUserID int    `json:"update_user_id"`
	UserName     string `json:"user_name"`
	UserNumber   int    `json:"user_number"`
	PostID       int    `json:"post_id"`
	ClassID      int    `json:"class_id"`
	MailAddress  string `json:"mail_address"`
}

// ReadDivisionで使用する構造体
type ReadDivisionRequest struct {
	UserID int `json:"user_id"`
}

type ReadDivisionResponse struct {
	DivisionID   int    `json:"division_id"`
	DivisionName string `json:"division_name"`
}


// ReadAlarmで使用する構造体
type ReadAlarmRequest struct {
	UserID int `json:"user_id"`
}
type StudentReadAlarm struct {
	DocumentID int       `json:"document_id"`
	RequestAt  time.Time `json:"request_at"`
	Status     int       `json:"status"`
}
type StudentReadAlarmResponse struct {
	DocumentID int    `json:"document_id"`
	RequestAt  string `json:"request_at"`
	Status     int    `json:"status"`
}

type TeacherReadAlarmResponse struct {
	DocumentID int    `json:"document_id"`
	UserName   string `json:"user_name"`
	ClassAbbr  string `json:"class_abbr"`
}

// ReadUserで使用する構造体
type ReadUserRequest struct {
	AccessUserID int `json:"access_user_id"`
	TargetUserID int `json:"target_user_id"`
}
type ReadUserResponse struct {
	UserName   string `json:"user_name"`
	UserNumber int    `json:"user_number"`
	ClassAbbr  string `json:"class_abbr"`
	PostID     int    `json:"post_id"`
}

// DeleteUserで使用する構造体
type DeleteUserRequest struct {
	AccessUserID int `json:"access_user_id"`
	TargetUserID int `json:"target_user_id"`
}
type DeleteUserResponse struct {
	Flag bool
}

// 役職ID取得用
type Post struct {
	PostID int
}

// ログイン者役職取得
type ReadUserPostRequest struct {
	UserID int `json:"user_id"`
}

type ReadUserPostResponse struct {
	PostID int `json:"post_id"`
}