package api

import (
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

// type ResubmitDocumentRequest struct {
// 	DocumentID     int       `json:"document_id"`
// 	UserID         int       `json:"user_id"`
// 	RequestAt      time.Time `json:"request_at"`
// 	StartTime      time.Time `json:"start_time"`
// 	StartFlame     int       `json:"start_flame"`
// 	EndTime        time.Time `json:"end_time"`
// 	EndFlame       int       `json:"end_flame"`
// 	Location       string    `json:"location"`
// 	StudentComment string    `json:"student_comment"`
// 	TeacherComment string    `json:"teacher_comment"`
// 	DivisionID     int       `json:"division_id"`
// }
// type ResubmitDocument struct {
// 	DocumentID     int       `json:"document_id"`     //書類ID
// 	UserID         int       `json:"user_id"`         //ユーザID
// 	RequestAt      time.Time `json:"request_at"`      // 申請日
// 	StartTime      time.Time `json:"start_time"`      // 欠席開始日
// 	StartFlame     int       `json:"start_flame"`     //開始時限
// 	EndTime        time.Time `json:"end_time"`        // 欠席終了日
// 	EndFlame       int       `json:"end_flame"`       //終了時限
// 	Location       string    `json:"location"`        // 場所
// 	Status         int       `json:"status"`          //ステータス
// 	ReadFlag       int       `json:"read_flag"`       //既読フラグ
// 	StudentComment string    `json:"student_comment"` // 学生コメント
// 	TeacherComment string    `json:"teacher_comment"` //教員コメント
// 	DivisionID     int       `json:"division_id"`     //区分ID
// }

func ResubmitDocument(c *gin.Context) {
	request := model.ResubmitDocumentRequest{}
	responseMessage := "RESUBMIT SUCCESS"

	// var test int

	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//更新用構造体
	oa := model.ResubmitDocument{
		DocumentID:     request.DocumentID,
		UserID:         request.UserID,
		RequestAt:      request.RequestAt,
		StartTime:      request.StartTime,
		StartFlame:     request.StartFlame,
		EndTime:        request.EndTime,
		EndFlame:       request.EndFlame,
		Location:       request.Location,
		Status:         1,
		ReadFlag:       false,
		StudentComment: request.StudentComment,
		DivisionID:     request.DivisionID,
	}

	//DB接続
	db := infra.DBInitGorm()

	//更新を実行
	//db.Table("oa").Where("document_id = ?", request.DocumentID).Updates(&oa)
	if err := db.Table("oa").Where("document_id = ?", request.DocumentID).Updates(&oa).Error; err != nil {
		// 更新中にエラーが発生した場合のエラーハンドリング
		responseMessage = "RESUBMIT ERROR"
	}

	//戻り値メッセージ
	c.JSON(http.StatusOK, gin.H{"document": responseMessage})

}
