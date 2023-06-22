package api

import{
	"fmt"

	"main/infra"
	"main/model"
	"github.com/gin-gonic/gin"
}

type struct UserData {
	UserID int "json/user_id"
}

type struct UserPosition{
	Position
}

func ReadAuthList(c *gin.Context) {
	request := UserData{}

	if err:= c.ShouldBindJSON(&request); err:= nil{
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responce := UserPosition{}
	db := infra.DBInitGorm()

	db.Table("users").Select("position").Where("user_id = ?",request).First(&responce)

	c.JSON(http.StatusOK, gin.H{"document": response.UserID})
	return

}