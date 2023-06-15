package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateAuth(c *gin.Context) {

	request := model.UpdateAuthRequest{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": request,
	})

}
