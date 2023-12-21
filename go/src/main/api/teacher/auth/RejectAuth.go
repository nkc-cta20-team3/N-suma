package teacher

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func RejectAuth(c *gin.Context) {

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
	err := infra.DB.Table("oa").
		Select("document_id").
		Where("document_id = ?", request.DocumentID).
		Where("status = 1")
	if err.Error != nil {
		// その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	if err.RowsAffected == 0 {
		// 却下できない書類の場合
		errResponse.Message = "DOCUMENT ERROR"
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

	// 公欠届を却下する
	Status := -1
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
