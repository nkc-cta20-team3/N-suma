package student

import (
	"errors"
	"main/infra"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DocumentArray struct {
	DocumentID int `json:"document_id"`
}

func NextDocument(c *gin.Context) {
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

	if PostID == 1 {
		//学生用処理
		pre_id := 0
		var search_id int

		//書類ID一覧を取得
		err = db.Table("oa").Select("document_id").Where("user_id = ?", request.UserID).Scan(&documentArray).Error

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

		//切り替え対象書類取得
		err := db.Debug().Table("oa").
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
			Joins("JOIN division dv ON oa.division_id = dv.division_id").
			Where("document_id = ?", search_id).
			Where("oa.user_id = ?", request.UserID).
			First(&result).Error

		if db.Error != nil {
			errMsg := "データベースからデータを取得できませんでした"
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"document": "書類が見つかりませんでした"})
			return
		}

		//日付変換(関数はreaddocument)
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "POST ERROR",
		})
		return
	}

}
