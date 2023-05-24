package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteAbsence(c *gin.Context) {

	aaa := c.Param("absence_list")

	// JSONに変えるときに使う
	// payload := gin.H{
	// 	"message": aaa,
	// }

	// ステータス200と、payloadを返します
	//c.JSON(http.StatusOK, payload)
	c.JSON(http.StatusOK, aaa)

}
