package student

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func NextDocument(c *gin.Context) {

	UserID := c.MustGet("UserID").(int)
	request := model.NextDocumentRequest{}
	responseWrap := model.ResponseWrap{}
	responseWrap.Message = "success"
	response := model.NextDocumentResponse{}
	errResponse := model.MessageError{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラー処理
		errResponse.Message = err.Error()
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	log.Println(request)

	documentArray := []model.DocumentArrayStruct{}

	//書類ID一覧を取得
	err := infra.DB.Table("oa").
		Select("document_id").
		Where("user_id = ?", UserID).
		Scan(&documentArray).
		Error
	if err != nil {
		//その他のエラーハンドリング
		errResponse.Message = "OTHER ERROR"
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}		

	//1つ前と後の書類IDを探索
	for i, v := range documentArray {
		if v.DocumentID == request.DocumentID {
			// 最初で最後の書類の場合
			if len(documentArray) == 1 {
				response.PrevDocumentID = -1
				response.NextDocumentID = -1
				break
			}
			//最初の書類の場合
			if i == 0 {
				response.PrevDocumentID = -1
				response.NextDocumentID = documentArray[i+1].DocumentID
				break
			}
			//最後の書類の場合	
			if i == len(documentArray)-1 {
				response.NextDocumentID = -1
				response.PrevDocumentID = documentArray[i-1].DocumentID
				break
			} 
			//前後の書類IDを取得
			response.PrevDocumentID = documentArray[i-1].DocumentID
			response.NextDocumentID = documentArray[i+1].DocumentID
			break
		}
		
	}

	log.Println(response)
	responseWrap.Document = response

	// レスポンスを返す
	c.JSON(http.StatusOK, responseWrap)
	return
}
