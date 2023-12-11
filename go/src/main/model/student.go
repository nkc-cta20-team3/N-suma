package model

import (
	"time"
)

// CreateDocumentRequest は、書類を作成時に必要なデータを取得するための構造体
type CreateDocumentRequest struct {
	UserID         int    `json:"user_id"`         //ユーザID
	StartTime      string `json:"start_time"`      // 欠席開始日
	EndTime        string `json:"end_time"`        // 欠席終了日
	Location       string `json:"location"`        // 場所
	StudentComment string `json:"student_comment"` // 学生コメント
	DivisionID     int    `json:"division_id"`     //区分ID
}

// CreateDocumentStruct は、書類を作成する際に、DBにデータを追加するときに使用する構造体
type CreateDocumentStruct struct {
	UserID         int    // ユーザID
	RequestAt      *time.Time // 申請日
	StartTime      *time.Time // 欠席開始日
	StartFlame     int    // 開始時限
	EndTime        *time.Time // 欠席終了日
	EndFlame       int    // 終了時限
	Location       string // 場所
	Status         int    // 認可状態
	StudentComment string // 学生コメント
	ReadFlag       bool   // 既読フラグ
	DivisionID     int    // 区分ID
}

// ReadDocsListRequest は、学生が自身の提出した書類を取得する際に必要なデータを保持するための構造体
type ReadDocsListRequest struct {
	UserID int `json:"user_id"` //ユーザID
}

// ReadDocsListResponse は、学生が自身の提出した書類を取得する際に使用する構造体
type ReadDocsListResponse struct {
	DocumentID   	int    	`json:"document_id"`		//書類ID
	RequestAt   	string	`json:"request_at"` 	   	//申請日
	DivisionName 	string	`json:"division_name"`		//区分名
	DivisionDetail 	string	`json:"division_detail"`	//区分詳細
}
