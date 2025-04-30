package postgres_gorm_repository_impl

import (
	"backend/domain/entity"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type useAuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) repository.AuthRepository {
	return &useAuthRepositoryImpl{
		DB: db,
	}
}

func (repo *useAuthRepositoryImpl) Save(schema entity.UserEntity) (*entity.UserEntity, error) {

	// create data jika ada error;
	if err := repo.DB.Create(&schema).Error; err != nil {

		return nil, err
	}

	// cari data hasil create tadi berdasarkan id
	find, err := repo.Detail(schema.ID)
	if err != nil {
		return nil, err
	}

	return find, err

}
func (repo *useAuthRepositoryImpl) Detail(id uint) (*entity.UserEntity, error) {
	var data entity.UserEntity

	// check data brand berdasarkan id jika ada
	if err := repo.DB.Where("id = ? ", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *useAuthRepositoryImpl) FindByUsername(username string) (*entity.UserEntity, error) {
	var data entity.UserEntity

	// check data brand berdasarkan id jika ada
	if err := repo.DB.Where("username = ? ", username).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
