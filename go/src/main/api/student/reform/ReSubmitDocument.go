package student

import (
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

// ResubmitDocumentで使用する構造体
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
// 	DivisionID int `json:"division_id"`
// }

// type ResubmitDocument struct {
// 	DocumentID int `json:"document_id"` //書類ID
// 	RequestAt      time.Time `json:"request_at"`      // 申請日
// 	StartTime      time.Time `json:"start_time"`      // 欠席開始日
// 	StartFlame     int       `json:"start_flame"`     //開始時限
// 	EndTime        time.Time `json:"end_time"`        // 欠席終了日
// 	EndFlame       int       `json:"end_flame"`       //終了時限
// 	Location       string    `json:"location"`        // 場所
// 	Status         int       `json:"status"`          //ステータス
// 	ReadFlag       bool      `json:"read_flag"`       //既読フラグ
// 	StudentComment string    `json:"student_comment"` // 学生コメント
// 	DivisionID int `json:"division_id"` //区分ID
// }

func ReSubmitDocument(c *gin.Context) {
	request := model.ResubmitDocumentRequest{}
	responseMessage := "RESUBMIT SUCCESS"

	var count int64

	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//再提出の書類と提出者の整合性確認SQLを発行
	db.Table("oa").Select("document_id").Where("user_id = ?", request.UserID).Where("document_id = ?", request.DocumentID).Where("status = -1").Count(&count)

	if count == 0 {
		//再提出する書類と提出者が一致しない、または再提出する書類のステータスが-1でない場合
		c.JSON(http.StatusBadRequest, gin.H{"document": "DOCUMENT ERROR"})
		return
	} else if request.StudentComment == "" {
		c.JSON(http.StatusBadRequest, gin.H{"document": "COMMENT ERROR"})
		return
	}

	//更新用構造体
	oa := model.ResubmitDocument{
		DocumentID:     request.DocumentID,
		RequestAt:      request.RequestAt,
		StartTime:      request.StartTime,
		EndTime:        request.EndTime,
		Location:       request.Location,
		Status:         1,
		ReadFlag:       true,
		StudentComment: request.StudentComment,
		DivisionID:     request.DivisionID,
	}

	//更新を実行
	//db.Table("oa").Where("document_id = ?", request.DocumentID).Updates(&oa)
	if err := db.Table("oa").Where("document_id = ?", request.DocumentID).Updates(&oa).Error; err != nil {
		// 更新中にエラーが発生した場合のエラーハンドリング
		responseMessage = "RESUBMIT ERROR"
	}

	//戻り値メッセージ
	c.JSON(http.StatusOK, gin.H{"document": responseMessage})

}
