package postgres_gorm_repository_impl

import (
	"backend/domain/entity"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type useTransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepositoryImpl(db *gorm.DB) repository.TransactionRepository {
	return &useTransactionRepositoryImpl{
		DB: db,
	}
}

func (repo *useTransactionRepositoryImpl) Save(entity *entity.TransactionEntity) (*entity.TransactionEntity, error) {
	if err := repo.DB.Create(&entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (repo *useTransactionRepositoryImpl) Detail(id uint) (*entity.TransactionEntity, error) {
	var transaction entity.TransactionEntity

	err := repo.DB.Model(&entity.TransactionEntity{}).Where("id = ?", id).Preload("Brand").Preload("TransactionDetails").First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
