package repository

import (
	"backend/domain/entity"
	brand "backend/domain/entity"
	branddto "backend/dtos/brand"
)

type BrandRepository interface {
	Save(dto branddto.CreateBrand) (*entity.BrandEntity, error)
	Detail(id uint) (*brand.BrandEntity, error)
}
