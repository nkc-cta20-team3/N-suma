package student

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func ReadDocsList(c *gin.Context) {

	UserID := c.MustGet("UserID").(int)
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.ReadDocsListResponse{}
	errResponse := model.MessageError{}

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	defer db.Close()
	
	// 書類一覧を取得
	err := db.Table("oa").
		Select(
			"oa.document_id",
			"oa.request_at",
			"dv.division_name",
			"dv.division_detail").
		Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
		Where("oa.user_id = ?", UserID).
		Scan(&response).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// 時刻の形式を変換
	for i := 0; i < len(response); i++ {
		response[i].RequestAt = (utils.StringToTime3(response[i].RequestAt)).Format("2006-01-02")
	}

	log.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
