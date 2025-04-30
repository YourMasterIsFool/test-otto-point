package postgres_gorm_repository_impl

import (
	"backend/domain/entity"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type useTransactionDetailRepository struct {
	DB *gorm.DB
}

func NewTransactionDetailRepository(db *gorm.DB) repository.TransactionDetailRepository {
	return &useTransactionDetailRepository{
		DB: db,
	}
}

func (repo *useTransactionDetailRepository) BulkSave(entities []entity.TransactionDetailEntity) ([]*entity.TransactionDetailEntity, error) {
	if err := repo.DB.Create(&entities).Error; err != nil {
		return nil, err
	}

	var result []*entity.TransactionDetailEntity
	for i := range entities {
		result = append(result, &entities[i])
	}

	return result, nil

}
