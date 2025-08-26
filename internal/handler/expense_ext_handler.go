package handler

import (
	"DayCost/internal/dto"
	"DayCost/internal/model"
	"DayCost/internal/service"
	"DayCost/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpenseExtHandler struct {
	BaseHandler                                  // 继承BaseHandler
	expenseExtService *service.ExpenseExtService // 内部创建要比传参方便一些
}

// 构造
func NewExpenseExtHandler() *ExpenseExtHandler {
	return &ExpenseExtHandler{
		BaseHandler:       *NewBaseHandler(), // 关键：初始化父结构体
		expenseExtService: service.NewExpenseExtService(),
	}
}
func (h *ExpenseExtHandler) AddExpenseExt(c *gin.Context) {
	userID, ok := h.GetUserID(c)
	if !ok {
		return // 错误已在GetUserID中返回
	}
	var req dto.ExpenseExtDto
	if !h.Bind(c, &req) {
		return
	}
	expenseExt := &model.ExpenseExt{
		ExpenseID:        req.ExpenseID,
		ExpenseType:      req.ExpenseType,
		StartDate:        req.StartDate,
		EstimatedEndDate: req.EstimatedEndDate,
		EndDate:          req.EndDate,
		TotalQuantity:    req.TotalQuantity,
		Remaining:        req.Remaining,
	}
	// 检查是否为该用户
	isOwner := h.CheckExpenseExtOwner(c, userID, req.ExpenseID)
	if !isOwner {
		return
	}
	// 添加
	err := h.expenseExtService.AddExpenseExt(userID, expenseExt)
	if err != nil {
		util.Result(c, 500, "添加失败: "+err.Error(), nil)
		return
	}
	util.Result(c, 200, "添加成功", nil)
}

// GetExpenseExtById
func (h *ExpenseExtHandler) GetExpenseExtById(c *gin.Context) {
	userID, ok := h.GetUserID(c)
	if !ok {
		return
	}

	idStr := c.Param("id") //查询id
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.Result(c, 400, "id参数错误", nil)
		return
	}
	// 检查是否为该用户
	isOwner := h.CheckExpenseExtOwner(c, userID, id)
	if !isOwner {
		return
	}
	expenseExt, err := h.expenseExtService.GetExpenseExtById(id)
	if err != nil {
		util.Result(c, 500, "获取失败: "+err.Error(), nil)
		return
	}
	util.Result(c, 200, "获取成功", expenseExt)

}
