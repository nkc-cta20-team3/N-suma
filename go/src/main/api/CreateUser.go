package api

import (
	"errors"
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	post := model.Post{}
	//post_idを取得(reqest.UserIDの部分に取得したいユーザのユーザIDを入れる)
	err := db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&post).Error

	if err != nil {
		//エラーハンドリング
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 行が見つからなかった場合の処理
			fmt.Println("行が見つかりませんでした")
			c.JSON(http.StatusBadRequest, gin.H{"message": "USER NOT FOUND"})
			return
		} else {
			//その他のエラーハンドリング
			c.JSON(http.StatusBadRequest, gin.H{"message": "OTHER ERROR"})
			return
		}

	}

	PostID := post.PostID

	fmt.Println(PostID)
	if PostID == 0 {
		//管理者処理

		fmt.Println(request)

		//ユーザ情報をDBに格納
		db.Table("user").
			Where("user_id = ?", request.UserID).
			Create(model.CreateUserRequest{
				UserName:    request.UserName,
				UserNumber:  request.UserNumber,
				PostID:      request.PostID,
				ClassID:     request.ClassID,
				MailAddress: request.MailAddress,
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
		return

	} else {
		fmt.Println("アクセス権限がありません")
		c.JSON(http.StatusBadRequest, gin.H{"message": "POST ERROR"})
		return
	}

}
