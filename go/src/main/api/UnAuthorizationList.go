package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnAuthorizationList(c *gin.Context) {

	aaa := c.Param("teacher_data")

	// JSONに変えるときに使う
	// payload := gin.H{
	// 	"message": aaa,
	// }

	// ステータス200と、payloadを返します
	//c.JSON(http.StatusOK, payload)
	c.JSON(http.StatusOK, aaa)

}
