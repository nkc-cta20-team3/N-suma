package model

// シンプルなレスポンス用の構造体

type MessageSuccess struct {
	Message string `json:"message"`
}

type MessageError struct {
	Message string `json:"error"`
}

type ResponseWrap struct {
	Message string `json:"message"`
	Document any `json:"document"`
}
