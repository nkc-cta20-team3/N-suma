package admin

import (
	"errors"
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadPrepareInformation(c *gin.Context) {
	
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.ReadPrepareInformationResponse{}
	errResponse := model.MessageError{}
	
	// 役職が未登録ユーザーのid,uuid,emailをリストで一覧取得
	err := infra.DB.Table("user").
		Select("user_id, user_uuid, mail_address").
		Where("post_id IS NULL").
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
