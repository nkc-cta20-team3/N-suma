package teacher

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func ReadDocument(c *gin.Context) {

	request := model.ReadAllDocumentRequest{}
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := model.ReadAllDocumentResponse{}
	errResponse := model.MessageError{}

	log.Println(request)
	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラー処理
		errResponse.Message = err.Error()
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	
	// データベースからデータを取得する
	err := db.Table("oa").
		Select(
			"oa.request_at",
			"dv.division_name",
			"dv.division_detail",
			"oa.start_time",
			"oa.end_time",
			"oa.start_flame",
			"oa.end_flame",
			"oa.location",
			"oa.student_comment",
			"oa.teacher_comment").
		Joins("JOIN division dv ON oa.division_id = dv.division_id").
		Where("document_id = ?", request.DocumentID).
		First(&response).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// 時刻の形式を変換
	response.RequestAt = utils.StringToTime3(response.RequestAt).Format("2006-01-02")
	response.StartTime = utils.StringToTime4(response.StartTime).Format("01-02 15:04")
	response.EndTime = utils.StringToTime4(response.EndTime).Format("01-02 15:04")

	log.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}