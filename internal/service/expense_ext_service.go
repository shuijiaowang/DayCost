package service

import (
	"DayCost/internal/model"
	"DayCost/internal/repository"
)

type ExpenseExtService struct {
	BaseService
	expenseExtRepo *repository.ExpenseExtRepository
	expenseRepo    *repository.ExpenseRepository
}

func NewExpenseExtService() *ExpenseExtService {
	return &ExpenseExtService{
		expenseExtRepo: &repository.ExpenseExtRepository{},
	}
}

// err := h.expenseExtService.AddExpenseExt(userID,req)
func (h *ExpenseExtService) AddExpenseExt(userID int, req *model.ExpenseExt) error {
	//err := h.baseRepo.CheckExpenseOwner(userID, req.ExpenseID)
	//if err != nil {
	//	return err
	//}
	err := h.expenseRepo.UpdateIsExtended(req.ExpenseID, userID, true)
	if err != nil {
		return err
	}
	err = h.expenseExtRepo.AddExpenseExt(req)
	if err != nil {
		return err
	}
	return nil
}
