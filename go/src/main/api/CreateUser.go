package api

import (
	"net/http"
	"fmt"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	request := model.CreateUserRequest{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//管理者の判別

	fmt.Println(request)
	
	
	//ユーザ情報をDBに格納
	db.Table("user").
	Where("user_id = ?", request.UserID).
	Create(model.CreateUserRequest{
		UserName:        request.UserName,
		UserNumber: 	 request.UserNumber,
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