package controller

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Firebase SDK のセットアップ
		opt := option.WithCredentialsFile(os.Getenv("GOOGLE_CREDENTIALS_JSON"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Printf("err: %v\n", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Printf("err: %v\n", err)
			os.Exit(1)
		}

		// AuthorizationヘッダからIDトークン(JWT)を取得
		authHandler := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHandler, "Bearer ", "", 1)

		// IDトークンを検証
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			// 検証エラー
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Error": err.Error(),
			})
			return
		}
		// ロギング
		log.Printf("Vertifed ID token: %v\n", token)
		c.Next()
	}
}
