package handler

import (
	"DayCost/internal/dto"
	"DayCost/internal/model"
	"DayCost/internal/service"
	"DayCost/pkg/util"
	"fmt"

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
	fmt.Println("---------读取参数")
	// JWT,从上下文获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		util.Result(c, 401, "未获取到用户信息", nil)
		return
	}
	//接收参数  //使用结构体来接收

	// 2. 绑定并验证前端请求参数
	var req dto.CreateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Result(c, 400, "无效的请求格式: "+err.Error(), nil)
		return
	}
	// 3. 转换DTO为数据库模型（只传递需要的字段）
	expense := &model.Expense{
		UserID:      userID.(int), // 从上下文获取，前端无法篡改
		Note:        req.Note,
		Amount:      req.Amount,
		Remarks:     req.Remarks,
		ExpenseDate: req.ExpenseDate.ToTime(),
		Category:    req.Category,
		// IsExtended默认false，CreatedAt/UpdatedAt由数据库自动生成，无需手动设置
	}
	fmt.Println(expense)
	// 4. 调用服务层保存数据
	h.expnseService.AddExpense(expense)
	//if err != nil {
	//	util.Result(c, 500, "添加失败: "+err.Error(), nil)
	//	return
	//}
	util.Result(c, 200, "添加成功", nil)

}
