package teacher

import (
	"errors"
	"net/http"
	"time"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DocumentStatus struct {
	Status int
}

func ReadDocument(c *gin.Context) {

	request := model.ReadDocumentRequest{}
	response := model.ReadDocumentResponse{}

	result := model.Document{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//log.Print(request.DocumentID)

	//データベース接続
	db := infra.DBInitGorm()
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	//引数定義
	post := model.Post{}
	//post_idを取得(reqest.UserIDの部分に取得したいユーザのユーザIDを入れる)
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&post)

	PostID := post.PostID

	// fmt.Println(PostID)
	//管理者処理
	if PostID == 0 {
		// fmt.Println("権限がありません")
		c.JSON(http.StatusBadRequest, gin.H{"document": "権限がありません"})
		return

	} else if PostID == 2 {
		//教員処理

		// データベースからデータを取得する
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
			//Where("document_id = ?", request.DocumentID).
			First(&result).Error
		if db.Error != nil {
			errMsg := "データベースからデータを取得できませんでした"
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"document": "書類が見つかりませんでした"})
			return
		}

		// log.Println(response)

		//日付変換処理
		requestAt := timeToString(result.RequestAt)
		startTime := timeToString(result.StartTime)
		endTime := timeToString(result.EndTime)
		response = model.ReadDocumentResponse{
			//DocumentID:     result.DocumentID,
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

		//値を返却
		c.JSON(http.StatusOK, gin.H{"document": response})
		return
	} else {
		// fmt.Println("権限がありません")

		c.JSON(http.StatusBadRequest, gin.H{"document": "権限がありません"})
		return
	}
}
func timeToString(t time.Time) string {
	str := t.Format("2006-01-02 15:04:05")
	return str
}
