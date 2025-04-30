package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	transactiondto "backend/dtos/transaction"
	"backend/pkg/response"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type TransactionUsecase interface {
	Save(dto transactiondto.CreateTransactionDto, userID uint) (*entity.TransactionEntity, error)
	Detail(id uint) (*entity.TransactionEntity, error)
}

// definisii struct buat usecase brandusecase dan voucher usecase buat pengecehkan data atau tidak
type useTransactionUsecase struct {
	transactionRepo          repository.TransactionRepository
	voucherUsecase           VoucherUsecase
	brandUsecase             BrandUsecase
	transactionDetailUsecase TransactionDetailUsecase
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository, voucherUsecase VoucherUsecase, brandUsecase BrandUsecase, transactionDetailUsecase TransactionDetailUsecase) TransactionUsecase {

	return &useTransactionUsecase{
		transactionRepo:          transactionRepo,
		voucherUsecase:           voucherUsecase,
		brandUsecase:             brandUsecase,
		transactionDetailUsecase: transactionDetailUsecase,
	}
}

func (uc *useTransactionUsecase) Save(dto transactiondto.CreateTransactionDto, userID uint) (*entity.TransactionEntity, error) {
	// defining usecase

	// check jika ada data brand

	brandData, err := uc.brandUsecase.Detail(uint(dto.BrandId))
	if err != nil {
		return nil, err
	}

	voucherData, err := uc.voucherUsecase.FindByMultipleId(dto.VocuhersId)
	if err != nil {
		return nil, response.NewErrorResponse(400, "Voucher Id tidak Valid", nil)
	}

	if len(voucherData) != len(dto.VocuhersId) {
		return nil, response.NewErrorResponse(400, "Salah Satu Voucher Tidak Valid", nil)
	}

	var transactionDetailEntities []entity.TransactionDetailEntity

	// get total point from voucher
	var total float64 = 0
	for _, voucher := range voucherData {
		total += voucher.Point
	}
	transactionEntity := entity.TransactionEntity{
		BrandName:  brandData.Name,
		BrandId:    brandData.ID,
		UserId:     userID,
		TotalPoint: total,
	}
	// created transaction
	createdTransaction, err := uc.transactionRepo.Save(&transactionEntity)
	if err != nil {
		return nil, response.NewErrorResponse(500, "Failed Created Transaction", nil)
	}

	// loop data voucher brand into transaction details entity
	for _, voucher := range voucherData {
		transactionDetailEntities = append(transactionDetailEntities, entity.TransactionDetailEntity{
			VoucherId: voucher.ID,
			BrandName: brandData.Name,
			BrandId:   brandData.ID,
			Point:     voucher.Point, TransactionId: createdTransaction.ID,
		})
	}

	//bulk inser transaction detail transaction
	_, err = uc.transactionDetailUsecase.BulkSave(transactionDetailEntities)
	if err != nil {
		return nil, response.NewErrorResponse(500, "Failed Created Detail Transaction", nil)

	}
	// get data transaction detail
	findTransaction, err := uc.Detail(createdTransaction.ID)
	if err != nil {
		return nil, err
	}
	return findTransaction, nil
}

func (uc *useTransactionUsecase) Detail(id uint) (*entity.TransactionEntity, error) {

	findData, err := uc.transactionRepo.Detail(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, response.NewErrorResponse(http.StatusNotFound, "transaction data tidak ada", nil)
	}
	return findData, nil
}
