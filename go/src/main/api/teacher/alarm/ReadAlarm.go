package teacher

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"
	
	"github.com/gin-gonic/gin"
)

func ReadAlarm(c *gin.Context) {

	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.TeacherReadAlarmResponse{}
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
		Select(
			"oa.document_id",
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

	log.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
