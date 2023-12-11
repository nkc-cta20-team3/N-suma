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

	type ReadUserResponse struct {
		UserUUID  	string 	`json:"user_uuid"`
		MailAddress string 	`json:"mail_address"`
		UserName   	string 	`json:"user_name"`
		UserNumber 	int    	`json:"user_number"`
		ClassID  	string 	`json:"class_id"`
		PostID    	int    	`json:"post_id"`
		UserFlag  	bool   	`json:"user_flag"`
	}

	// ユーザー情報を取得
	err := db.Table("user").
		Select("user.uuid,user.mail_address,user.user_name,user.user_number,user.class_id,user.post_id,user.user_flag").
		Where("user_id = ?", request.TargetUserID).
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
