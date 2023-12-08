package hello

import (
	"net/http"

	"main/model"
	"github.com/gin-gonic/gin"
)

// GETメソッドで動作し、ステータス200と、HelloWorldを返します
// @Summary HelloWorldを返す
// @Tags Hello
// @Description GETメソッドで動作し、ステータス200と、HelloWorldを返します
// @Accept json
// @Produce json
// @Success 200 {object} model.MessageSuccess
// @Failure 400 {object} model.MessageError
// @Router /hello/get [get]
func GetHello(c *gin.Context) {
	// ステータス200と、HelloWorldを返します
	message := model.MessageSuccess{}
	message.Message = "HelloWorld. Get Method Success."
	c.JSON(http.StatusOK, message)
	return
}

// POSTメソッドで動作し、ステータス200と、HelloWorldを返します
// @Summary HelloWorldを返す
// @Tags Hello
// @Description POSTメソッドで動作し、ステータス200と、HelloWorldを返します
// @Accept json
// @Produce json
// @Success 200 {object} model.MessageSuccess
// @Failure 400 {object} model.MessageError
// @Router /hello/post [post]
func PostHello(c *gin.Context) {
	// ステータス200と、HelloWorldを返します
	message := model.MessageSuccess{}
	message.Message = "HelloWorld. Post Method Success."
	c.JSON(http.StatusOK, message)
	return

}

