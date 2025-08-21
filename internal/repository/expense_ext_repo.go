package repository

import (
	"DayCost/internal/model"
	"DayCost/pkg/database"
)

// expense_ext_repo.go 继承基础Repository
type ExpenseExtRepository struct {
	BaseRepository // 嵌入基础Repository，复用CheckExpenseOwner
}

// 使用时直接调用父类方法
func (r *ExpenseExtRepository) Test1(ext *model.ExpenseExt, userID int) error {
	if err := r.CheckExpenseOwner(ext.ExpenseID, userID); err != nil {
		return err
	}
	return nil
}

// h.expenseExtRepo.AddExpenseExt(req)
// 函数声明开始
func (r *ExpenseExtRepository) AddExpenseExt(expenseExt *model.ExpenseExt) error {
	tx := database.DB.Create(expenseExt)
	return tx.Error
}
