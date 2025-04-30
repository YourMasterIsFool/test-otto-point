package postgres_gorm_repository_impl

import (
	"backend/domain/entity"
	"backend/domain/repository"
	branddto "backend/dtos/brand"

	"gorm.io/gorm"
)

type useBrandRepositoryImpl struct {
	DB *gorm.DB
}

func NewBrandRepositoryImpl(db *gorm.DB) repository.BrandRepository {
	return &useBrandRepositoryImpl{
		DB: db,
	}
}

func (repo *useBrandRepositoryImpl) Save(dto branddto.CreateBrand) (*entity.BrandEntity, error) {

	// convert data kedalam entity berdasarkan gorm
	var data = entity.BrandEntity{
		Name: dto.Name,
	}

	// create data jika ada error;
	if err := repo.DB.Create(&data).Error; err != nil {
		return nil, err
	}

	// cari data hasil create tadi berdasarkan id
	find, err := repo.Detail(data.ID)
	if err != nil {
		return nil, err
	}

	return find, err

}
func (repo *useBrandRepositoryImpl) Detail(id uint) (*entity.BrandEntity, error) {
	var data entity.BrandEntity

	// check data brand berdasarkan id jika ada
	if err := repo.DB.Where("id = ? ", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
