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

func ReadUserList(c *gin.Context) {

	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.ReadUserListResponse{}
	errResponse := model.MessageError{}

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	defer db.Close()

	// ユーザー一覧を取得
	err := db.Table("user").
		Select("user.user_id,user.user_name,cs.class_abbr").
		Joins("LEFT OUTER JOIN classification cs ON user.class_id = cs.class_id").
		Scan(&response).Error
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
