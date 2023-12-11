package model

// ReadAllDocumentListResponse は、書類一覧を取得する際に使用する構造体
type ReadAllDocumentListResponse struct {
	DocumentID   int    `json:"document_id"`
	DivisionName string `json:"division_name"`
	DivisionDetail string `json:"division_detail"`
 	UserName     string `json:"user_name"`
}