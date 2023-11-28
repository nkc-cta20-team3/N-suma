package api

import (
	"fmt"
	"net/http"
	"time"

	"main/infra"

	"github.com/gin-gonic/gin"
)

type SQLresult struct {
	StartTime time.Time `json:"start_time"`
}
type Response struct {
	StartTime string `json:"start_time"`
}

func TestDate(c *gin.Context) {

	result := SQLresult{}
	response := Response{}

	//データベース接続
	db := infra.DBInitGorm()
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}
	// データベースからデータを取得する
	err := db.Debug().Table("oa").
		Select("start_time").
		Where("document_id = ?", 1).
		Where("user_id = ?", 1).
		First(&result).Error

	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(result.StartTime)

	str := timeToString(result.StartTime)
	fmt.Println(str)

	response.StartTime = str

	// parsedTime, err := time.Parse(time.RFC3339, response.StartTime)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(parsedTime.Format("2006-01-02 15:04:05"))

	c.JSON(http.StatusOK, gin.H{"document": response})

}

func timeToString(t time.Time) string {
	str := t.Format("2006-01-02 15:04:05")
	return str
}
