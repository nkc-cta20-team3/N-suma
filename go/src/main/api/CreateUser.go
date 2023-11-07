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

	//引数定義
	var PostID int
	//post_idを取得(reqest.UserIDの部分に取得したいユーザのユーザIDを入れる)
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&PostID)

    if PostID == 0 {
  	//管理者処理
        
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
	}else{
		fmt.Println("アクセス権限がありません")
	}

}