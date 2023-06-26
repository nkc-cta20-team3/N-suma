package api

import (
	"fmt"
	"net/http"

	"main/infra"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	UserID int "json/user_id"
}

type UserPosition struct {
	Position int "json/position_id"
}

func ReadAuthList(c *gin.Context) {
	request := UserData{}

	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responceDocument := UserPosition{}
	db := infra.DBInitGorm()

	//役職IDを取得する
	db.Table("users").Select("position").Where("user_id = ?", request).First(&responceDocument)
	if db.Error != nil {
		fmt.Print("ERROR!")
	}

	//戻り値
	c.JSON(http.StatusOK, gin.H{responceDocument})
	return

}
