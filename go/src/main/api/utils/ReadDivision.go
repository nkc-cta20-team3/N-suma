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

// GETメソッドで動作し、ステータス200と、DBの区分一覧を返します
// @Summary 区分一覧を返す
// @Tags Utils
// @Description GETメソッドで動作し、ステータス200と、DBの区分一覧を返します
// @Accept json
// @Produce json
// @Success 200 {object} model.MessageSuccess
// @Failure 400 {object} model.MessageError
// @Router /utils/read/division [get]
func ReadDivision(c *gin.Context) {

	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.ReadDivisionResponse{}
	errResponse := model.MessageError{}

	// 区分一覧を取得
	err := infra.DB.Table("division").Select("division_id,division_name,division_detail").Find(&response).Error
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
