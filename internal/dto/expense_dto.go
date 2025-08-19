// DayCost/internal/dto/expense_dto.go
package dto

import (
	//"github.com/go-playground/validator/v10"
	"DayCost/internal/model"
)

// CreateExpenseRequest 接收前端新增消费的请求参数
type CreateExpenseRequest struct {
	Note        string         `json:"note" binding:"required,max=100"`                          // 物品名称/摘要（必填，最长100字符）
	Amount      float64        `json:"amount" binding:"required,gt=0"`                           // 金额（必填，大于0）
	Remarks     string         `json:"remarks" binding:"omitempty,max=500"`                      // 备注（可选，最长500字符）
	ExpenseDate model.JSONDate `json:"expense_date" binding:"required" time_format:"2006-01-02"` // 消费日期（必填）
	Category    int8           `json:"category" binding:"required,gte=0,lte=9"`                  // 分类（必填，0-9范围内）
}
