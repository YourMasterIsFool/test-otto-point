package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	branddto "backend/dtos/brand"
	"backend/pkg/response"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type BrandUsecase interface {
	Save(createDto branddto.CreateBrand) (*entity.BrandEntity, error)
	Detail(id uint) (*entity.BrandEntity, error)
}

type useBrandUsecase struct {
	brandRepo repository.BrandRepository
}

func NewBrandUsecase(repo repository.BrandRepository) *useBrandUsecase {
	return &useBrandUsecase{
		brandRepo: repo,
	}
}

func (u *useBrandUsecase) Save(dtos branddto.CreateBrand) (*entity.BrandEntity, error) {
	data, err := u.brandRepo.Save(dtos)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *useBrandUsecase) Detail(id uint) (*entity.BrandEntity, error) {
	data, err := u.brandRepo.Detail(id)

	// check data brand kalo tidak ada retunr custom message
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// fmt.Println("test ")
		return nil, response.NewErrorResponse(http.StatusNotFound, "data brand tidak ada", nil)
	}
	return data, nil
}
