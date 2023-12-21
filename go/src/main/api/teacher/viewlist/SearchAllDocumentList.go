package teacher

import (
	"errors"
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SearchAllDocumentList(c *gin.Context) {

	request := model.SearchAllDocumentListRequest{}
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.SearchAllDocumentListResponse{}
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
	defer db.Close()

	// 書類を検索
	err := db.Table("oa").
		Select(
			"oa.document_id",
			"dv.division_name",
			"dv.division_detail",
			"user.user_name").
		Joins("JOIN user ON oa.user_id = user.user_id").
		Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
		Where("user_number LIKE ?", "%"+request.UserNumber+"%").
		Scan(&response).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 行が見つからなかった場合の処理
			errResponse.Message = "RECORD NOT FOUND"
			c.JSON(http.StatusInternalServerError, errResponse)
			return
		} else {
			//その他のエラーハンドリング
			errResponse.Message = "OTHER ERROR"
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, errResponse)
			return
		}
	}
	
	log.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
