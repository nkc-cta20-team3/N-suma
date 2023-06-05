package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	
	payload := gin.H{
		"message": "Hello World",
	}
	
	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, payload)
	
}

