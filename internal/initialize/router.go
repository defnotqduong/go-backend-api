package initialize

import (
	"fmt"
	"go/go-backend-api/internal/controller"
	"go/go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("Alter --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("Alter --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("Alter --> CC")
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	// use the middleware
	r.Use(middlewares.AuthMiddlewware(), BB(), CC)

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.GET("/user", controller.NewUserController().GetUserByID)
	}

	return r
}
