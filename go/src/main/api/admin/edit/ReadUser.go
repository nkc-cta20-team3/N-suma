package admin

import (
	"errors"
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadUser(c *gin.Context) {
	
	request := model.ReadUserRequest{}
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := []model.ReadUserResponse{}
	errResponse := model.MessageError{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラー処理
		errResponse.Message = err.Error()
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	fmt.Println(request)

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	err := db.Table("user").
		Select("user.user_name,user.user_number,cs.class_abbr,user.post_id").
		Joins("LEFT OUTER JOIN classification cs ON user.class_id = cs.class_id").
		Where("user_flag = true").
		Where("user_id = ?", request.TargetUserID).
		Scan(&response).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 行が見つからなかった場合の処理
			fmt.Println("行が見つかりませんでした")
			c.JSON(http.StatusBadRequest, gin.H{"message": "TABLE NOT FOUND"})
			return
		} else {
			//その他のエラーハンドリング
			c.JSON(http.StatusBadRequest, gin.H{"message": "OTHER ERROR"})
			return
		}
	}

	fmt.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
