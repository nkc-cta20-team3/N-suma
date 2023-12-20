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

// ReadPositionで使用する構造体
type UserData struct {
	UserID int `json:"user_id"`
}

type UserPosition struct {
	PostID int `json:"post_id"`
}

type TakePostID struct {
	PostID int `json:"post_id"` //役職ID
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
