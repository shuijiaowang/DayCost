package routes

import (
	"DayCost/internal/handler"
	"DayCost/pkg/middleware" // 导入中间件包

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 注册全局异常处理中间件
	r.Use(middleware.ErrorHandler())

	authHandler := handler.NewAuthHandler()
	expenseHandler := handler.NewExpenseHandler()

	// 用户路由
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/login", authHandler.Login)
		//userGroup.POST("/register", authHandler.Register)
	}

	// 需要认证的路由组
	// 增删改查
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.JWTInterceptor()) // 应用JWT拦截器
	{
		// 消费记录路由（需要认证）
		expenseGroup := apiGroup.Group("/expenses")
		{
			expenseGroup.POST("/", expenseHandler.AddExpense) // 添加消费记录
			//expenseGroup.GET("/:id", expenseHandler.GetExpense) // 获取单个消费记录
			expenseGroup.GET("/", expenseHandler.ListExpense) // 获取消费记录列表
			//expenseGroup.PUT("/:id", expenseHandler.UpdateExpense) // 更新消费记录
			//expenseGroup.DELETE("/:id", expenseHandler.DeleteExpense) // 删除消费记录

		}

	}

	return r
}
