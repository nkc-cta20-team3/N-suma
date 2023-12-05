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

// type ReadClassResponse struct {
// 	ClassID   int    `json:"class_id"`
// 	ClassAbbr string `json:"class_abbr"`
// 	ClassName string `json:"class_name"`
// }

func ReadClass(c *gin.Context) {

	response := []model.ReadClassResponse{}

	db := infra.DBInitGorm()
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	err := db.Table("classification").Select("class_id,class_abbr,class_name").Scan(&response).Error

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
