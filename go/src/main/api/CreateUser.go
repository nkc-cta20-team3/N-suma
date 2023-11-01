package api

import (
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func UpdateAuth(c *gin.Context) {

	request := model.CreateUserRequest{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()


	//ユーザ情報をDBに格納
	db.Table("user").
	Where("user_id = ?", request.UserID)
	Create(model.CreateUserRequest{
		UserNumber:      request.UserNumber,
		UserName:        request.UserName,
		CreateUserNumber: request.CreateUserNumber,
		PostID:          request.PostID,
		ClassID:         request.ClassID,
		MailAddress:     request.MailAddress,
	})

	//エラーハンドリング
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusOK,
	})
}