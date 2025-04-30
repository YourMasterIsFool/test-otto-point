package repository

import "backend/domain/entity"

type TransactionDetailRepository interface {
	BulkSave(entity []entity.TransactionDetailEntity) ([]*entity.TransactionDetailEntity, error)
}
