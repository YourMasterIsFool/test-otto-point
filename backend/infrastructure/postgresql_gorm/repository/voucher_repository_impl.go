package postgres_gorm_repository_impl

import (
	"backend/domain/entity"
	"backend/domain/repository"
	voucherdto "backend/dtos/voucher"

	"gorm.io/gorm"
)

type useVoucherRepositoryImpl struct {
	DB *gorm.DB
}

func NewVoucherRepositoryImpl(db *gorm.DB) repository.VoucherRepository {
	return &useVoucherRepositoryImpl{
		DB: db,
	}
}
func (repo *useVoucherRepositoryImpl) Save(dto voucherdto.CreateVoucherDto) (*entity.VoucherEntity, error) {

	// convert data kedalam entity berdasarkan gorm
	data := entity.VoucherEntity{
		BrandId: dto.BrandId,
		Point:   dto.Point,
	}

	// create data jika ada error;
	if err := repo.DB.Create(&data).Error; err != nil {
		return nil, err
	}

	// find voucher data
	find, err := repo.Detail(data.ID)
	if err != nil {
		return nil, err
	}

	return find, err

}
func (repo *useVoucherRepositoryImpl) Detail(id uint) (*entity.VoucherEntity, error) {
	var data entity.VoucherEntity

	// check data Voucher berdasarkan id jika ada
	if err := repo.DB.Where("id = ? ", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// check voucher berdasarkan brandId
func (repo *useVoucherRepositoryImpl) FindByBrandId(brandId uint) ([]*entity.VoucherEntity, error) {

	// define list voucher entitty
	var listVoucher []*entity.VoucherEntity

	// get list voucher entity
	if err := repo.DB.Model(&entity.VoucherEntity{}).Where("brand_id = ?", brandId).Find(&listVoucher).Error; err != nil {
		return nil, err
	}
	return listVoucher, nil
}

// check multiple data voucher
func (repo *useVoucherRepositoryImpl) FindByMultipleId(multipleID []uint) ([]*entity.VoucherEntity, error) {
	var listVoucher []*entity.VoucherEntity

	if err := repo.DB.Where("id in ?", multipleID).Find(&listVoucher).Error; err != nil {
		return nil, err
	}
	return listVoucher, nil
}
