package handler

import (
	"DayCost/internal/service"
	"DayCost/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	expnseService *service.ExpenseService
}

func NewExpenseHandler() *ExpenseHandler {
	return &ExpenseHandler{
		expnseService: service.NewExpenseService(),
	}
}
func (h *ExpenseHandler) AddExpense(c *gin.Context) {
	//接收参数  //使用结构体来接收
	var req struct {
		Note        string    `json:"note" binding:"required"`
		Amount      float64   `json:"amount" binding:"required"`
		Category    int       `json:"category" binding:"required"`
		ExpenseDate time.Time `json:"expense_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Result(c, 404, "无效的请求格式", nil)
		return
	}
	h.expnseService.AddExpense(req.Note, req.Amount, req.Category, req.ExpenseDate)
	//if !ok{
	//	util.Result(c, 500, "添加失败", nil)
	//	return
	//}
	//响应数据
	util.Result(c, 200, "添加成功", nil)
}
