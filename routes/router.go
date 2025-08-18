package routes

import (
	"DayCost/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	authHandler := handler.NewAuthHandler()

	// 用户路由
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/login", authHandler.Login)
	}

	return r
}
