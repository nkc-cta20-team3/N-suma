package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	// ステータス200と、messageを返します
	c.JSON(http.StatusOK, gin.H{"message": "test!!"})
	return

}
