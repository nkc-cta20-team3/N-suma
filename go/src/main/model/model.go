package model

import (
	"time"
)

// ReadAuthListで使用する構造体
type ReadAuthListRequest struct {
	UserID int `json:"user_id"`
}

type ReadAuthListResponse struct {
	ClassName    string `json:"class_name"`
	UserName     string `json:"user_name"`
	DivisionName string `json:"division_name"`
	DocumentID   int    `json:"document_id"`
}

type TakeClassID struct {
	PostID  int    `json:"post_id"`
	ClassID string `json:"class_id"`
}

// ReadDocumentで使用する構造体
type ReadDocumentRequest struct {
	DocumentID int `json:"document_id"`
}

type ReadDocumentResponse struct {
	DocumentID     int       `json:"document_id"`
	RequestAt      time.Time `json:"request_at"`
	StartDate      time.Time `json:"start_time"`
	StartFlame     int       `json:"start_flame"`
	EndDate        time.Time `json:"end_time"`
	EndFlame       int       `json:"end_flame"`
	Location       string    `json:"location"`
	StudentComment string    `json:"student_comment"`
	TeacherComment string    `json:"teacher_comment"`
	UserNumber     int       `json:"user_number"`
	ClassName      string    `json:"class_name"`
	UserName       string    `json:"user_name"`
}

// UpdateAuthで使用する構造体
type UpdateAuthRequest struct {
	DocumentID     int    `json:"document_id"`     //ドキュメントID
	UserNumber     int    `json:"user_number"`     //学内識別番号
	TeacherComment string `json:"teacher_comment"` //教員コメント
}

type UpdateDocument struct {
	Status         int    `json:"status"`          // ステータス
	TeacherComment string `json:"teacher_comment"` // 教員コメント
}

// RejectAuthで使用する構造体
type RejectAuthRequest struct {
	DocumentID     int    `json:"document_id"`
	UserNumber     int    `json:"user_number"`     //学内識別番号
	TeacherComment string `json:"teacher_comment"` // 教員コメント
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
	UserID         int       `json:"user_id"`         //ユーザID
	RequestAt      time.Time `json:"request_at"`      // 申請日
	StartTime      time.Time `json:"start_time"`      // 欠席開始日
	EndTime        time.Time `json:"end_time"`        // 欠席終了日
	Location       string    `json:"location"`        // 場所
	StudentComment string    `json:"student_comment"` // 学生コメント
	DivisionID     int       `json:"division_id"`     //区分ID
}

type CreateDocument struct {
	UserID         int       `json:"user_id"`         //ユーザID
	RequestAt      time.Time `json:"request_at"`      // 申請日
	StartTime      time.Time `json:"start_time"`      // 欠席開始日
	StartFlame     int       `json:"start_flame"`     //開始時限
	EndTime        time.Time `json:"end_time"`        // 欠席終了日
	EndFlame       int       `json:"end_flame"`       //開始時限
	Location       string    `json:"location"`        // 場所
	Status         int       `json:"status"`          //ステータス
	ReadFlag       bool      `json:"read_flag"`       //既読フラグ
	StudentComment string    `json:"student_comment"` // 学生コメント
	DivisionID     int       `json:"division_id"`     //区分ID
}

// DeleteDocumentで使用する構造体
// type DeleteDocumentRequest struct {
// 	DocumentID int `json:"document_id"`
// }
