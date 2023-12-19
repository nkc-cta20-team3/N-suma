package student

import (
	"errors"
	"fmt"
	"main/infra"
	"main/model"
	"net/http"
	
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ReadAlarmで使用する構造体
// type ReadAlarmRequest struct {
// 	UserID int `json:"user_id"`
// }
// type StudentReadAlarmResponse struct {
// 	DocumentID int       `json:"document_id"`
// 	RequestAt  time.Time `json:"request_at"`
// 	Status     int       `json:"status"`
// }

func ReadAlarm(c *gin.Context) {
	request := model.ReadAlarmRequest{}
	studentResponse := []model.StudentReadAlarmResponse{}
	post := model.Post{}

	result := []model.StudentReadAlarm{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := infra.DBInitGorm()

	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).Scan(&post)

	if post.PostID == 1 {
		//学生の処理

		err := db.Table("oa").
			Select("document_id,request_at,status").
			Where("user_id = ? AND ((status = ? AND read_flag = ?) OR status = ?)", request.UserID, 2, 1, -1).
			Scan(&result).Error

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

		//日付変換処理
		for _, t := range result {

			//日付変換処理
			requestAt := timeToString(t.RequestAt)
			document := model.StudentReadAlarmResponse{
				DocumentID: t.DocumentID,
				RequestAt:  requestAt,
				Status:     t.Status,
			}
			studentResponse = append(studentResponse, document)
		}

		//戻り値
		c.JSON(http.StatusOK, gin.H{"document": studentResponse})
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"document": "POST ERROR"})
		return
	}
}

func timeToString(t time.Time) string {
	str := t.Format("2006-01-02 15:04:05")
	return str
}
