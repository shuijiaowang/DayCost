package handler

import (
	"DayCost/internal/dto"
	"DayCost/internal/model"
	"DayCost/internal/service"
	"DayCost/pkg/util"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ExpenseHandler 继承BaseHandler，复用公共方法
type ExpenseHandler struct {
	BaseHandler                            // 继承基础Handler  //添加
	expenseService *service.ExpenseService // 内部创建要比传参方便一些
}

// NewExpenseHandler 构造函数，
func NewExpenseHandler() *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: service.NewExpenseService(),
	}
}
func (h *ExpenseHandler) AddExpense(c *gin.Context) {

	// 复用BaseHandler的GetUserID，减少重复代码
	userID, ok := h.GetUserID(c)
	if !ok {
		return // 错误已在GetUserID中返回
	}

	// 2. 绑定并验证前端请求参数
	// 复用BaseHandler的Bind方法
	var req dto.ExpenseCreateRequest
	if !h.Bind(c, &req) {
		return // 绑定错误已处理
	}
	// 3. 转换DTO为数据库模型（只传递需要的字段）
	expense := &model.Expense{
		UserID:      userID, // 从上下文获取，前端无法篡改
		Note:        req.Note,
		Amount:      req.Amount,
		Remarks:     req.Remarks,
		ExpenseDate: req.ExpenseDate,
		Category:    req.Category,
		// IsExtended默认false，CreatedAt/UpdatedAt由数据库自动生成，无需手动设置
	}
	fmt.Println(expense)
	// 4. 调用服务层保存数据
	h.expenseService.AddExpense(expense)
	//if err != nil {
	//	util.Result(c, 500, "添加失败: "+err.Error(), nil)
	//	return
	//}
	util.Result(c, 200, "添加成功", nil)

}
func (h *ExpenseHandler) GetExpenseById(c *gin.Context) {

	userID, ok := h.GetUserID(c)
	if !ok {
		return
	}
	//怎么获取路径参数
	id := c.Param("id")
	// 将当前用户ID和要查询的ID一起传给Service
	expense, err := h.expenseService.GetExpenseById(id, strconv.Itoa(userID))
	if err != nil {
		// Service层返回了错误（例如：记录不存在或权限不足）
		util.Result(c, 403, err.Error(), nil) // 或者404，根据错误类型判断
		return
	}
	result := dto.ToResultExpense(expense)
	util.Result(c, 200, "查询成功", result)
}
func (h *ExpenseHandler) ListExpense(c *gin.Context) {
	userID, ok := h.GetUserID(c)
	if !ok {
		return
	}
	expense, err := h.expenseService.ListExpense(strconv.Itoa(userID)) //返回切片
	if err != nil {
		util.Result(c, 403, "false", nil)
	}
	var expenseList []dto.ExpenseResponse
	for i := 0; i < len(expense); i++ {
		expenseList = append(expenseList, dto.ToResultExpense(expense[i]))
	}
	util.Result(c, 200, "查询成功", expenseList)
}

// 条件查询+分页查询
// 示例：Repository 层处理分页
// Handler 层返回
func (h *ExpenseHandler) ListExpenseByCondition(c *gin.Context) {
	userID, ok := h.GetUserID(c)
	if !ok {
		return
	}
	var req dto.ExpensePagesQuery
	if !h.Bind(c, &req) {
		return // 绑定错误已处理
	}
	req.UserID = userID
	fmt.Println(req)
	expenses, total, err := h.expenseService.ListExpenseByCondition(req)
	if err != nil {
		util.Result(c, 500, "查询失败: "+err.Error(), nil)
		return
	}
	// 包装分页响应
	resp := dto.PaginationResponse{
		Total:    total,        //总页数
		Page:     req.Page,     //查询页码，第几页
		PageSize: req.PageSize, //每页条数
		Data:     expenses,     // 转换为DTO后的数据
	}
	util.Result(c, 200, "查询成功", resp)
}
