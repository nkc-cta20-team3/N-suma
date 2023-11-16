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

// ReadDivisionで使用する構造体
// type ReadDivisionRequest struct {
// 	UserID int `json:"user_id"`
// }

// type ReadDivisionResponse struct {
// 	DivisionID   int    `json:"division_id"`
// 	DivisionName string `json:"division_name"`
// }

func ReadDivision(c *gin.Context) {

	request := model.ReadDivisionRequest{}
	response := []model.ReadDivisionResponse{}
	post := model.Post{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//役職ID取得
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).Scan(&post)

	if post.PostID == 1 {
		//学生のみアクセス可能
		err := db.Table("division").Select("division_id,division_name").Find(&response).Error

		//エラーハンドリング
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
