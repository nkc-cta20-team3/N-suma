package admin

import (
	"errors"
	"fmt"
	"main/infra"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUser(c *gin.Context) {

	request := model.UpdateUserRequest{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//DBに接続
	db := infra.DBInitGorm()

	//更新処理
	db.Table("user").
		Where("user_id = ?", request.UpdateUserID).
		Updates(model.UpdateUserRequest{
			UserName:    request.UserName,
			UserNumber:  request.UserNumber,
			PostID:      request.PostID,
			ClassID:     request.ClassID,
			MailAddress: request.MailAddress})

	//エラーハンドリング
	if db.Error != nil {
		errMsg := "UPDATE ERROR"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flag": true})
	return

}
