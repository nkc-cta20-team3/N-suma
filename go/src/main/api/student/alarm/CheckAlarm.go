package student

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func CheckAlarm(c *gin.Context) {

	UserID := c.MustGet("UserID").(int)
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	errResponse := model.MessageError{}
	responseWrap.Document = false
	
	// 再提出の書類があるかどうかを確認
	var count int64
	err := infra.DB.Table("oa").
		Select("document_id").
		Where("status = -1").
		Where("user_id = ?", UserID).
		Count(&count).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	if count != 0  {
		//再提出の書類がある
		responseWrap.Document = true
		log.Println("RESUBMIT DOCUMENT EXIST")
	}

	/*
	// 認可完了していて、自分が未読の書類があるかどうかを確認
	count = 0	
	err = infra.DB.Table("oa").
		Select("document_id").
		Where("status = 2 AND read_flag = 1").
		Where("user_id = ?", UserID).
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
		//認可済みかつ未読の書類がある
		responseWrap.Document = true
		log.Println("UNREAD DOCUMENT EXIST")
	}
	*/

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
