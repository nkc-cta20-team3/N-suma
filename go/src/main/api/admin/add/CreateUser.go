package admin

import (
	"strconv"
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	request := model.CreateUserRequest{}
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
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

	// userNumberというnil許容の数値型を定義し、
	// user_numberが空文字の場合、nilを格納する
	var userNumber *int
	if request.UserNumber == "" {
		userNumber = nil
	} else {
		// user_numberが空文字でない場合、数値型に変換して格納する
		userNumberInt, err := strconv.Atoi(request.UserNumber)
		if err != nil {
			errResponse.Message = "ユーザー番号の数値変換エラー"
			c.JSON(http.StatusInternalServerError, errResponse)
			return
		}
		userNumber = &userNumberInt
	}

	//ユーザ情報をDBに格納
	err := db.Table("user").
		Where("user_id = ?", request.UserID).
		Updates(model.CreateUserStruct{
			UserName:    request.UserName,
			UserNumber:  userNumber,
			PostID:      &request.PostID,
			ClassID:     &request.ClassID,
		}).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
