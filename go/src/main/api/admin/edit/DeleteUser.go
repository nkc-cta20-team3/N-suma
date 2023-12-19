package admin

import (
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {

	request := model.DeleteUserRequest{}
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
	fmt.Println(request)

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	/*
	userflag := false
	update := UpdateUser{UserFlag: &userflag}
	err := db.Table("user").
		Where("user_id = ?", request.TargetUserID).
		Updates(update).Error	
	*/

	// ユーザー情報を削除(ユーザーを凍結処理)
	userflag := false
	err := db.Table("user").
		Where("user_id = ?", request.TargetUserID).
		Updates(model.DeleteUserStruct{
			UserFlag: &userflag }).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
