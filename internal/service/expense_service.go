package service

import (
	"DayCost/internal/repository"
	"time"
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
func (s *ExpenseService) AddExpense(note string, amount float64, category int, expenseDate time.Time) {
	s.expenseRepo.AddExpense(note, amount, category, expenseDate)
	return
}
