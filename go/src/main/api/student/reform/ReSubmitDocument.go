package student

import (
	"time"
	"log"
	"net/http"

	"main/infra"
	"main/model"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func ReSubmitDocument(c *gin.Context) {

	UserID := c.MustGet("UserID").(int)
	request := model.ReSubmitDocumentRequest{}
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	errResponse := model.MessageError{}
	
	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラー処理
		errResponse.Message = err.Error()
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	log.Println(request)

	//再提出の書類と提出者の整合性確認SQLを発行
	var count int64
	err := infra.DB.Table("oa").
		Select("document_id").
		Where("user_id = ?", UserID).
		Where("document_id = ?", request.DocumentID).
		Where("status = -1").
		Count(&count)
	if err.Error != nil {
		// その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// 取得した行数をカウントする
	log.Println(count)
	
	// 書類が存在するかを確認する
	if count == 0 {
		// 書類が存在しない、もしくは再提出する書類と提出者が一致しない、または再提出する書類のステータスが-1でない場合
		errResponse.Message = "DOCUMENT NOT FOUND"
		log.Println(errResponse.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	} else if request.StudentComment == "" {
		// 再提出した書類の学生コメントが空の場合
		errResponse.Message = "COMMENT ERROR"
		log.Println(errResponse.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	// 現在時刻yyyy-MM-dd HH:mm形式で取得
	// TODO: 申請日を再提出日に変更するか検討
	requestAt := time.Now()

	// 時刻型に変換
	startTime := utils.StringToTime2(request.StartTime)
	endTime := utils.StringToTime2(request.EndTime)

	// 公欠届のデータを更新
	err = infra.DB.Table("oa").
		Where("user_id = ?", UserID).
		Where("document_id = ?", request.DocumentID).
		Updates(model.ReSubmitDocumentStruct{
			RequestAt:      &requestAt,
			StartTime:      &startTime,
			EndTime:        &endTime,
			Location:       request.Location,
			Status:         1,
			StudentComment: request.StudentComment,
			DivisionID:     request.DivisionID,
		})
	if err.Error != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
