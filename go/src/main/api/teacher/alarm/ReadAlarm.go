package teacher

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"
	"main/utils"
	
	"github.com/gin-gonic/gin"
)

func ReadAlarm(c *gin.Context) {

	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.TeacherReadAlarmResponse{}
	errResponse := model.MessageError{}

	err := infra.DB.Table("oa").
		Select(
			"oa.document_id",
			"oa.request_at",
			"user.user_name",
			"cs.class_abbr").
		Joins("JOIN user ON oa.user_id = user.user_id").
		Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
		Where("status = ?", 1).
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
