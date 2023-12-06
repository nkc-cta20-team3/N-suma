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

// type ReadUserPostRequest struct {
// 	UserID int `json:"user_id"`
// }

// type ReadUserPostResponse struct {
// 	PostID int `json:"post_id"`
// }

func ReadUserPost(c *gin.Context) {
	request := model.ReadUserPostRequest{}
	response := model.ReadUserPostResponse{}

	//引数受け取り
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := infra.DBInitGorm()
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	err := db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&response).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 行が見つからなかった場合の処理
			fmt.Println("行が見つかりませんでした")
			c.JSON(http.StatusBadRequest, gin.H{"message": "DATA NOT FOUND"})
			return
		} else {
			//その他のエラーハンドリング
			c.JSON(http.StatusBadRequest, gin.H{"message": "OTHER ERROR"})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{"document": response})
	return

}
