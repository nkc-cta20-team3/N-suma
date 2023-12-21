package student

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func ReadAlarm(c *gin.Context) {

	UserID := c.MustGet("UserID").(int)
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.StudentReadAlarmResponse{}
	errResponse := model.MessageError{}

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	defer db.Close()

	err := db.Table("oa").
		Select("document_id,request_at,status").
		Where("user_id = ?", UserID).
		Where("(status = ? AND read_flag = ?) OR status = ?" , 2, 1, -1).
		Scan(&response).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	for i := 0; i < len(response); i++ {
		// 時刻の形式を変換
		response[i].RequestAt = utils.StringToTime3(response[i].RequestAt).Format("2006-01-02")
	}
	
	
	log.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
