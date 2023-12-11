package auth

import (
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

type GetIDRequest struct {
	UserUUID string `json:"user_uuid"` //ユーザUUID
}
type GetIDResponse struct {
	UserID int `json:"user_id"` //ユーザID
}

func GetID(c *gin.Context) {

	request := GetIDRequest{}
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := GetIDResponse{}
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

	// ユーザーIDを取得
	err := db.Table("user").
		Select("user.user_id").
		Where("user_uuid = ?", request.UserUUID).
		First(&response).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	
	fmt.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
