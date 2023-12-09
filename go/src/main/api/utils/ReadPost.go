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

func ReadPost(c *gin.Context) {

	response := []model.ReadPostPresponse{}

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	// 役職一覧を取得
	err := db.Table("post").Select("post_id,post_name").Scan(&response).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 行が見つからなかった場合の処理
			fmt.Println("行が見つかりませんでした")
			c.JSON(http.StatusBadRequest, gin.H{"message": "TABLE NOT FOUND"})
			return
		} else {
			//その他のエラーハンドリング
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "OTHER ERROR" })
			return
		}
	}
	
	fmt.Println(response)

	// レスポンスを返す
	c.JSON(http.StatusOK, gin.H{"document": response})
	return

}
