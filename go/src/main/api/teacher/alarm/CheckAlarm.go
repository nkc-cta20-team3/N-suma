package teacher

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func CheckAlarm(c *gin.Context) {

	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	errResponse := model.MessageError{}
	responseWrap.Document = false
	
	// 未認可の書類があるかどうかを確認
	var count int64
	err := infra.DB.Table("oa").
		Select("document_id").
		Where("oa.status = ?", 1).
		Count(&count).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	if count > 0 {
		//再提出の書類がある
		responseWrap.Document = true
		log.Println("RESUBMIT DOCUMENT EXIST")
	}

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
