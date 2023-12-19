package auth

import (
	"fmt"
	"net/http"

	// "main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

type GetIDResponse struct {
	UserID int `json:"user_id"` //ユーザID
}

func GetPost(c *gin.Context) {
	
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := GetIDResponse{}
	// errResponse := model.MessageError{}

	// ユーザーIDを取得
	post_id := c.MustGet("PostID").(int)
	fmt.Println(post_id)

	// post_idが0の場合は、役職が未定義のだと判断する

	// これ以下TODO

	/*

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// ユーザーIDを取得
	err := db.Table("user").
		Select("user_id").
		Where("user_uuid = ?", uuid).
		First(&response).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	*/
	
	fmt.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
