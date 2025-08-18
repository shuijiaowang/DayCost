package handler

import (
	"DayCost/internal/service"
	"DayCost/pkg/util"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendBadRequest(c, "无效的请求格式")
		return
	}

	user, ok := h.authService.Login(req.Name, req.Password)
	if !ok {
		util.SendUnauthorized(c, "用户名或密码错误")
		return
	}

	// 简化响应（实际应返回JWT token）
	util.SendSuccess(c, gin.H{
		"id":   user.ID,
		"name": user.Name,
	})
}
