package handler

import (
	"DayCost/internal/dto"
	"DayCost/internal/model"
	"DayCost/internal/service"
	"DayCost/pkg/util"

	"github.com/gin-gonic/gin"
)

type ExpenseExtHandler struct {
	BaseHandler                                  // 继承BaseHandler
	expenseExtService *service.ExpenseExtService // 内部创建要比传参方便一些
}

// 构造
func NewExpenseExtHandler() *ExpenseExtHandler {
	return &ExpenseExtHandler{
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

	ok := h.expenseExtService.CheckExpenseOwner(userID, req.ExpenseID)
	if !ok {
		util.Result(c, 500, "添加失败: "+err.Error(), nil)
		return
	}
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
	id := c.Param("id")
	expenseExt, err := h.expenseExtService.GetExpenseExtById(userID, id)
	if err != nil {
		util.Result(c, 500, "获取失败: "+err.Error(), nil)
		return
	}
	util.Result(c, 200, "获取成功", expenseExt)

}
