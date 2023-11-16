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

// SortUserで使用する構造体
// type SortUserRequest struct {
// 	UserID       int    `json:"user_id"`
// 	SearchString string `json:"search_string"`
// }
// type SortUserResponse struct {
// 	UserID    int    `json:"user_id"`
// 	ClassAbbr string `json:"class_abbr"`
// }

func SortUser(c *gin.Context) {
	request := model.SortUserRequest{}
	response := []model.SortUserResponse{}
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

		// i, _ := strconv.Atoi(request.SearchString)

		err := db.Table("user").
			Select("user.user_id,cs.class_abbr").
			Joins("LEFT OUTER JOIN classification cs ON user.class_id = cs.class_id").
			Where("user_flag = true").
			Where("user_number LIKE ?", "%"+request.SearchString+"%").
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
		//管理者でない場合
		c.JSON(http.StatusBadRequest, gin.H{"document": "POST ERROR"})
		return
	}
}
