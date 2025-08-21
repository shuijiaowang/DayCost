package service

import (
	"DayCost/internal/model"
	"DayCost/internal/repository"
)

type ExpenseExtService struct {
	expenseExtRepo *repository.ExpenseExtRepository
	baseRepo       *repository.BaseRepository
	expenseRepo    *repository.ExpenseRepository
}

func NewExpenseExtService() *ExpenseExtService {
	return &ExpenseExtService{
		expenseExtRepo: &repository.ExpenseExtRepository{},
	}
}

// err := h.expenseExtService.AddExpenseExt(userID,req)
func (h *ExpenseExtService) AddExpenseExt(userID int, req *model.ExpenseExt) error {
	err := h.baseRepo.CheckExpenseOwner(userID, req.ExpenseID)
	if err != nil {
		return err
	}
	h.expenseRepo.UpdateIsExtended(req.ExpenseID, userID, true)
	return h.expenseExtRepo.AddExpenseExt(req)
}
