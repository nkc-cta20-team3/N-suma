package main 

import (
    "log"

    "github.com/gin-gonic/gin"

    // ローカルモジュールのインポート
    "main/infra"
    "main/handler"
    "main/service"


)

func main() {

    //
	engine := infra.DBInit()
	factory := service.NewService(engine)
	defer func() {
		log.Println("engine closed")
		engine.Close()
	}()

    // 
    g := gin.Default()
	g.Use(service.ServiceFactoryMiddleware(factory))

    // 
    routes := g.Group("/v1")
	{
        // user
		routes.POST("/users", handler.Create)
		routes.GET("/users", handler.GetAll)
		routes.GET("/users/:user-id", handler.GetOne)
		routes.PUT("/users/:user-id", handler.Update)
		routes.DELETE("/users/:user-id", handler.Delete)
        
    }

    routes2 := g.Group("/v2")
    {
        // hello world
        routes2.GET("/hello", handler.Hello)
    }

    g.Run(":8080")
}
