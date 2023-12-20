package auth

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := model.GetPostResponse{}
	errResponse := model.MessageError{}

	// APIにアクセスしたユーザーの役職IDを取得
	post_id := c.MustGet("PostID").(int)
	log.Println(post_id)

	// post_idが-1の場合は、役職が未定義だと判断する
	if post_id == -1 {
		responseWrap.Document = nil
		// レスポンスを返す
		c.JSON(http.StatusOK, responseWrap)
		return
	}

	//DB接続とエラーハンドリング
	db := infra.DBInitGorm()
	if db.Error != nil {
		errResponse.Message = "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// 役職IDをもとに、DBから役職名を取得する
	err := db.Table("post").
		Select("post_name").
		Where("post_id = ?", post_id).
		First(&response).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	log.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
