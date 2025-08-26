// DayCost/internal/handler/base_handler.go
package handler

import (
	"DayCost/internal/service"
	"DayCost/pkg/util"

	"github.com/gin-gonic/gin"
)

// BaseHandler 基础Handler，封装公共逻辑
type BaseHandler struct {
	baseService *service.BaseService
}

// 初始化
func NewBaseHandler() *BaseHandler {
	return &BaseHandler{
		baseService: service.NewBaseService(),
	}
}

// Bind 通用参数绑定方法，自动处理绑定错误
func (h *BaseHandler) Bind(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		util.Result(c, 400, "无效的请求格式: "+err.Error(), nil)
		return false
	}
	return true
}

// GetUserID 从上下文获取用户ID，封装重复的断言和错误处理
func (h *BaseHandler) GetUserID(c *gin.Context) (int, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		util.Result(c, 401, "未获取到用户信息", nil)
		return 0, false
	}

	id, ok := userID.(int) //断言用户ID类型
	if !ok {
		util.Result(c, 401, "无效的用户ID类型", nil)
		return 0, false
	}
	return id, true
}

// 判断expenseID和userId是否匹配来判断是否有权限？
func (h *BaseHandler) CheckExpenseExtOwner(c *gin.Context, userID int, expenseID int) bool {
	err := h.baseService.CheckExpenseExtOwner(userID, expenseID)
	if err != nil {
		util.Result(c, 401, "无权限", nil)
	}
	return true
}
