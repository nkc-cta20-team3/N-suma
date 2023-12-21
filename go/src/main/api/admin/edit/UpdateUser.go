package admin

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	request := model.UpdateUserRequest{}
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

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// ユーザー情報を更新
	err := db.Table("user").
		Where("user_id = ?", request.UserID).
		Updates(model.UpdateUserStruct{
			UserName:   request.UserName,
			UserNumber: &request.UserNumber,
			PostID:     &request.PostID,
			ClassID:    &request.ClassID,
			UserFlag:	&request.UserFlag}).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
