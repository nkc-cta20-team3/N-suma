package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestAPIRequest struct {
	UserID int `json:"user_id"`
}

func TestAPI(c *gin.Context) {

	request := TestAPIRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"document": request.UserID})
}
