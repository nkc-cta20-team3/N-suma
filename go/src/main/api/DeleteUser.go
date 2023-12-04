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

// DeleteUserで使用する構造体
//
//	type DeleteUserRequest struct {
//		AccessUserID int `json:"access_user_id"`
//		TargetUserID int `json:"target_user_id"`
//	}
//
//	type DeleteUserResponse struct {
//		Flag bool
//	}

func DeleteUser(c *gin.Context) {
	request := model.DeleteUserRequest{}
	// response := model.DeleteUserResponse{}
	post := model.Post{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	db.Table("user").Select("post_id").Where("user_id = ?", request.AccessUserID).Scan(&post)

	if post.PostID == 0 {
		//管理者のみ実行できる

		//削除(無効化)用準備
		type UpdateUser struct {
			UserFlag *bool `json:"user_flag"`
		}
		userflag := false
		update := UpdateUser{UserFlag: &userflag}

		//
		err := db.Table("user").
			Where("user_id = ?", request.TargetUserID).
			Updates(update).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 行が見つからなかった場合の処理
				fmt.Println("行が見つかりませんでした")
				c.JSON(http.StatusBadRequest, gin.H{"message": "TABLE NOT FOUND"})
				return
			} else {
				//その他のエラーハンドリング
				c.JSON(http.StatusBadRequest, gin.H{"message": "OTHER ERROR"})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"document": true})
		return
	} else {
		//管理者でない場合
		c.JSON(http.StatusBadRequest, gin.H{"document": "POST ERROR"})
		return
	}
}
