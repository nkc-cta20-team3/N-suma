package model

// ReadAllDocumentListResponse は、書類一覧を取得する際に使用する構造体
type ReadAllDocumentListResponse struct {
	DocumentID   int    `json:"document_id"`
	DivisionName string `json:"division_name"`
	DivisionDetail string `json:"division_detail"`
 	UserName     string `json:"user_name"`
}

// SearchAllDocumentListRequest は、書類一覧を取得する際に必要なデータを保持するための構造体
type SearchAllDocumentListRequest struct {
	UserNumber  string `json:"user_number"`
}

// SearchAllDocumentListResponse は、書類一覧を取得する際に使用する構造体
type SearchAllDocumentListResponse struct {
	DocumentID   int    `json:"document_id"`
	DivisionName string `json:"division_name"`
	DivisionDetail string `json:"division_detail"`
 	UserName     string `json:"user_name"`
}

// NextAllDocumentRequest は、現在閲覧している書類の前後の書類を取得する際に必要なデータを保持するための構造体
type NextAllDocumentRequest struct {
	DocumentID int `json:"document_id"`	//書類ID
}

// AllDocumentArrayStruct は、現在閲覧している書類の前後の書類を取得する際に使用する構造体
type AllDocumentArrayStruct struct {
	DocumentID int `json:"document_id"` //書類ID
}

// NextAllDocumentResponse は、現在閲覧している書類の前後の書類を取得する際に使用する構造体
type NextAllDocumentResponse struct {
	NextDocumentID	int  `json:"document_id_next"`		//書類ID
	PrevDocumentID  int	`json:"document_id_prev"`		//書類ID
}

// ReadAllDocumentRequest は、学生が自身の提出した書類を取得する際に必要なデータを保持するための構造体
type ReadAllDocumentRequest struct {
	DocumentID int `json:"document_id"`	//書類ID
}

// ReadAllDocumentResponse は、学生が自身の提出した書類を取得する際に使用する構造体
type ReadAllDocumentResponse struct {
	RequestAt      string `json:"request_at"`      //申請日
	DivisionName   string `json:"division_name"`   //区分名
	DivisionDetail string `json:"division_detail"` //区分詳細
	StartTime      string `json:"start_time"`      //公欠開始日時
	EndTime        string `json:"end_time"`        //公欠終了日時
	StartFlame     int    `json:"start_flame"`     //公欠開始時限
	EndFlame       int    `json:"end_flame"`       //公欠終了時限
	Location       string `json:"location"`        //場所
	StudentComment string `json:"student_comment"` //学生コメント
	TeacherComment string `json:"teacher_comment"` //教員コメント
}

// TeacherReadAlarmResponse は、通知の内容を取得する際に、データを保持するのに使用する構造体
type TeacherReadAlarmResponse struct {
	DocumentID int    `json:"document_id"`
	UserName   string `json:"user_name"`
	ClassAbbr  string `json:"class_abbr"`
}

// UpdateAuthRequest は、書類を更新する際に必要なデータを保持するための構造体
type UpdateAuthRequest struct {
	DocumentID 		int 	`json:"document_id"`		//書類ID
	StartFlame     	int    	`json:"start_flame"`     	//公欠開始時限
	EndFlame       	int    	`json:"end_flame"`       	//公欠終了時限
	TeacherComment 	string 	`json:"teacher_comment"` 	//教員コメント
}

// UpdateAuthStruct は、書類を更新する際に使用する構造体
type UpdateAuthStruct struct {
	StartFlame     	int    `json:"start_flame"`     //公欠開始時限
	EndFlame       	int    `json:"end_flame"`       //公欠終了時限
	TeacherComment 	string `json:"teacher_comment"` //教員コメント
	Status			int    `json:"status"`			//ステータス
}