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

// GETメソッドで動作し、ステータス200と、DBの役職一覧を返します
// @Summary 役職一覧を返す
// @Tags Utils
// @Description GETメソッドで動作し、ステータス200と、DBの役職一覧を返します
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseWrap
// @Failure 400 {object} model.MessageError
// @Router /utils/read/post [get]
func ReadPost(c *gin.Context) {

	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.ReadPostPresponse{}
	errResponse := model.MessageError{}

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// 役職一覧を取得
	err := db.Table("post").Select("post_id,post_name").Scan(&response).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 行が見つからなかった場合の処理
			errResponse.Message = "RECORD NOT FOUND"
			c.JSON(http.StatusNotFound, errResponse)
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
