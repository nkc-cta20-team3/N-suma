package teacher

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func UpdateAuth(c *gin.Context) {

	request := model.UpdateAuthRequest{}
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
	log.Println(request)

	// 却下可能な書類かどうかの確認
	var count int64
	err := infra.DB.Table("oa").
		Select("document_id").
		Where("document_id = ?", request.DocumentID).
		Where("status = 1").
		Count(&count)
	if err.Error != nil {
		// その他のエラーハンドリング
		log.Println(err.Error.Error())

		errResponse.Message = "OTHER ERROR"		
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
		
	// 取得した行数をカウントする
	log.Println(count)

	// 書類が存在するかを確認する
	if count == 0 {
		// 書類が存在しない場合
		errResponse.Message = "DOCUMENT NOT FOUND"
		log.Println(errResponse.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	} else if request.TeacherComment == "" {
		// 教員コメントが空の場合
		errResponse.Message = "COMMENT ERROR"
		log.Println(errResponse.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	
	// 公欠届を認可する
	Status := 2
	err = infra.DB.Table("oa").
		Where("document_id = ?", request.DocumentID).
		Updates(model.UpdateAuthStruct{
			StartFlame:     request.StartFlame,
			EndFlame:      	request.EndFlame,
			TeacherComment: request.TeacherComment,
			Status:         Status,
		})
	if err.Error != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
