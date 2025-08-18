package repository

import (
	"DayCost/internal/model"
	"DayCost/pkg/database"
	"time"
)

type ExpenseRepository struct{}

func (r *ExpenseRepository) AddExpense(note string, amount float64, category int, expenseDate time.Time) {
	expense := model.Expense{
		UserId:      1,
		Note:        note,
		Amount:      amount,
		Category:    category,
		ExpenseDate: expenseDate,
	}
	database.DB.Create(expense)
	return
}
