package api

import (
	"errors"
	"fmt"
	"main/infra"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Post struct {
	PostID int
}

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
	//引数定義
	post := Post{}
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

	//管理者処理
	if PostID == 0 {

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

		//新たに管理者を増やすことができない
		/*

		 */
	} else {
		fmt.Println("権限がありません")
		c.JSON(http.StatusBadRequest, gin.H{"message": "POST ERROR"})
		return
	}

}
