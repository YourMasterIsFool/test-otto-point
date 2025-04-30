package repository

import (
	"backend/domain/entity"
	voucherdto "backend/dtos/Voucher"
)

type VoucherRepository interface {
	Save(dto voucherdto.CreateVoucherDto) (*entity.VoucherEntity, error)
	Detail(id uint) (*entity.VoucherEntity, error)
	FindByBrandId(brandID uint) ([]*entity.VoucherEntity, error)
	FindByMultipleId(multipleID []uint) ([]*entity.VoucherEntity, error)
}
