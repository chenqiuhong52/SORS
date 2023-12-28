package router

import (
	"SORS/biz/handler"
	"SORS/biz/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", handler.Login)
	auth := r.Group("/")
	auth.Use(middleware.JWT())
	{
		auth.GET("/user/info", handler.GetUserInfo)
	}

	return r
}
