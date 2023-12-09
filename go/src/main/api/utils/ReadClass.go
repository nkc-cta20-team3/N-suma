package api

import (
	"errors"
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GETメソッドで動作し、ステータス200と、DBのクラス一覧を返します
// @Summary クラス一覧を返す
// @Tags Utils
// @Description GETメソッドで動作し、ステータス200と、DBのクラス一覧を返します
// @Accept json
// @Produce json
// @Success 200 {object} model.MessageSuccess
// @Failure 400 {object} model.MessageError
// @Router /utils/read/class [get]
func ReadClass(c *gin.Context) {

	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.ReadClassResponse{}
	errResponse := model.MessageError{}

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// クラス一覧を取得
	err := db.Table("classification").Select("class_id,class_abbr,class_name").Scan(&response).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 行が見つからなかった場合の処理
			errResponse.Message = "RECORD NOT FOUND"
		c.JSON(http.StatusInternalServerError, errResponse)
			return
		} else {
			//その他のエラーハンドリング
			errResponse.Message = "OTHER ERROR"
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, errResponse)
			return
		}
	}

	fmt.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
