package main

import (
	"github.com/boymeetsblockchain/go-jwt/controllers"
	"github.com/boymeetsblockchain/go-jwt/initializers"
	"github.com/boymeetsblockchain/go-jwt/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/sign-up", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate-token", middlewares.RequireAuth, controllers.ValidateToken)
	r.Run() // listen and serve on
}
