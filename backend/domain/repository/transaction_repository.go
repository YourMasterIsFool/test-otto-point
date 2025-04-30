package repository

import (
	"backend/domain/entity"
)

type TransactionRepository interface {
	Save(entity *entity.TransactionEntity) (*entity.TransactionEntity, error)
	Detail(id uint) (*entity.TransactionEntity, error)
}
