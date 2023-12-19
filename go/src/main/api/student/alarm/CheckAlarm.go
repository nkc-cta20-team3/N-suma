package student

import (
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func CheckAlarm(c *gin.Context) {

	request := model.StudentCheckAlarmRequest{}
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	errResponse := model.MessageError{}
	responseWrap.Document = false

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラー処理
		errResponse.Message = err.Error()
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	fmt.Println(request)

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	
	// 再提出の書類があるかどうかを確認
	var count int64
	err := db.Table("oa").
		Select("document_id").
		Where("status = -1").
		Where("user_id = ?", request.UserID).
		Count(&count).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	if count > 0 {
		//再提出の書類がある
		responseWrap.Document = false
		fmt.Println("RESUBMIT DOCUMENT EXIST")
	}

	// 認可完了していて、自分が未読の書類があるかどうかを確認
	count = 0	
	err = db.Table("oa").
		Select("document_id").
		Where("status = 2 AND read_flag = 1").
		Where("user_id = ?", request.UserID).
		Count(&count).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	if count > 0 {
		//認可済みかつ未読の書類がある
		responseWrap.Document = true
		fmt.Println("UNREAD DOCUMENT EXIST")
	}

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
