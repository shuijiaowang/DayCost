package repository

import (
	"DayCost/internal/model"
	"DayCost/pkg/database"
)

type ExpenseRepository struct{}

func (r *ExpenseRepository) AddExpense(expense *model.Expense) {
	database.DB.Create(expense) //GORM 在插入记录后，可能需要修改传入的结构体实例，以便将数据库生成的值（如自增ID、时间戳等）回填到结构体中。所以需要传入指针
	return
}
