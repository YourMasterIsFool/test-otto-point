package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/pkg/response"
)

type useTransactionDetailUsecase struct {
	repo repository.TransactionDetailRepository
}

type TransactionDetailUsecase interface {
	BulkSave([]entity.TransactionDetailEntity) ([]*entity.TransactionDetailEntity, error)
}

func NewTransactionDetailUsecase(repo repository.TransactionDetailRepository) TransactionDetailUsecase {
	return &useTransactionDetailUsecase{
		repo: repo,
	}
}

func (uc *useTransactionDetailUsecase) BulkSave(entities []entity.TransactionDetailEntity) ([]*entity.TransactionDetailEntity, error) {
	datas, err := uc.repo.BulkSave(entities)

	if err != nil {
		return nil, response.NewErrorResponse(500, "Gagal membuat transaction detail", err)
	}

	return datas, nil
}
