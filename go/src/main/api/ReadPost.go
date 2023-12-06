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

// type ReadPostPresponse struct {
// 	PostID   int    `json:"post_id"`
// 	PostName string `json:"post_name"`
// }

func ReadPost(c *gin.Context) {

	response := []model.ReadPostPresponse{}

	db := infra.DBInitGorm()
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	err := db.Table("post").Select("post_id,post_name").Scan(&response).Error

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

}
