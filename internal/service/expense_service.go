package service

import (
	"DayCost/internal/model"
	"DayCost/internal/repository"
)

type ExpenseService struct {
	expenseRepo *repository.ExpenseRepository
}

func NewExpenseService() *ExpenseService {
	return &ExpenseService{
		expenseRepo: &repository.ExpenseRepository{},
	}
}

// 这里得一个一个参数传递，还是封装个vo对象
func (s *ExpenseService) AddExpense(expense *model.Expense) {
	s.expenseRepo.AddExpense(expense)
	return
}
