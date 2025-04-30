package usecase

import (
	entity "backend/domain/entity"
	repository "backend/domain/repository"
	voucherdto "backend/dtos/Voucher"
	response "backend/pkg/response"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

// defining interface voucher interface
type VoucherUsecase interface {
	Save(dto voucherdto.CreateVoucherDto) (*entity.VoucherEntity, error)
	FindByMultipleId(multipleID []uint) ([]*entity.VoucherEntity, error)
	Detail(id uint) (*entity.VoucherEntity, error)
	FindByBrandId(brand_id uint) ([]*entity.VoucherEntity, error)
}

// defining use vouhcer usecase
type useVoucherUsecase struct {
	brandUsecase BrandUsecase
	repo         repository.VoucherRepository
}

func NewVoucherUsecase(repo repository.VoucherRepository, brandUsecase BrandUsecase) VoucherUsecase {
	return &useVoucherUsecase{
		brandUsecase: brandUsecase,
		repo:         repo,
	}
}

func (uc *useVoucherUsecase) Save(dto voucherdto.CreateVoucherDto) (*entity.VoucherEntity, error) {
	// data, err := uc.repo.save(dto)

	// check data brand jika tidak ada
	_, err := uc.brandUsecase.Detail(dto.BrandId)
	if err != nil {
		return nil, err
	}
	// membuat data voucher
	data, err := uc.repo.Save(dto)
	if err != nil {
		return nil, response.NewErrorResponse(500, "Gagal membuat data voucher", nil)
	}

	return data, nil

}

func (uc *useVoucherUsecase) FindByMultipleId(multipleId []uint) ([]*entity.VoucherEntity, error) {
	data, err := uc.repo.FindByMultipleId(multipleId)

	if err != nil {
		return nil, response.NewErrorResponse(400, "Voucher ID tidak valid", nil)

	}

	return data, nil
}

func (uc *useVoucherUsecase) Detail(id uint) (*entity.VoucherEntity, error) {
	data, err := uc.repo.Detail(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, response.NewErrorResponse(http.StatusNotFound, "Data voucher tidak ada", nil)
	}

	return data, nil
}

func (uc *useVoucherUsecase) FindByBrandId(id uint) ([]*entity.VoucherEntity, error) {
	data, err := uc.repo.FindByBrandId(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, response.NewErrorResponse(http.StatusNotFound, "Data voucher tidak ada", nil)
	}

	return data, nil
}
