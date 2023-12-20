package student

import (
	"time"
	"log"
	"net/http"

	"main/infra"
	"main/model"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func CreateDocument(c *gin.Context) {

	UserID := c.MustGet("UserID").(int)
	request := model.CreateDocumentRequest{}
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	errResponse := model.MessageError{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラー処理
		errResponse.Message = err.Error()
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	log.Println(request)

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// 現在時刻yyyy-MM-dd HH:mm形式で取得
	requestAt := time.Now()

	// 時刻型に変換
	startTime := utils.StringToTime2(request.StartTime)
	endTime := utils.StringToTime2(request.EndTime)
	
	// ユーザ情報をDBに格納
	err := db.Table("oa").
		Create(model.CreateDocumentStruct{
			UserID:         UserID,
			RequestAt:      &requestAt,
			StartTime:      &startTime,
			StartFlame:     0,
			EndTime:        &endTime,
			EndFlame:       0,
			Location:       request.Location,
			Status:         1,
			StudentComment: request.StudentComment,
			ReadFlag:       true,
			DivisionID:     request.DivisionID,
		}).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
