package routes

import (
	"DayCost/internal/handler"

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
	expenseGroup := r.Group("/api/expenses")
	{
		expenseGroup.POST("/", expenseHandler.AddExpense)
		//expenseGroup.GET("/", expenseHandler.ListExpense)
	}

	return r
}
