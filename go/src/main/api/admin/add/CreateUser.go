package admin

import (
	"strconv"
	"net/http"

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
	
	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	userNumber := nil
	if(request.UserNumber != ""){
		userNumber = strconv.Atoi(request.UserNumber)	
	}

	//ユーザ情報をDBに格納
	err := db.Table("user").
		Where("user_id = ?", request.UserID).
		Updates(model.CreateUserRequest{
			UserName:    request.UserName,
			UserNumber:  userNumber,
			PostID:      request.PostID,
			ClassID:     request.ClassID,
		})
	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
	return

}
