package main

import (
	"DayCost/pkg/database"
	"DayCost/routes"
	"log"
)

func main() {
	// 初始化数据库
	if err := database.InitDB(); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 创建路由
	r := routes.SetupRouter()

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
