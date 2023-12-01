package teacher

import (
	"errors"
	"fmt"
	"main/infra"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// model.go記載の構造体
// type NextDocumentRequest struct {
// 	DocumentID int  `json:"document_id"` //書類ID(基準点)
// 	UserID     int  `json:"user_id"`     //ユーザID(誰の書類？)
// 	NextFlag   bool `json:"next_flag"`   //フラグ
// }
// type NextDocumentResponse struct {
// 	DocumentID     int       `json:"document_id"`     //書類ID
// 	RequestAt      time.Time `json:"request_at"`      //申請日
// 	StartTime      time.Time `json:"start_time"`      //公欠開始日時
// 	StartFlame     int       `json:"start_flame"`     //公欠開始時限
// 	EndTime        time.Time `json:"end_time"`        //公欠終了日時
// 	EndFlame       int       `json:"end_flame"`       //公欠終了時限
// 	Location       string    `json:"location"`        //場所
// 	StudentComment string    `json:"student_comment"` //学生コメント
// 	TeacherComment string    `json:"teacher_comment"` //教員コメント
// 	UserNumber     int       `json:"user_number"`     //学籍番号
// 	ClassAddr      string    `json:"class_addr"`      //クラス略称
// 	DivisionName   string    `json:"division_name"`   //区分名
// }

type DocumentArray struct {
	DocumentID int `json:"document_id"`
}

func NextDocument(c *gin.Context) {
	//必要な変数(引数・戻り値)の定義
	request := model.NextDocumentRequest{}
	response := model.NextDocumentResponse{}
	documentArray := []DocumentArray{}

	result := model.Document{}

	// var testUserID int
	post := model.Post{}
	var PostID int

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DBに接続
	db := infra.DBInitGorm()

	//ここでユーザIDと書類の整合性を確認(教員かどうか確認)
	err := db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 行が見つからなかった場合の処理
			c.JSON(http.StatusBadRequest, gin.H{"message": "DOCUMENT NOT FOUND"})
			return
		} else {
			//その他のエラーハンドリング
			c.JSON(http.StatusBadRequest, gin.H{"message": "OTHER ERROR"})
			return
		}
	}
	PostID = post.PostID

	if PostID == 2 {
		//教員処理
		pre_id := 0
		var search_id int

		//書類ID一覧を取得
		err = db.Table("oa").Select("document_id").Where("status = ?", PostID-1).Order("document_id ASC").Scan(&documentArray).Error

		fmt.Println(documentArray)

		for _, id := range documentArray {

			if request.DocumentID == id.DocumentID && !request.NextFlag {
				//前の書類IDを取得
				search_id = pre_id
				break
			} else if request.DocumentID == id.DocumentID && request.NextFlag {
				//次の書類IDを取得(下記elseifで検索IDを取得させる)
				pre_id = -1
			} else if pre_id == -1 {
				search_id = id.DocumentID
				break
			} else {
				//次のループに行くとき
				pre_id = id.DocumentID
			}
		}

		db.Table("oa").
			Select(
				"oa.document_id",
				"oa.request_at",
				"oa.start_time",
				"oa.start_flame",
				"oa.end_time",
				"oa.end_flame",
				"oa.location",
				"oa.student_comment",
				"oa.teacher_comment",
				"dv.division_name").
			Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
			Where("document_id = ?", search_id).
			First(&result)
		if db.Error != nil {
			fmt.Print("SQL ERROR!")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "SQL ERROR",
			})
		}

		requestAt := timeToString(result.RequestAt)
		startTime := timeToString(result.StartTime)
		endTime := timeToString(result.EndTime)

		response = model.NextDocumentResponse{
			DocumentID:     result.DocumentID,
			RequestAt:      requestAt,
			StartTime:      startTime,
			StartFlame:     result.StartFlame,
			EndTime:        endTime,
			EndFlame:       result.EndFlame,
			Location:       result.Location,
			StudentComment: result.StudentComment,
			TeacherComment: result.TeacherComment,
			DivisionName:   result.DivisionName,
		}

		//戻り値
		c.JSON(http.StatusOK, gin.H{
			"message": response,
		})
		return

	} else {
		//教員でも学生でもなかったとき(不正なデータ)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "POST ERROR",
		})
		return
	}
}
