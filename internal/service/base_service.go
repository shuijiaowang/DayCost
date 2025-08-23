package service

import "DayCost/internal/repository"

type BaseService struct {
	baseRepo *repository.BaseRepository
}

func NewBaseService() *BaseService {
	return &BaseService{
		baseRepo: &repository.BaseRepository{},
	}
}

func (b *BaseService) CheckExpenseOwner(userID, expenseID int) error {
	return b.baseRepo.CheckExpenseOwner(expenseID, userID)
}
