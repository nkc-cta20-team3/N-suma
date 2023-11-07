package api

import (
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

// CreateDocumentで使用する構造体
// type CreateDocumentRequest struct {
// 	UserID         int       `json:"user_id"`         //ユーザID
// 	RequestAt      time.Time `json:"request_at"`      // 申請日
// 	StartTime      time.Time `json:"start_date"`      // 欠席開始日
// 	EndTime        time.Time `json:"end_date"`        // 欠席終了日
// 	Location       string    `json:"location"`        // 場所
// 	StudentComment string    `json:"student_comment"` // 学生コメント
// 	DivisionID     int       `json:"division_id"`     //区分ID
// }
// type CreateDocument struct {
// 	UserID         int       `json:"user_id"`         //ユーザID
// 	RequestAt      time.Time `json:"request_at"`      // 申請日
// 	StartTime      time.Time `json:"start_time"`      // 欠席開始日
// 	StartFlame     int       `json:"start_flame"`     //開始時限
// 	EndTime        time.Time `json:"end_time"`        // 欠席終了日
// 	EndFlame       int       `json:"end_flame"`       //開始時限
// 	Location       string    `json:"location"`        // 場所
// 	Status         int       `json:"status"`          //ステータス
// 	ReadFlag       int       `json:"read_flag"`       //既読フラグ
// 	StudentComment string    `json:"student_comment"` // 学生コメント
// 	DivisionID     int       `json:"division_id"`     //区分ID
// }

type Post struct {
	PostID int
}

func CreateDocument(c *gin.Context) {
	//必要な変数定義
	request := model.CreateDocumentRequest{}

	post := Post{}

	//引数受け取り
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//役職取得
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&post)

	//学生チェック
	if post.PostID == 1 {
		//正規の処理
		oa := model.CreateDocument{
			UserID:         request.UserID,
			RequestAt:      request.RequestAt,
			StartTime:      request.StartTime,
			StartFlame:     0,
			EndTime:        request.EndTime,
			EndFlame:       0,
			Location:       request.Location,
			Status:         0,
			StudentComment: request.StudentComment,
			ReadFlag:       false,
			DivisionID:     request.DivisionID,
		}

		if err := db.Table("oa").Create(&oa).Error; err != nil {
			// 更新中にエラーが発生した場合のエラーハンドリング
			c.JSON(http.StatusBadRequest, gin.H{"document": "CREATE ERROR"})
			return
		}

	} else {
		//学生ではない時
		c.JSON(http.StatusBadRequest, gin.H{"document": "CREATE ERROR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"document": "CREATE SUCCESS"})
	return

}
