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

// ReadUserで使用する構造体
// type ReadUserRequest struct {
// 	AccessUserID int `json:"access_user_id"`
// 	TargetUserID int `json:"target_user_id"`
// }
// type ReadUserResponse struct {
// 	UserName   string `json:"user_name"`
// 	UserNumber int    `json:"user_number"`
// 	ClassAbbr  string `json:"class_abbr"`
// 	PostID     int    `json:"post_id"`
// }

func ReadUser(c *gin.Context) {
	request := model.ReadUserRequest{}
	response := model.ReadUserResponse{}
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
		err := db.Table("user").
			Select("user.user_name,user.user_number,cs.class_abbr,user.post_id").
			Joins("LEFT OUTER JOIN classification cs ON user.class_id = cs.class_id").
			Where("user_flag = true").
			Where("user_id = ?", request.TargetUserID).
			Scan(&response).Error

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
		c.JSON(http.StatusOK, gin.H{"document": response})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"document": "POST ERROR"})
		return
	}
}
