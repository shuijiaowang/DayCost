package routes

import (
	"DayCost/internal/handler"
	"DayCost/pkg/middleware" // 导入中间件包

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	authHandler := handler.NewAuthHandler()
	expenseHandler := handler.NewExpenseHandler()

	// 用户路由
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/login", authHandler.Login)
	}

	// 需要认证的路由组
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.JWTInterceptor()) // 应用JWT拦截器
	{
		// 消费记录路由（需要认证）
		expenseGroup := apiGroup.Group("/expenses")
		{
			expenseGroup.POST("/", expenseHandler.AddExpense)
			// expenseGroup.GET("/", expenseHandler.ListExpense)
		}
	}

	return r
}
