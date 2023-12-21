package teacher

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func ReadUnAuthList(c *gin.Context) {

	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.ReadAllDocumentListResponse{}
	errResponse := model.MessageError{}

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	defer db.Close()

	// 未認可書類一覧を取得
	err := db.Table("oa").
		Select(
			"oa.document_id",
			"dv.division_name",
			"dv.division_detail",
			"user.user_name").
		Joins("JOIN user ON oa.user_id = user.user_id").
		Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
		Where("oa.status = ?", 1).
		Scan(&response).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	log.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
