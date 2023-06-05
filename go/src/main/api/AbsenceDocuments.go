package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AbsenceDocument(c *gin.Context) {

	aaa := c.Param("document_id")

	// JSONに変えるときに使う
	// payload := gin.H{
	// 	"message": aaa,
	// }

	// ステータス200と、payloadを返します
	//c.JSON(http.StatusOK, payload)
	c.JSON(http.StatusOK, aaa)

}
