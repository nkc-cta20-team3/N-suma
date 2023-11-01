package api

import (
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	request := model.UpdateUserRequest{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()
	//管理者の判定(未完成)
	// post := post.PostID
	
	/*
	if post == ? {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "権限がありません"})
		return
	}*/
	
	//更新処理
	db.Table("user").
	Where("user_id = ?", request.UserID).
	Updates(model.UpdateUserRequest{
		UserID: request.UserID,
		UserName: request.UserName,
		UserNumber: request.UserNumber,
		PostID: request.PostID,
		ClassID: request.ClassID,
		MailAddress: request.MailAddress})
	
	//エラーハンドリング
	if db.Error != nil {
		errMsg := "UPDATE ERROR"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flag": "データの更新が成功しました"})
	
	//新たに管理者を増やすことができない(分からんそもそもUpdateUserで作るの？)
	

}