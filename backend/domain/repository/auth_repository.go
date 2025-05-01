package repository

import (
	"backend/domain/entity"
)

// definign auth repository
type AuthRepository interface {
	FindByUsername(email string) (*entity.UserEntity, error)
	Save(schema entity.UserEntity) (*entity.UserEntity, error)
	Detail(id uint) (*entity.UserEntity, error)
}
