package student

import (
	"time"
	"fmt"
	"net/http"

	"main/infra"
	"main/model"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func ReadAlarm(c *gin.Context) {

	request := model.StudentReadAlarmRequest{}
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.StudentReadAlarmResponse{}
	errResponse := model.MessageError{}
	
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

	err := db.Table("oa").
		Select("document_id,request_at,status").
		Where("user_id = ?", request.UserID).
		Where("(status = ? AND read_flag = ?) OR status = ?" , 2, 1, -1).
		Scan(&response).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// 時刻の形式を変換
	response.RequestAt = utils.StringToTime3(response.RequestAt).Format("2006-01-02")
	
	fmt.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
