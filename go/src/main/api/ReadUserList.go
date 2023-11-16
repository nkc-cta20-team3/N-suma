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

// ReadUserListで使用する構造体
// type ReadUserListRequest struct {
// 	UserID int `json:"user_id"`
// }
// type ReadUserListResponse struct {
// 	UserID    int    `json:"user_id"`
// 	ClassAbbr string `json:"class_abbr"`
// }

func ReadUserList(c *gin.Context) {
	request := model.ReadUserListRequest{}
	response := []model.ReadUserListResponse{}
	post := model.Post{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).Scan(&post)

	if post.PostID == 0 {
		//管理者のみ実行できる
		err := db.Table("user").
			Select("user.user_id,user.user_name,cs.class_abbr").
			Joins("LEFT OUTER JOIN classification cs ON user.class_id = cs.class_id").
			Where("user_flag = true").
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
